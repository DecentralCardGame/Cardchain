package keeper

import (
	"context"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SubmitMatchReporterProposal(goCtx context.Context, msg *types.MsgSubmitMatchReporterProposal) (*types.MsgSubmitMatchReporterProposalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSubmitMatchReporterProposalResponse{}, nil
}
