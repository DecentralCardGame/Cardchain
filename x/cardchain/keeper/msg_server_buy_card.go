package keeper

import (
	"context"

	sdkerrors "cosmossdk.io/errors"
	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) BuyCard(goCtx context.Context, msg *types.MsgBuyCard) (*types.MsgBuyCardResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	buyer, err := k.GetUserFromString(ctx, msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrUserDoesNotExist, err.Error())
	}

	sellOffer := k.SellOffers.Get(ctx, msg.SellOfferId)

	sellerAddr, err := sdk.AccAddressFromBech32(sellOffer.Seller)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidAccAddress, "Unable to convert to AccAddress")
	}

	if sellOffer.Status != types.SellOfferStatus_open || sellOffer.Seller == "" {
		return nil, sdkerrors.Wrapf(types.ErrNoOpenSellOffer, "Status: %v", sellOffer.Status)
	}

	err = k.BankKeeper.SendCoins(ctx, buyer.Addr, sellerAddr, sdk.Coins{sellOffer.Price})
	if err != nil {
		return nil, sdkerrors.Wrap(errors.ErrInsufficientFunds, err.Error())
	}

	buyer.Cards = append(buyer.Cards, sellOffer.Card)
	sellOffer.Buyer = msg.Creator
	sellOffer.Status = types.SellOfferStatus_sold

	k.SellOffers.Set(ctx, msg.SellOfferId, sellOffer)
	k.SetUserFromUser(ctx, buyer)

	return &types.MsgBuyCardResponse{}, nil
}
