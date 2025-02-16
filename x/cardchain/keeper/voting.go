package keeper

import (
	"fmt"

	sdkerrors "cosmossdk.io/errors"
	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) voteCard(
	ctx sdk.Context,
	voter *User,
	cardId uint64,
	voteType types.VoteType,
) error {
	// if the vote right is valid, get the Card
	card := k.Cards.Get(ctx, cardId)

	// check if card status is valid
	if card.Status != types.CardStatus_permanent && card.Status != types.CardStatus_trial {
		return sdkerrors.Wrap(errors.ErrUnknownRequest, "Voting on a card is only possible if it is in trial or a permanent card")
	}

	err := k.RegisterVote(voter, cardId)
	if err != nil {
		return err
	}

	switch voteType {
	case types.VoteType_fairEnough:
		card.FairEnoughVotes++
	case types.VoteType_inappropriate:
		card.InappropriateVotes++
	case types.VoteType_overpowered:
		card.OverpoweredVotes++
	case types.VoteType_underpowered:
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
	err = k.MintCoinsToAddr(ctx, voter.Addr, sdk.Coins{amount})
	if err != nil {
		return err
	}
	k.SubPoolCredits(ctx, BalancersPoolKey, amount)
	k.Cards.Set(ctx, cardId, card)

	return nil
}

func (k Keeper) incVotesAverageBy(ctx sdk.Context, n int64) {
	votes := k.RunningAverages.Get(ctx, Votes24ValueKey)
	votes.Arr[len(votes.Arr)-1] += n
	k.RunningAverages.Set(ctx, Votes24ValueKey, votes)
}

func (k Keeper) multiVote(ctx sdk.Context, voter *User, votes []*types.SingleVote) error {
	var n = 0

	for i, vote := range votes {
		err := k.voteCard(ctx, voter, vote.CardId, vote.VoteType)
		if err != nil {
			return err
		}
		n = i
	}

	k.incVotesAverageBy(ctx, int64(n+1))

	return nil
}
