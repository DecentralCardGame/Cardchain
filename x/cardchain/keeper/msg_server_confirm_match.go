package keeper

import (
	"context"

	"golang.org/x/exp/slices"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) ConfirmMatch(goCtx context.Context, msg *types.MsgConfirmMatch) (*types.MsgConfirmMatchResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var (
		player      *types.MatchPlayer
		otherPlayer *types.MatchPlayer
	)

	match := k.Matches.Get(ctx, msg.MatchId)

	switch msg.Creator {
	case match.PlayerA.Addr:
		player = match.PlayerA
		otherPlayer = match.PlayerB
	case match.PlayerB.Addr:
		player = match.PlayerB
		otherPlayer = match.PlayerA
	default:
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Didn't participate in match")
	}

	if player.Confirmed {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Already reported")
	}
	if match.Outcome == types.Outcome_Aborted {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Can't report, because match was aborted")
	}

	player.Outcome = msg.Outcome

	// Prefilter voted cards cards
	var votedCards []uint64
	for _, card := range msg.VotedCards {
		if slices.Contains(otherPlayer.Deck, card) {
			votedCards = append(votedCards, card)
		}
	}

	player.VotedCards = votedCards
	player.Confirmed = true
	k.Matches.Set(ctx, msg.MatchId, match)

	return &types.MsgConfirmMatchResponse{}, nil
}
