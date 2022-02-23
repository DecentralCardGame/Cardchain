package keeper

import (
	"context"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CommitCouncilResponse(goCtx context.Context, msg *types.MsgCommitCouncilResponse) (*types.MsgCommitCouncilResponseResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCommitCouncilResponseResponse{}, nil
}
