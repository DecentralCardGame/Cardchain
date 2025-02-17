package keeper

import (
	"context"

	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) MatchReporterAppoint(goCtx context.Context, msg *types.MsgMatchReporterAppoint) (*types.MsgMatchReporterAppointResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	return &types.MsgMatchReporterAppointResponse{}, k.SetMatchReporter(ctx, msg.Reporter)
}
