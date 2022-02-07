package keeper

import (
	"context"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) ApointMatchReporter(goCtx context.Context, msg *types.MsgApointMatchReporter) (*types.MsgApointMatchReporterResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgApointMatchReporterResponse{}, nil
}
