package keeper

import (
	"context"

	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CouncilResponseCommit(goCtx context.Context, msg *types.MsgCouncilResponseCommit) (*types.MsgCouncilResponseCommitResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCouncilResponseCommitResponse{}, nil
}
