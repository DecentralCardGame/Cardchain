package keeper

import (
	"slices"

	sdkerrors "cosmossdk.io/errors"
	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetVoteReward Calculates winner rewards
func (k Keeper) GetVoteReward(ctx sdk.Context) sdk.Coin {
	params := k.GetParams(ctx)

	pool := k.Pools.Get(ctx, BalancersPoolKey)
	reward := QuoCoin(*pool, params.VotePoolFraction)
	if reward.Amount.Int64() > params.VotingRewardCap {
		return sdk.NewInt64Coin(reward.Denom, params.VotingRewardCap)
	}
	return reward
}

// AddVoteRightToUser Adds a voteright for a certain card to a certain user
func (k Keeper) AddVoteRightToUser(ctx sdk.Context, user *types.User, cardId uint64) {
	card := k.Cards.Get(ctx, cardId)
	if !card.BalanceAnchor {
		if !slices.Contains(user.VotableCards, cardId) && !slices.Contains(user.VotedCards, cardId) {
			user.VotableCards = append(user.VotableCards, cardId)
		}
	}
}

// AddVoteRightsToAllUsers Adds voterights to all cards to all users
func (k Keeper) AddVoteRightsToAllUsers(ctx sdk.Context) {
	votables := k.GetAllVotableCards(ctx)
	allUsers, allAddrs := k.GetAllUsers(ctx)
	for idx, user := range allUsers {
		user.VotableCards = votables
		k.SetUser(ctx, allAddrs[idx], *user)
	}
}

// RegisterVote Removes a voteright from a user
func (k Keeper) RegisterVote(voter *User, cardId uint64) error {
	var rightsIndex = slices.Index(voter.VotableCards, cardId)

	// check if voting rights are true
	if rightsIndex < 0 {
		return sdkerrors.Wrap(types.ErrVoterHasNoVotingRights, "No Voting Rights")
	}
	voter.VotableCards = slices.Delete(voter.VotableCards, rightsIndex, rightsIndex+1)
	voter.VotedCards = append(voter.VotedCards, cardId)
	return nil
}

// GetAllVotableCards Gets the voterights to all cards
func (k Keeper) GetAllVotableCards(ctx sdk.Context) (votables []uint64) {
	iter := k.Cards.GetItemIterator(ctx)

	for ; iter.Valid(); iter.Next() {
		// here only give right if card is not a scheme or banished
		idx, gottenCard := iter.Value()
		if !gottenCard.BalanceAnchor {
			switch gottenCard.Status {
			case types.Status_permanent, types.Status_trial:
				votables = append(votables, idx)
			}
		}
	}
	return
}

// RemoveExpiredVoteRights Removes all expied voteRights
func (k Keeper) RemoveExpiredVoteRights(ctx sdk.Context) {
	allUsers, allAddrs := k.GetAllUsers(ctx)

	for idx, user := range allUsers {
		user.VotableCards = []uint64{}
		user.VotedCards = []uint64{}
		k.SetUser(ctx, allAddrs[idx], *user)
	}
}

// ResetAllVotes Resets all cards votes
func (k Keeper) ResetAllVotes(ctx sdk.Context) {
	iter := k.Cards.GetItemIterator(ctx)

	for ; iter.Valid(); iter.Next() {
		idx, resetCard := iter.Value()
		if resetCard.Status != types.Status_trial {
			resetCard.ResetVotes()
		}
		k.Cards.Set(ctx, uint64(idx), resetCard)
	}
}
