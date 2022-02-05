package keeper

import (
	"context"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SubmitCopyrightProposal(goCtx context.Context, msg *types.MsgSubmitCopyrightProposal) (*types.MsgSubmitCopyrightProposalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSubmitCopyrightProposalResponse{}, nil
}
