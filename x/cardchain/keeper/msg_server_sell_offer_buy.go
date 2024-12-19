package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) SellOfferBuy(goCtx context.Context, msg *types.MsgSellOfferBuy) (*types.MsgSellOfferBuyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	buyer, err := k.GetUserFromString(ctx, msg.Creator)
	if err != nil {
		return nil, errorsmod.Wrap(types.ErrUserDoesNotExist, err.Error())
	}

	sellOffer := k.SellOffers.Get(ctx, msg.SellOfferId)

	sellerAddr, err := sdk.AccAddressFromBech32(sellOffer.Seller)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "Unable to convert to AccAddress")
	}

	if sellOffer.Status != types.SellOfferStatus_open || sellOffer.Seller == "" {
		return nil, errorsmod.Wrapf(sdkerrors.ErrUnauthorized, "Selloffer Status: %v", sellOffer.Status)
	}

	err = k.BankKeeper.SendCoins(ctx, buyer.Addr, sellerAddr, sdk.Coins{sellOffer.Price})
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInsufficientFunds, err.Error())
	}

	buyer.Cards = append(buyer.Cards, sellOffer.Card)
	sellOffer.Buyer = msg.Creator
	sellOffer.Status = types.SellOfferStatus_sold

	k.SellOffers.Set(ctx, msg.SellOfferId, sellOffer)
	k.SetUserFromUser(ctx, buyer)

	return &types.MsgSellOfferBuyResponse{}, nil
}
