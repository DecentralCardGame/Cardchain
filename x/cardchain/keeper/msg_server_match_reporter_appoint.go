package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) MatchReporterAppoint(goCtx context.Context, msg *types.MsgMatchReporterAppoint) (*types.MsgMatchReporterAppointResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if k.authority != msg.Authority {
		return nil, errorsmod.Wrapf(types.ErrInvalidSigner, "expected %s got %s", k.authority, msg.Authority)
	}

	return &types.MsgMatchReporterAppointResponse{}, k.SetMatchReporter(ctx, msg.Reporter)
}
