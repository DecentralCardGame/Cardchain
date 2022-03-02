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
	games := float32(k.GetGames(ctx))
	votes := float32(k.GetVotes(ctx))
	gVR := float32(k.GetParams(ctx).GameVoteRatio) / 100
	return games / (votes*gVR + games)
}

func (k Keeper) GetBalancerIncentives(ctx sdk.Context) float32 {
	games := float32(k.GetGames(ctx))
	votes := float32(k.GetVotes(ctx))
	gVR := float32(k.GetParams(ctx).GameVoteRatio) / 100
	return (votes * gVR) / (votes*gVR + games)
}

func (k Keeper) GetGames(ctx sdk.Context) (num int64) {
	games := k.GetRunningAverage(ctx, Games24ValueKey)
  for _, val := range games.Arr {
    num += val
  }
  if num == 0 {
    num = 1
  }
  return
}

func (k Keeper) GetVotes(ctx sdk.Context) (num int64) {
	votes := k.GetRunningAverage(ctx, Votes24ValueKey)
  for _, val := range votes.Arr {
    num += val
  }
  if num == 0 {
    num = 1
  }
  return
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
func (k Keeper) GetAllRunningAverages(ctx sdk.Context) (allRunningAverages []*types.RunningAverage) {
	for _, runningAverageName := range k.RunningAverageKeys {
    gottenAverage := k.GetRunningAverage(ctx, runningAverageName)
		allRunningAverages = append(allRunningAverages, &gottenAverage)
	}
	return
}
