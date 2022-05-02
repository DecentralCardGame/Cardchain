package keeper

import (
	"context"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"golang.org/x/exp/slices"
)

func (k msgServer) RemoveCardFromCollection(goCtx context.Context, msg *types.MsgRemoveCardFromCollection) (*types.MsgRemoveCardFromCollectionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	collection := k.Collections.Get(ctx, msg.CollectionId)
	if !slices.Contains(collection.Contributors, msg.Creator) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Invalid contributor")
	}
	if collection.Status != types.CStatus_design {
		return nil, types.ErrCollectionNotInDesign
	}

	newCards, err := PopItemFromArr(msg.CardId, collection.Cards)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "Card: %d", msg.CardId)
	}

	collection.Cards = newCards

	k.Collections.Set(ctx, msg.CollectionId, collection)

	return &types.MsgRemoveCardFromCollectionResponse{}, nil
}
