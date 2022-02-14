package keeper

import (
	"context"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) RemoveSellOffer(goCtx context.Context, msg *types.MsgRemoveSellOffer) (*types.MsgRemoveSellOfferResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	creator, err := k.GetUserFromString(ctx, msg.Creator)
	if err != nil {
		return nil, err
	}

	sellOffer := k.GetSellOffer(ctx, msg.SellOfferId)

	if sellOffer.Seller != msg.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Seller")
	}

	if sellOffer.Status != types.SellOfferStatus_open{
		return nil, sdkerrors.Wrapf(types.ErrNoOpenSellOffer, "Status: %v", sellOffer.Status)
	}

	sellOffer.Status = types.SellOfferStatus_removed
	creator.Cards = append(creator.Cards, sellOffer.Card)

	k.SetSellOffer(ctx, msg.SellOfferId, sellOffer)
	k.SetUserFromUser(ctx, creator)

	return &types.MsgRemoveSellOfferResponse{}, nil
}
