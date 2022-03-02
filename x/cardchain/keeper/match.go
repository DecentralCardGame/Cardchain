package keeper

import (
  "github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// ApointMatchReporter Makes a user a match reporter
func (k Keeper) ApointMatchReporter(ctx sdk.Context, address string) error {
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
	reward := k.GetParams(ctx).WinnerReward // TODO make a fraction
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
	return
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
