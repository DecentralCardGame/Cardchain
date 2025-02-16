package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) MatchConfirm(goCtx context.Context, msg *types.MsgMatchConfirm) (*types.MsgMatchConfirmResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var (
		player *types.MatchPlayer
	)

	match := k.matches.Get(ctx, msg.MatchId)

	switch msg.Creator {
	case match.PlayerA.Addr:
		player = match.PlayerA
	case match.PlayerB.Addr:
		player = match.PlayerB
	default:
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "Didn't participate in match")
	}

	if player.Confirmed {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "Already reported")
	}
	if match.Outcome == types.Outcome_Aborted {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "Can't report, because match was aborted")
	}

	player.Outcome = msg.Outcome
	player.VotedCards = msg.VotedCards
	player.Confirmed = true
	match.Timestamp = uint64(ctx.BlockHeight())

	err := k.TryHandleMatchOutcome(ctx, match)
	if err != nil {
		return nil, err
	}

	k.matches.Set(ctx, msg.MatchId, match)

	return &types.MsgMatchConfirmResponse{}, nil
}
