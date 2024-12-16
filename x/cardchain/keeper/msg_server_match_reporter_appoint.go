package keeper

import (
	"context"

	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) MatchReporterAppoint(goCtx context.Context, msg *types.MsgMatchReporterAppoint) (*types.MsgMatchReporterAppointResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgMatchReporterAppointResponse{}, nil
}
