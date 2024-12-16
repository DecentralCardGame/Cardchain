package keeper

import (
	"context"

	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SetStoryAdd(goCtx context.Context, msg *types.MsgSetStoryAdd) (*types.MsgSetStoryAddResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSetStoryAddResponse{}, nil
}
