package keeper

import (
	"context"
	"encoding/binary"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// GetProductDetailsCount get the total number of productDetails
func (k Keeper) GetProductDetailsCount(ctx context.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.ProductDetailsCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetProductDetailsCount set the total number of productDetails
func (k Keeper) SetProductDetailsCount(ctx context.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.ProductDetailsCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendProductDetails appends a productDetails in the store with a new id and update the count
func (k Keeper) AppendProductDetails(
	ctx context.Context,
	productDetails types.ProductDetails,
) uint64 {
	// Create the productDetails
	count := k.GetProductDetailsCount(ctx)

	// Set the ID of the appended value
	productDetails.Id = count

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ProductDetailsKey))
	appendedValue := k.cdc.MustMarshal(&productDetails)
	store.Set(GetProductDetailsIDBytes(productDetails.Id), appendedValue)

	// Update productDetails count
	k.SetProductDetailsCount(ctx, count+1)

	return count
}

// SetProductDetails set a specific productDetails in the store
func (k Keeper) SetProductDetails(ctx context.Context, productDetails types.ProductDetails) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ProductDetailsKey))
	b := k.cdc.MustMarshal(&productDetails)
	store.Set(GetProductDetailsIDBytes(productDetails.Id), b)
}

// GetProductDetails returns a productDetails from its id
func (k Keeper) GetProductDetails(ctx context.Context, id uint64) (val types.ProductDetails, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ProductDetailsKey))
	b := store.Get(GetProductDetailsIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveProductDetails removes a productDetails from the store
func (k Keeper) RemoveProductDetails(ctx context.Context, id uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ProductDetailsKey))
	store.Delete(GetProductDetailsIDBytes(id))
}

// GetAllProductDetails returns all productDetails
func (k Keeper) GetAllProductDetails(ctx context.Context) (list []types.ProductDetails) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ProductDetailsKey))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ProductDetails
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetProductDetailsIDBytes returns the byte representation of the ID
func GetProductDetailsIDBytes(id uint64) []byte {
	bz := types.KeyPrefix(types.ProductDetailsKey)
	bz = append(bz, []byte("/")...)
	bz = binary.BigEndian.AppendUint64(bz, id)
	return bz
}
