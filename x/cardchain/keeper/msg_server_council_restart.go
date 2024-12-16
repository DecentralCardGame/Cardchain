package keeper

import (
	"context"

	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CouncilRestart(goCtx context.Context, msg *types.MsgCouncilRestart) (*types.MsgCouncilRestartResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCouncilRestartResponse{}, nil
}
