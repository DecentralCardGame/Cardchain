package keeper

import (
	"context"

	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) MatchConfirm(goCtx context.Context, msg *types.MsgMatchConfirm) (*types.MsgMatchConfirmResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgMatchConfirmResponse{}, nil
}
