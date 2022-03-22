package keeper

import (
	"context"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) ConfirmMatch(goCtx context.Context, msg *types.MsgConfirmMatch) (*types.MsgConfirmMatchResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var (
		player  *types.MatchPlayer
	)

	match := k.Matches.Get(ctx, msg.MatchId)

	switch msg.Creator {
	case match.PlayerA.Addr:
		player = match.PlayerA
	case match.PlayerB.Addr:
		player = match.PlayerB
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
	player.Confirmed = true
	k.Matches.Set(ctx, msg.MatchId, match)

	outcome, err := k.GetOutcome(ctx, *match)

	// Distribute coins
	if match.CoinsDistributed || err != nil {  // Ensures that money isn't dropped twice
		return &types.MsgConfirmMatchResponse{}, nil
	}

	err = k.DistributeCoins(ctx, match, outcome)
	if err != nil {
		return nil, err
	}

	k.Matches.Set(ctx, msg.MatchId, match)

	return &types.MsgConfirmMatchResponse{}, nil
}
