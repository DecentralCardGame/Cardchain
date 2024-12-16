package keeper

import (
	"context"

	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CouncilResponseReveal(goCtx context.Context, msg *types.MsgCouncilResponseReveal) (*types.MsgCouncilResponseRevealResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCouncilResponseRevealResponse{}, nil
}
