package keeper

import (
	"context"
	"time"

	sdkerrors "cosmossdk.io/errors"
	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) ConfirmMatch(goCtx context.Context, msg *types.MsgConfirmMatch) (*types.MsgConfirmMatchResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var (
		player *types.MatchPlayer
	)

	match := k.Matches.Get(ctx, msg.MatchId)

	switch msg.Creator {
	case match.PlayerA.Addr:
		player = match.PlayerA
	case match.PlayerB.Addr:
		player = match.PlayerB
	default:
		return nil, sdkerrors.Wrap(errors.ErrUnauthorized, "Didn't participate in match")
	}

	if player.Confirmed {
		return nil, sdkerrors.Wrap(errors.ErrUnauthorized, "Already reported")
	}
	if match.Outcome == types.Outcome_Aborted {
		return nil, sdkerrors.Wrap(errors.ErrUnauthorized, "Can't report, because match was aborted")
	}

	player.Outcome = msg.Outcome
	player.VotedCards = msg.VotedCards
	player.Confirmed = true
	match.Timestamp = uint64(time.Now().Unix())

	err := k.TryHandleMatchOutcome(ctx, match)
	if err != nil {
		return nil, err
	}

	k.Matches.Set(ctx, msg.MatchId, match)

	return &types.MsgConfirmMatchResponse{}, nil
}
