package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) SellOfferRemove(goCtx context.Context, msg *types.MsgSellOfferRemove) (*types.MsgSellOfferRemoveResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	creator, err := k.GetUserFromString(ctx, msg.Creator)
	if err != nil {
		return nil, errorsmod.Wrap(types.ErrUserDoesNotExist, err.Error())
	}

	sellOffer := k.SellOffers.Get(ctx, msg.SellOfferId)

	if sellOffer.Seller != msg.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Seller")
	}

	if sellOffer.Status != types.SellOfferStatus_open {
		return nil, errorsmod.Wrapf(sdkerrors.ErrUnauthorized, "Invalid status: %v", sellOffer.Status)
	}

	sellOffer.Status = types.SellOfferStatus_removed
	creator.Cards = append(creator.Cards, sellOffer.Card)

	k.SellOffers.Set(ctx, msg.SellOfferId, sellOffer)
	k.SetUserFromUser(ctx, creator)

	return &types.MsgSellOfferRemoveResponse{}, nil
}
