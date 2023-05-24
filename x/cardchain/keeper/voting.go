package keeper

import (
	"fmt"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"golang.org/x/exp/slices"
)

func (k msgServer) voteCard(
	ctx sdk.Context,
	voter *User,
	voteRights *[]*types.VoteRight,
	cardId uint64,
	voteType string,
) error {
	rightsIndex := slices.IndexFunc(*voteRights, func(s *types.VoteRight) bool { return s.CardId == cardId })

	// check if voting rights are true
	if rightsIndex < 0 {
		return sdkerrors.Wrap(types.ErrVoterHasNoVotingRights, "No Voting Rights")
	}

	//check if voting rights are timed out
	if ctx.BlockHeight() > (*voteRights)[rightsIndex].ExpireBlock {
		k.RemoveVoteRight(ctx, voter, rightsIndex)
		return sdkerrors.Wrap(types.ErrVoteRightIsExpired, "Voting Right has expired")
	}

	// if the vote right is valid, get the Card
	card := k.Cards.Get(ctx, cardId)

	// check if card status is valid
	if card.Status != types.Status_permanent && card.Status != types.Status_trial {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Voting on a card is only possible if it is in trial or a permanent card")
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
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
	}
	card.Voters = append(card.Voters, voter.Addr.String())

	// check for specific bounty on the card
	if !card.VotePool.IsZero() {
		reward := k.GetParams(ctx).TrialVoteReward
		err := k.MintCoinsToAddr(ctx, voter.Addr, sdk.Coins{reward})
		if err != nil {
			return sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, err.Error())
		}
		card.VotePool.Sub(reward) // TODO actually if there is less than 1cr then it should be adjusted
	}

	amount := k.GetVoteReward(ctx)
	k.MintCoinsToAddr(ctx, voter.Addr, sdk.Coins{amount})
	k.SubPoolCredits(ctx, BalancersPoolKey, amount)
	k.RemoveVoteRight(ctx, voter, rightsIndex)

	*voteRights = append((*voteRights)[:rightsIndex], (*voteRights)[rightsIndex+1:]...)

	k.Cards.Set(ctx, cardId, card)

	return nil
}

func (k msgServer) incVotesAverageBy(ctx sdk.Context, n int64) {
	votes := k.RunningAverages.Get(ctx, Votes24ValueKey)
	votes.Arr[len(votes.Arr)-1] += n
	k.RunningAverages.Set(ctx, Votes24ValueKey, votes)
}
