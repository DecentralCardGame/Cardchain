package keeper

import (
	"context"
	"strconv"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) FinalizeCollection(goCtx context.Context, msg *types.MsgFinalizeCollection) (*types.MsgFinalizeCollectionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	collectionSize := int(k.GetParams(ctx).CollectionSize)

	collection := k.GetCollection(ctx, msg.CollectionId)
	if msg.Creator != collection.Contributors[0] {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Invalid creator")
	}

	if collection.Status != types.CStatus_design {
		return nil, types.ErrCollectionNotInDesign
	}

	if len(collection.Cards) != collectionSize {
		return nil, sdkerrors.Wrap(types.ErrCollectionSize, "Has to be "+strconv.Itoa(collectionSize))
	}

	// TODO Add checking for rarity

	collection.Status = types.CStatus_finalized

	k.SetCollection(ctx, msg.CollectionId, collection)

	return &types.MsgFinalizeCollectionResponse{}, nil
}
