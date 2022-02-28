package keeper

import (
  "github.com/DecentralCardGame/Cardchain/x/cardchain/types"
  sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) GetCard(ctx sdk.Context, cardId uint64) types.Card {
	store := ctx.KVStore(k.CardsStoreKey)
	bz := store.Get(sdk.Uint64ToBigEndian(cardId))

	var gottenCard types.Card
	k.cdc.MustUnmarshal(bz, &gottenCard)
	return gottenCard
}

func (k Keeper) SetCard(ctx sdk.Context, cardId uint64, newCard types.Card) {
	store := ctx.KVStore(k.CardsStoreKey)
	store.Set(sdk.Uint64ToBigEndian(cardId), k.cdc.MustMarshal(&newCard))
}

// returns the current price of the card scheme auction
func (k Keeper) GetCardAuctionPrice(ctx sdk.Context) sdk.Coin {
	store := ctx.KVStore(k.InternalStoreKey)
	bz := store.Get([]byte("currentCardSchemeAuctionPrice"))
	var price sdk.Coin
	k.cdc.MustUnmarshal(bz, &price)
	return price
}

// sets the current price of the card scheme auction
func (k Keeper) SetCardAuctionPrice(ctx sdk.Context, price sdk.Coin) {
	store := ctx.KVStore(k.InternalStoreKey)
	store.Set([]byte("currentCardSchemeAuctionPrice"), k.cdc.MustMarshal(&price))
}

func (k Keeper) AddOwnedCardScheme(ctx sdk.Context, cardId uint64, address sdk.AccAddress) {
	store := ctx.KVStore(k.UsersStoreKey)
	bz := store.Get(address)

	var gottenUser types.User
	k.cdc.MustUnmarshal(bz, &gottenUser)

	gottenUser.OwnedCardSchemes = append(gottenUser.OwnedCardSchemes, cardId)

	store.Set(address, k.cdc.MustMarshal(&gottenUser))
}

func (k Keeper) GetCardsIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.CardsStoreKey)
	return sdk.KVStorePrefixIterator(store, nil)
}

func (k Keeper) GetAllCards(ctx sdk.Context) []*types.Card {
	var allCards []*types.Card
	iterator := k.GetCardsIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {

		var gottenCard types.Card
		k.cdc.MustUnmarshal(iterator.Value(), &gottenCard)

		allCards = append(allCards, &gottenCard)
	}
	return allCards
}

func (k Keeper) GetCardsNumber(ctx sdk.Context) uint64 {
	var cardId uint64
	iterator := k.GetCardsIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {
		cardId++
	}
	return cardId
}
