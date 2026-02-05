package keeper

import (
	"context"

	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SetStoryWriterSet(goCtx context.Context, msg *types.MsgSetStoryWriterSet) (*types.MsgSetStoryWriterSetResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	set := k.SetK.Get(ctx, msg.SetId)
	err := checkSetEditable(set, msg.Creator)
	if err != nil {
		return nil, err
	}

	set.StoryWriter = msg.StoryWriter

	k.SetK.Set(ctx, msg.SetId, set)

	return &types.MsgSetStoryWriterSetResponse{}, nil
}
