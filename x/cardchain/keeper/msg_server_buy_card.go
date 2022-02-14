package keeper

import (
	"context"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) BuyCard(goCtx context.Context, msg *types.MsgBuyCard) (*types.MsgBuyCardResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	buyerAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidAccAddress, "Unable to convert to AccAddress")
	}
	buyer, err := k.GetUser(ctx, buyerAddr)
	if err != nil {
		return nil, err
	}

	sellOffer := k.GetSellOffer(ctx, msg.SellOfferId)

	sellerAddr, err := sdk.AccAddressFromBech32(sellOffer.Seller)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidAccAddress, "Unable to convert to AccAddress")
	}

	if sellOffer.Status != types.SellOfferStatus_open || sellOffer.Seller == "" {
		return nil, sdkerrors.Wrapf(types.ErrNoOpenSellOffer, "Status: %v", sellOffer.Status)
	}

	err = k.BankKeeper.SendCoins(ctx, buyerAddr, sellerAddr, sdk.Coins{sdk.NewInt64Coin("credits", int64(sellOffer.Price))})
	if err != nil {
		return nil, err
	}

	buyer.Cards = append(buyer.Cards, sellOffer.Card)
	sellOffer.Buyer = msg.Creator
	sellOffer.Status = types.SellOfferStatus_sold

	k.SetSellOffer(ctx, msg.SellOfferId, sellOffer)
	k.SetUser(ctx, buyerAddr, buyer)

	return &types.MsgBuyCardResponse{}, nil
}
