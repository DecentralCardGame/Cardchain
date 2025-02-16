package keeper

import (
	"context"
	"slices"

	errorsmod "cosmossdk.io/errors"
	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CouncilResponseCommit(goCtx context.Context, msg *types.MsgCouncilResponseCommit) (*types.MsgCouncilResponseCommitResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	collateralDeposit := k.GetParams(ctx).CollateralDeposit

	creator, err := k.GetUserFromString(ctx, msg.Creator)
	if err != nil {
		return nil, errorsmod.Wrap(types.ErrUserDoesNotExist, err.Error())
	}

	council := k.councils.Get(ctx, msg.CouncilId)
	if !slices.Contains(council.Voters, msg.Creator) {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "Invalid Voter")
	}

	if council.Status != types.CouncelingStatus_councilCreated {
		return nil, errorsmod.Wrapf(sdkerrors.ErrUnauthorized, "invalid status, have '%s', want '%s'", council.Status.String(), types.CouncelingStatus_councilCreated.String())
	}

	var allreadyVoted []string
	for _, response := range council.HashResponses {
		allreadyVoted = append(allreadyVoted, response.User)
	}

	if slices.Contains(allreadyVoted, msg.Creator) {
		return nil, errorsmod.Wrap(errors.ErrUnauthorized, "Allready voted")
	}

	resp := types.WrapHashResponse{
		User: msg.Creator,
		Hash: msg.Response,
	}
	council.HashResponses = append(council.HashResponses, &resp)
	if msg.Suggestion != "" { // Direcly reveal when a suggestion is made
		clearResp := types.WrapClearResponse{
			User:       msg.Creator,
			Response:   types.Response_Suggestion,
			Suggestion: msg.Suggestion,
		}
		council.ClearResponses = append(council.ClearResponses, &clearResp)
	}

	if len(council.HashResponses) == 5 {
		council.Status = types.CouncelingStatus_commited
	}

	err = k.BurnCoinsFromAddr(ctx, creator.Addr, sdk.Coins{collateralDeposit})
	if err != nil {
		return nil, errorsmod.Wrap(errors.ErrInsufficientFunds, "Voter does not have enough coins")
	}
	council.Treasury = council.Treasury.Add(collateralDeposit)

	err = k.TryEvaluate(ctx, council)
	if err != nil {
		return nil, err
	}

	k.councils.Set(ctx, msg.CouncilId, council)

	return &types.MsgCouncilResponseCommitResponse{}, nil
}
