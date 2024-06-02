package keeper

import (
    "fmt"

    "github.com/DecentralCardGame/Cardchain/x/cardchain/types"
    sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetZealy Sets a zealy id
func (k Keeper) SetZealy(ctx sdk.Context, zealyId string, zealy types.Zealy) {
    store := ctx.KVStore(k.zealyStoreKey)
    store.Set([]byte(zealyId), k.cdc.MustMarshal(&zealy))
}

// GetZealy Gets a zealy from store
func (k Keeper) GetZealy(ctx sdk.Context, zealyId string) (zealy types.Zealy, err error) {
    store := ctx.KVStore(k.zealyStoreKey)
    bz := store.Get([]byte(zealyId))

    if bz == nil {
        err = fmt.Errorf("zealyId `%s` not in store", zealyId)
    } else {
        k.cdc.MustUnmarshal(bz, &zealy)
    }
    return
}

// GetZealyIterator Returns an iterator for all zealys
func (k Keeper) GetZealyIterator(ctx sdk.Context) sdk.Iterator {
    store := ctx.KVStore(k.zealyStoreKey)
    return sdk.KVStorePrefixIterator(store, nil)
}

// GetAllZealys Gets all zealys from store
func (k Keeper) GetAllZealys(ctx sdk.Context) (allZealys []*types.Zealy, allZealyIds []string) {
    iterator := k.GetZealyIterator(ctx)
    for ; iterator.Valid(); iterator.Next() {

        var gotten types.Zealy
        k.cdc.MustUnmarshal(iterator.Value(), &gotten)

        allZealys = append(allZealys, &gotten)
        allZealyIds = append(allZealyIds, string(iterator.Key()))
    }
    return
}
