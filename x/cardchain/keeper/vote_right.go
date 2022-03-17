package keeper

import (
	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetVoteReward Calculates winner rewards
func (k Keeper) GetVoteReward(ctx sdk.Context) sdk.Coin {
	pool := k.GetPool(ctx, BalancersPoolKey)
	reward := QuoCoin(pool, k.GetParams(ctx).VoterReward)
	if reward.Amount.Int64() > 1000000 {
		return sdk.NewInt64Coin(reward.Denom, 1000000)
	}
	return reward
}

// AddVoteRight Adds a voteright for a certain card to a certain user
func (k Keeper) AddVoteRight(ctx sdk.Context, userAddress sdk.AccAddress, cardId uint64) error {
	user, err := k.GetUser(ctx, userAddress)
	if err != nil {
		return err
	}
	right := types.NewVoteRight(cardId, ctx.BlockHeight()+k.GetParams(ctx).VotingRightsExpirationTime)
	user.VoteRights = append(user.VoteRights, &right)
	k.SetUser(ctx, userAddress, user)
	return nil
}

// AddVoteRightsToAllUsers Adds voterights to all cards to all users
func (k Keeper) AddVoteRightsToAllUsers(ctx sdk.Context, expireBlock int64) {
	votingRights := k.GetVoteRightToAllCards(ctx, expireBlock)
	allUsers, allAddrs := k.GetAllUsers(ctx)

	for idx, user := range allUsers {
		user.VoteRights = votingRights
		k.SetUser(ctx, allAddrs[idx], *user)
	}
}

// RemoveVoteRight Removes a voteright from a user
func (k Keeper) RemoveVoteRight(ctx sdk.Context, userAddress sdk.AccAddress, rightsIndex int) error {
	user, err := k.GetUser(ctx, userAddress)
	if err != nil {
		return err
	}
	user.VoteRights[rightsIndex] = user.VoteRights[len(user.VoteRights)-1]
	//user.VoteRights[len(user.VoteRights)-1] = null
	user.VoteRights = user.VoteRights[:len(user.VoteRights)-1]
	k.SetUser(ctx, userAddress, user)
	return nil
}

// GetVoteRightToAllCards Gets the voterights to all cards
func (k Keeper) GetVoteRightToAllCards(ctx sdk.Context, expireBlock int64) (votingRights []*types.VoteRight) {
	allCards := k.Card.GetAll(ctx)

	for idx, gottenCard := range allCards {
		// here only give right if card is not a scheme or banished
		if gottenCard.Status == types.Status_permanent || gottenCard.Status == types.Status_trial || gottenCard.Status == types.Status_prototype {
			right := types.NewVoteRight(uint64(idx), expireBlock)
			votingRights = append(votingRights, &right)
		}
	}
	return
}

// GetVoteRights Gets the voterights of a user
func (k Keeper) GetVoteRights(ctx sdk.Context, voter sdk.AccAddress) []*types.VoteRight {
	user, _ := k.GetUser(ctx, voter)
	return user.VoteRights
}

// ResetAllVotes Resets all cards votes
func (k Keeper) ResetAllVotes(ctx sdk.Context) {
	allCards := k.Card.GetAll(ctx)

	for idx, resetCard := range allCards {
		if resetCard.Status != types.Status_trial {
			resetCard.ResetVotes()
		}
		k.Card.Set(ctx, uint64(idx), resetCard)
	}
}
