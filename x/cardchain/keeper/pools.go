package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Keys to access pools
const (
	PublicPoolKey    = "public"
	WinnersPoolKey   = "winners"
	BalancersPoolKey = "balancers"
)

// AddPoolCredits Adds ucredits to a pool
func (k Keeper) AddPoolCredits(ctx sdk.Context, poolName string, amount sdk.Coin) {
	pool := k.Pools.Get(ctx, poolName)
	newPool := pool.Add(amount)
	k.Pools.Set(ctx, poolName, &newPool)
}

// SubPoolCredits Subtracts ucredits from a pool
func (k Keeper) SubPoolCredits(ctx sdk.Context, poolName string, amount sdk.Coin) {
	pool := k.Pools.Get(ctx, poolName)
	newPool := pool.Sub(amount)
	k.Pools.Set(ctx, poolName, &newPool)
}
