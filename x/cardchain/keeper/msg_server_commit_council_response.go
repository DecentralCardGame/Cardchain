package keeper

import (
	"context"
	"slices"

	sdkerrors "cosmossdk.io/errors"
	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CommitCouncilResponse(goCtx context.Context, msg *types.MsgCommitCouncilResponse) (*types.MsgCommitCouncilResponseResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	collateralDeposit := k.GetParams(ctx).CollateralDeposit

	creator, err := k.GetUserFromString(ctx, msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrUserDoesNotExist, err.Error())
	}

	council := k.Councils.Get(ctx, msg.CouncilId)
	if !slices.Contains(council.Voters, msg.Creator) {
		return nil, sdkerrors.Wrap(errors.ErrUnauthorized, "Invalid Voter")
	}

	if council.Status != types.CouncelingStatus_councilCreated {
		return nil, sdkerrors.Wrapf(types.ErrCouncilStatus, "Have '%s', want '%s'", council.Status.String(), types.CouncelingStatus_councilCreated.String())
	}

	var allreadyVoted []string
	for _, response := range council.HashResponses {
		allreadyVoted = append(allreadyVoted, response.User)
	}

	if slices.Contains(allreadyVoted, msg.Creator) {
		return nil, sdkerrors.Wrap(errors.ErrUnauthorized, "Allready voted")
	}

	commitCouncilResponse(council, creator, msg.Response, msg.Suggestion)

	err = k.BurnCoinsFromAddr(ctx, creator.Addr, sdk.Coins{collateralDeposit})
	if err != nil {
		return nil, sdkerrors.Wrap(errors.ErrInsufficientFunds, "Voter does not have enough coins")
	}
	council.Treasury = council.Treasury.Add(collateralDeposit)

	err = k.TryEvaluate(ctx, council)
	if err != nil {
		return nil, err
	}

	k.Councils.Set(ctx, msg.CouncilId, council)

	return &types.MsgCommitCouncilResponseResponse{}, nil
}
