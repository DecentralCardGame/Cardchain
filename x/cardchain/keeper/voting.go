package keeper

import (
	"fmt"
	"slices"

	sdkerrors "cosmossdk.io/errors"
	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) voteCard(
	ctx sdk.Context,
	voter *User,
	cardId uint64,
	voteType string,
	ignoreVoteRights bool,
) error {
	var rightsIndex int
	if !ignoreVoteRights {
		rightsIndex = slices.IndexFunc(voter.VoteRights, func(s *types.VoteRight) bool { return s.CardId == cardId })

		// check if voting rights are true
		if rightsIndex < 0 {
			return sdkerrors.Wrap(types.ErrVoterHasNoVotingRights, "No Voting Rights")
		}

		//check if voting rights are timed out
		if ctx.BlockHeight() > (voter.VoteRights)[rightsIndex].ExpireBlock {
			k.RemoveVoteRight(ctx, voter, rightsIndex)
			return sdkerrors.Wrap(types.ErrVoteRightIsExpired, "Voting Right has expired")
		}
	}

	// if the vote right is valid, get the Card
	card := k.Cards.Get(ctx, cardId)

	// check if card status is valid
	if card.Status != types.Status_permanent && card.Status != types.Status_trial {
		return sdkerrors.Wrap(errors.ErrUnknownRequest, "Voting on a card is only possible if it is in trial or a permanent card")
	}

	switch voteType {
	case "fair_enough":
		card.FairEnoughVotes++
	case "inappropriate":
		card.InappropriateVotes++
	case "overpowered":
		card.OverpoweredVotes++
	case "underpowered":
		card.UnderpoweredVotes++
	default:
		errMsg := fmt.Sprintf("Unrecognized card vote type: %s", voteType)
		return sdkerrors.Wrap(errors.ErrUnknownRequest, errMsg)
	}
	card.Voters = append(card.Voters, voter.Addr.String())

	// check for specific bounty on the card
	if !card.VotePool.IsZero() {
		reward := k.GetParams(ctx).TrialVoteReward
		err := k.MintCoinsToAddr(ctx, voter.Addr, sdk.Coins{reward})
		if err != nil {
			return sdkerrors.Wrap(errors.ErrInsufficientFunds, err.Error())
		}
		card.VotePool.Sub(reward) // TODO actually if there is less than 1cr then it should be adjusted
	}

	amount := k.GetVoteReward(ctx)
	err := k.MintCoinsToAddr(ctx, voter.Addr, sdk.Coins{amount})
	if err != nil {
		return err
	}
	k.SubPoolCredits(ctx, BalancersPoolKey, amount)
	if !ignoreVoteRights {
		k.RemoveVoteRight(ctx, voter, rightsIndex)
	}

	k.Cards.Set(ctx, cardId, card)

	return nil
}

func (k Keeper) incVotesAverageBy(ctx sdk.Context, n int64) {
	votes := k.RunningAverages.Get(ctx, Votes24ValueKey)
	votes.Arr[len(votes.Arr)-1] += n
	k.RunningAverages.Set(ctx, Votes24ValueKey, votes)
}

func (k Keeper) multiVote(ctx sdk.Context, voter *User, votes []*types.SingleVote, ignoreVotingRights bool) error {
	var n = 0

	for i, vote := range votes {
		err := k.voteCard(ctx, voter, vote.CardId, vote.VoteType, ignoreVotingRights)
		if err != nil {
			return err
		}
		n = i
	}

	k.incVotesAverageBy(ctx, int64(n+1))

	return nil
}
