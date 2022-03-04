package keeper

import (
  "github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetMatchReporter Makes a user a match reporter
func (k Keeper) SetMatchReporter(ctx sdk.Context, address string) error {
	reporter, err := k.GetUserFromString(ctx, address)
	if err != nil {
		return err
	}
	reporter.ReportMatches = true

	k.SetUserFromUser(ctx, reporter)
	return nil
}

// CalculateMatchReward Calculates the match winning rewards
func (k Keeper) CalculateMatchReward(ctx sdk.Context, outcome types.Outcome) (amountA sdk.Coin, amountB sdk.Coin) {
	reward := k.GetMatchReward(ctx)
	amountA = sdk.NewInt64Coin("ucredits", 0)
	amountB = sdk.NewInt64Coin("ucredits", 0)

	if outcome == types.Outcome_AWon {
		amountA = reward
	} else if outcome == types.Outcome_BWon {
		amountB = reward
	} else if outcome == types.Outcome_Draw {
		amountA = QuoCoin(reward, 2)
		amountB = QuoCoin(reward, 2)
	}
  if outcome != types.Outcome_Aborted {
    k.SubPoolCredits(ctx, WinnersPoolKey, reward)
  }
	return
}

// GetMatchReward Calculates winner rewards
func (k Keeper) GetMatchReward(ctx sdk.Context) sdk.Coin {
  pool := k.GetPool(ctx, WinnersPoolKey)
  reward := QuoCoin(pool, k.GetParams(ctx).WinnerReward)
  if reward.Amount.Int64() > 1000000 {
    return sdk.NewInt64Coin(reward.Denom, 1000000)
  }
  return reward
}

// GetMatch Gets a match from store
func (k Keeper) GetMatch(ctx sdk.Context, matchId uint64) (gottenMatch types.Match) {
	store := ctx.KVStore(k.MatchesStoreKey)
	bz := store.Get(sdk.Uint64ToBigEndian(matchId))

	k.cdc.MustUnmarshal(bz, &gottenMatch)
	return
}

// SetMatch Sets a match in store
func (k Keeper) SetMatch(ctx sdk.Context, matchId uint64, newMatch types.Match) {
	store := ctx.KVStore(k.MatchesStoreKey)
	store.Set(sdk.Uint64ToBigEndian(matchId), k.cdc.MustMarshal(&newMatch))
}

// GetMatchesIterator Returns an iterator for all matches
func (k Keeper) GetMatchesIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.MatchesStoreKey)
	return sdk.KVStorePrefixIterator(store, nil)
}

// GetAllMatches Gets all matches from store
func (k Keeper) GetAllMatches(ctx sdk.Context) (allMatches []*types.Match) {
	iterator := k.GetMatchesIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {

		var gottenMatch types.Match
		k.cdc.MustUnmarshal(iterator.Value(), &gottenMatch)

		allMatches = append(allMatches, &gottenMatch)
	}
	return
}

// GetMatchesNumber Returns the number of all reported matches
func (k Keeper) GetMatchesNumber(ctx sdk.Context) (matchId uint64) {
	iterator := k.GetMatchesIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {
		matchId++
	}
	return
}
