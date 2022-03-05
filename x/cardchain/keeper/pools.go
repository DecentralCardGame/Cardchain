package keeper

import (
	"fmt"

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
	pool := k.GetPool(ctx, poolName)
	pool = pool.Add(amount)
	k.SetPool(ctx, poolName, pool)
}

// SubPoolCredits Subtracts ucredits from a pool
func (k Keeper) SubPoolCredits(ctx sdk.Context, poolName string, amount sdk.Coin) {
	pool := k.GetPool(ctx, poolName)
	pool = pool.Sub(amount)
	k.SetPool(ctx, poolName, pool)
}

// GetPool Returns a given pool
func (k Keeper) GetPool(ctx sdk.Context, poolName string) sdk.Coin {
	store := ctx.KVStore(k.PoolsStoreKey)
	bz := store.Get([]byte(poolName))

	var gottenPool sdk.Coin
	k.cdc.MustUnmarshal(bz, &gottenPool)
	return gottenPool
}

// SetPool Sets a given pool
func (k Keeper) SetPool(ctx sdk.Context, poolName string, newPool sdk.Coin) {
	store := ctx.KVStore(k.PoolsStoreKey)
	if newPool.Amount.Int64() < 0 {
		k.Logger(ctx).Error(fmt.Sprintf(":: Pool amount negativ!: %s: %d", poolName, newPool.Amount.Int64()))
	}
	store.Set([]byte(poolName), k.cdc.MustMarshal(&newPool))
}

// GetAllPools Returns all pools
func (k Keeper) GetAllPools(ctx sdk.Context) []sdk.Coin {
	var allPools []sdk.Coin
	for _, poolName := range k.PoolKeys {
		allPools = append(allPools, k.GetPool(ctx, poolName))
	}
	return allPools
}
