package keeper

import (
  "encoding/binary"

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

	userStore := ctx.KVStore(k.UsersStoreKey)

	userIterator := sdk.KVStorePrefixIterator(userStore, nil)

	for ; userIterator.Valid(); userIterator.Next() {
		var user types.User
		k.cdc.MustUnmarshal(userIterator.Value(), &user)
		user.VoteRights = votingRights
		userStore.Set(userIterator.Key(), k.cdc.MustMarshal(&user))
	}

	userIterator.Close()
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
	cardStore := ctx.KVStore(k.CardsStoreKey)
	cardIterator := sdk.KVStorePrefixIterator(cardStore, nil)

	for ; cardIterator.Valid(); cardIterator.Next() {
		// here only give right if card is not a scheme or banished
		var gottenCard types.Card
		k.cdc.MustUnmarshal(cardIterator.Value(), &gottenCard)

		if gottenCard.Status == types.Status_permanent || gottenCard.Status == types.Status_trial || gottenCard.Status == types.Status_prototype {
			right := types.NewVoteRight(binary.BigEndian.Uint64(cardIterator.Key()), expireBlock)
			votingRights = append(votingRights, &right)
		}
	}
	cardIterator.Close()

	return
}

// GetVoteRights Gets the voterights of a user
func (k Keeper) GetVoteRights(ctx sdk.Context, voter sdk.AccAddress) []*types.VoteRight {
	user, _ := k.GetUser(ctx, voter)
	return user.VoteRights
}

// ResetAllVotes Resets all cards votes
func (k Keeper) ResetAllVotes(ctx sdk.Context) {
	store := ctx.KVStore(k.CardsStoreKey)
	iterator := sdk.KVStorePrefixIterator(store, nil)

	for ; iterator.Valid(); iterator.Next() {
		var resetCard types.Card
		k.cdc.MustUnmarshal(iterator.Value(), &resetCard)
		if resetCard.Status != types.Status_trial {
			resetCard.ResetVotes()
		}

		store.Set(iterator.Key(), k.cdc.MustMarshal(&resetCard))
	}
}

// SearchVoteRights Searches voterights for a certain card
func SearchVoteRights(cardID uint64, rights []*types.VoteRight) int {
	if rights == nil {
		return -1
	}
	for i, b := range rights {
		if b.CardId == cardID {
			return i
		}
	}
	return -1
}
