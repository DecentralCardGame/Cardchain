package keeper

import (
	"context"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SetCollectionStoryWriter(goCtx context.Context, msg *types.MsgSetCollectionStoryWriter) (*types.MsgSetCollectionStoryWriterResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	collection := k.Collections.Get(ctx, msg.CollectionId)
	err := checkCollectionEditable(collection, msg.Creator)
	if err != nil {
		return nil, err
	}

	collection.StoryWriter = msg.StoryWriter

	k.Collections.Set(ctx, msg.CollectionId, collection)

	return &types.MsgSetCollectionStoryWriterResponse{}, nil
}
