package keeper

import (
	"context"

	"golang.org/x/exp/slices"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) ReportMatch(goCtx context.Context, msg *types.MsgReportMatch) (*types.MsgReportMatchResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	match := k.Matches.Get(ctx, msg.MatchId)

	creator, err := k.GetUserFromString(ctx, msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrUserDoesNotExist, err.Error())
	}
	if creator.ReportMatches == false {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Reporter")
	}
	if msg.Creator != match.Reporter {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "Wrong Reporter, reporter is %s", match.Reporter)
	}
	if !match.PlayerA.Confirmed && !match.PlayerB.Confirmed {
		return nil, sdkerrors.Wrapf(types.ErrWaitingForPlayers, "Waiting for players to report")
	}
	if !match.CoinsDistributed {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Match already reported")
	}

	match.Outcome = msg.Outcome

	// Evaluate Outcome
	outcomes := []types.Outcome{msg.Outcome, match.PlayerA.Outcome, match.PlayerB.Outcome}
	slices.Sort(outcomes)
	outcomes = slices.Compact(outcomes)
	switch i := uint64(len(outcomes)); i {
	case 1:
		k.ReportServerMatch(ctx, match.Reporter, 1, true)
	default:
		k.ReportServerMatch(ctx, match.Reporter, i-1, false)
	}

	outcome, err := k.GetOutcome(ctx, *match)
	match.Outcome = outcome

	err = k.DistributeCoins(ctx, match, outcome)
	if err != nil {
		return nil, err
	}

	// TODO: Votes

	k.Matches.Set(ctx, msg.MatchId, match)

	return &types.MsgReportMatchResponse{}, nil
}
