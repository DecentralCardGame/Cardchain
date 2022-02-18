package keeper

import (
	"context"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) AddStoryToCollection(goCtx context.Context, msg *types.MsgAddStoryToCollection) (*types.MsgAddStoryToCollectionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	collection := k.GetCollection(ctx, msg.CollectionId)
	if collection.StoryWriter != msg.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect StoryWriter")
	}
	if collection.Status != types.CStatus_design {
		return nil, types.ErrCollectionNotInDesign
	}

	collection.Story = msg.Story

	k.SetCollection(ctx, msg.CollectionId, collection)

	return &types.MsgAddStoryToCollectionResponse{}, nil
}
