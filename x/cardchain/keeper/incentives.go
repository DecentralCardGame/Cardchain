package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Keys to access RunningAverages
const (
	Games24ValueKey = "games/24h"
	Votes24ValueKey = "votes/24h"
)

// GetWinnerIncentives calculates incentives for the winners pool
func (k Keeper) GetWinnerIncentives(ctx sdk.Context) float32 {
	games := float32(k.GetGames(ctx))
	votes := float32(k.GetVotes(ctx))
	gVR := float32(k.GetParams(ctx).GameVoteRatio) / 100
	return games / (votes*gVR + games)
}

// GetBalancerIncentives calculates incentives for the ballancers pool
func (k Keeper) GetBalancerIncentives(ctx sdk.Context) float32 {
	games := float32(k.GetGames(ctx))
	votes := float32(k.GetVotes(ctx))
	gVR := float32(k.GetParams(ctx).GameVoteRatio) / 100
	return (votes * gVR) / (votes*gVR + games)
}

// GetRunningAverageTotal Sums up a certain running average
func (k Keeper) GetRunningAverageTotal(ctx sdk.Context, key string) (num int64) {
	runningAverage := k.RunningAverages.Get(ctx, key)
	for _, val := range runningAverage.Arr {
		num += val
	}
	if num == 0 {
		num = 1
	}
	return
}

// GetGames Gets the number of games played in the last 24 hours
func (k Keeper) GetGames(ctx sdk.Context) int64 {
	return k.GetRunningAverageTotal(ctx, Games24ValueKey)
}

// GetVotes Gets the number of votes made in the last 24 hours
func (k Keeper) GetVotes(ctx sdk.Context) int64 {
	return k.GetRunningAverageTotal(ctx, Votes24ValueKey)
}
