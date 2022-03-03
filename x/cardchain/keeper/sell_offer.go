package keeper

import (
	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetSellOffer Gets selloffer form store
func (k Keeper) GetSellOffer(ctx sdk.Context, sellOfferId uint64) (gottenSellOffer types.SellOffer) {
	store := ctx.KVStore(k.SellOffersStoreKey)
	bz := store.Get(sdk.Uint64ToBigEndian(sellOfferId))

	k.cdc.MustUnmarshal(bz, &gottenSellOffer)
	return
}

// SetSellOffer Sets selloffer in store
func (k Keeper) SetSellOffer(ctx sdk.Context, sellOfferId uint64, newSellOffer types.SellOffer) {
	store := ctx.KVStore(k.SellOffersStoreKey)
	store.Set(sdk.Uint64ToBigEndian(sellOfferId), k.cdc.MustMarshal(&newSellOffer))
}

// GetSellOffersIterator Returns an interator for all selloffers
func (k Keeper) GetSellOffersIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.SellOffersStoreKey)
	return sdk.KVStorePrefixIterator(store, nil)
}

// GetAllSellOffers Gets all selloffers from store
func (k Keeper) GetAllSellOffers(ctx sdk.Context) (allSellOffers []*types.SellOffer) {
	iterator := k.GetSellOffersIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {

		var gottenSellOffer types.SellOffer
		k.cdc.MustUnmarshal(iterator.Value(), &gottenSellOffer)

		allSellOffers = append(allSellOffers, &gottenSellOffer)
	}
	return
}

// GetSellOffersNumber Returns the number of all selloffers
func (k Keeper) GetSellOffersNumber(ctx sdk.Context) (sellOfferId uint64) {
	iterator := k.GetSellOffersIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {
		sellOfferId++
	}
	return
}
