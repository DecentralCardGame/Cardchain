package keeper

import (
	"context"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CommitCouncilResponse(goCtx context.Context, msg *types.MsgCommitCouncilResponse) (*types.MsgCommitCouncilResponseResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	council := k.GetCouncil(ctx, msg.CouncilId)
	if !stringItemInList(msg.Creator, council.Voters) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Invalid Voter")
	}

	var allreadyVoted []string
	for _, response := range council.HashResponses {
		allreadyVoted = append(allreadyVoted, response.User)
	}

	if stringItemInList(msg.Creator, allreadyVoted) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Allready voted")
	}

	resp := types.WrapHashResponse{msg.Creator, msg.Response}
	council.HashResponses = append(council.HashResponses, &resp)

	if len(council.HashResponses) == 5 {
		council.Status = types.CouncelingStatus_commited
	}
	k.SetCouncil(ctx, msg.CouncilId, council)

	return &types.MsgCommitCouncilResponseResponse{}, nil
}
