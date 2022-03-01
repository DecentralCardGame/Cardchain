package keeper

import (
  "github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	Games24ValueKey = "games/24h"
	Votes24ValueKey = "votes/24h"
)

func (k Keeper) GetWinnerIncentives(ctx sdk.Context) float32 {
	games := float32(k.GetGames())
	votes := float32(k.GetVotes())
	gVR := float32(k.GetParams(ctx).GameVoteRatio) / 100
	if games == 0 || votes == 0 {
		games = 1
		votes = 1
	}
	return games / (votes*gVR + games)
}

func (k Keeper) GetBalancerIncentives(ctx sdk.Context) float32 {
	games := float32(k.GetGames())
	votes := float32(k.GetVotes())
	gVR := float32(k.GetParams(ctx).GameVoteRatio) / 100
	if games == 0 || votes == 0 {
		games = 1
		votes = 1
	}
	return (votes * gVR) / (votes*gVR + games)
}

func (k Keeper) GetGames() int64 {
	return 0
}

func (k Keeper) GetVotes() int64 {
	return 0
}

// GetRunningAverage Returns a given runningAverage
func (k Keeper) GetRunningAverage(ctx sdk.Context, runningAverageName string) types.RunningAverage {
	store := ctx.KVStore(k.RunningAveragesStoreKey)
	bz := store.Get([]byte(runningAverageName))

	var gottenRunningAverage types.RunningAverage
	k.cdc.MustUnmarshal(bz, &gottenRunningAverage)
	return gottenRunningAverage
}

// SetRunningAverage Sets a given runningAverage
func (k Keeper) SetRunningAverage(ctx sdk.Context, runningAverageName string, newRunningAverage types.RunningAverage) {
	store := ctx.KVStore(k.RunningAveragesStoreKey)
	store.Set([]byte(runningAverageName), k.cdc.MustMarshal(&newRunningAverage))
}

// GetAllRunningAverages Returns all runningAverages
func (k Keeper) GetAllRunningAverages(ctx sdk.Context) (allRunningAverages []types.RunningAverage) {
	for _, runningAverageName := range k.RunningAverageKeys {
		allRunningAverages = append(allRunningAverages, k.GetRunningAverage(ctx, runningAverageName))
	}
	return
}
