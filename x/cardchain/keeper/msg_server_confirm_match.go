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
		outcome types.Outcome
		playerOutcome *types.Outcome
	)

	match := k.Matches.Get(ctx, msg.MatchId)

	switch msg.Creator {
	case match.PlayerA:
		playerOutcome = &match.OutcomeA
	case match.PlayerB:
		playerOutcome = &match.OutcomeB
	default:
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Didn't participate in match")
	}

	if *playerOutcome != types.Outcome_Aborted {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Already reported")
	}

	if match.Outcome == types.Outcome_Aborted {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Can't report, because match was aborted")
	}

	*playerOutcome = msg.Outcome
	k.Matches.Set(ctx, msg.MatchId, match)

	// majority based vote
	outcomes := []types.Outcome{match.Outcome, match.OutcomeA, match.OutcomeB}
	slices.Sort(outcomes)
	if outcomes[0] == outcomes[1] {
		outcome = outcomes[0]
	} else if outcomes[2] == outcomes[1] {
		outcome = outcomes[3]
	} else {
		outcome = types.Outcome_Aborted
		return &types.MsgConfirmMatchResponse{}, nil
	}

	if match.CoinsDistributed {  // Ensures that money isn't dropped twice
		return &types.MsgConfirmMatchResponse{}, nil
	}

	addresses, err := k.GetMatchAddresses(ctx, *match)
	if err != nil {
		return nil, err
	}

	amountA, amountB := k.CalculateMatchReward(ctx, outcome)
	amounts := []sdk.Coin{amountA, amountB}
	for idx := range addresses {
		err := k.MintCoinsToAddr(ctx, addresses[idx], sdk.Coins{amounts[idx]})
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, err.Error())
		}
		k.SubPoolCredits(ctx, WinnersPoolKey, amounts[idx])
	}

	games := k.RunningAverages.Get(ctx, Games24ValueKey)
	games.Arr[len(games.Arr)-1]++
	k.RunningAverages.Set(ctx, Games24ValueKey, games)

	match.CoinsDistributed = true

	k.Matches.Set(ctx, msg.MatchId, match)

	return &types.MsgConfirmMatchResponse{}, nil
}
