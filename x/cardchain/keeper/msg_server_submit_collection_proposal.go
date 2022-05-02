package keeper

import (
	"context"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SubmitCollectionProposal(goCtx context.Context, msg *types.MsgSubmitCollectionProposal) (*types.MsgSubmitCollectionProposalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSubmitCollectionProposalResponse{}, nil
}
