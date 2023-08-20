package keeper

import (
	"context"

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

	err = k.TryHandleMatchOutcome(ctx, match)
	if err != nil {
		return nil, err
	}

	k.Matches.Set(ctx, msg.MatchId, match)

	return &types.MsgReportMatchResponse{}, nil
}
