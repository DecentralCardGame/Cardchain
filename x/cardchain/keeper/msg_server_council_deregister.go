package keeper

import (
	"context"

	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CouncilDeregister(goCtx context.Context, msg *types.MsgCouncilDeregister) (*types.MsgCouncilDeregisterResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCouncilDeregisterResponse{}, nil
}
