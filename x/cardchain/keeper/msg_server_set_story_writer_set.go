package keeper

import (
	"context"

	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SetStoryWriterSet(goCtx context.Context, msg *types.MsgSetStoryWriterSet) (*types.MsgSetStoryWriterSetResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSetStoryWriterSetResponse{}, nil
}
