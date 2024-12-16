package keeper

import (
	"context"

	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) BoosterPackOpen(goCtx context.Context, msg *types.MsgBoosterPackOpen) (*types.MsgBoosterPackOpenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgBoosterPackOpenResponse{}, nil
}
