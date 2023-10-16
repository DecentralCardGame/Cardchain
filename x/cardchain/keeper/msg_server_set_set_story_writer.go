package keeper

import (
	"context"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SetSetStoryWriter(goCtx context.Context, msg *types.MsgSetSetStoryWriter) (*types.MsgSetSetStoryWriterResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	set := k.Sets.Get(ctx, msg.SetId)
	err := checkSetEditable(set, msg.Creator)
	if err != nil {
		return nil, err
	}

	set.StoryWriter = msg.StoryWriter

	k.Sets.Set(ctx, msg.SetId, set)

	return &types.MsgSetSetStoryWriterResponse{}, nil
}
