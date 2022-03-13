package keeper

import (
	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetCard Gets card from store
func (k Keeper) GetCard(ctx sdk.Context, cardId uint64) (gottenCard types.Card) {
	store := ctx.KVStore(k.CardsStoreKey)
	bz := store.Get(sdk.Uint64ToBigEndian(cardId))

	k.cdc.MustUnmarshal(bz, &gottenCard)
	return
}

// SetCard Sets card in store
func (k Keeper) SetCard(ctx sdk.Context, cardId uint64, newCard types.Card) {
	store := ctx.KVStore(k.CardsStoreKey)
	store.Set(sdk.Uint64ToBigEndian(cardId), k.cdc.MustMarshal(&newCard))
}

// GetCardAuctionPrice Returns the current price of the card scheme auction
func (k Keeper) GetCardAuctionPrice(ctx sdk.Context) (price sdk.Coin) {
	store := ctx.KVStore(k.InternalStoreKey)
	bz := store.Get([]byte("currentCardSchemeAuctionPrice"))
	k.cdc.MustUnmarshal(bz, &price)
	return
}

// SetCardAuctionPrice Sets the current price of the card scheme auction
func (k Keeper) SetCardAuctionPrice(ctx sdk.Context, price sdk.Coin) {
	store := ctx.KVStore(k.InternalStoreKey)
	store.Set([]byte("currentCardSchemeAuctionPrice"), k.cdc.MustMarshal(&price))
}

// AddOwnedCardScheme Adds a cardscheme to a user
func (k Keeper) AddOwnedCardScheme(ctx sdk.Context, cardId uint64, address sdk.AccAddress) {
	user, _ := k.GetUser(ctx, address)
	user.OwnedCardSchemes = append(user.OwnedCardSchemes, cardId)
	k.SetUser(ctx, address, user)
}

// GetCardsIterator Returns an iterator for all cards
func (k Keeper) GetCardsIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.CardsStoreKey)
	return sdk.KVStorePrefixIterator(store, nil)
}

// GetAllCards Gets all cards form store
func (k Keeper) GetAllCards(ctx sdk.Context) (allCards []*types.Card) {
	iterator := k.GetCardsIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {

		var gottenCard types.Card
		k.cdc.MustUnmarshal(iterator.Value(), &gottenCard)

		allCards = append(allCards, &gottenCard)
	}
	return
}

// GetCardsNumber Gets the number of all card in store
func (k Keeper) GetCardsNumber(ctx sdk.Context) (cardId uint64) {
	iterator := k.GetCardsIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {
		cardId++
	}
	return
}

// SetCardToTrial Sets a card to trial
func (k Keeper) SetCardToTrial(ctx sdk.Context, cardId uint64, votePool sdk.Coin) {
	card := k.GetCard(ctx, cardId)
	card.ResetVotes()
	card.VotePool = card.VotePool.Add(votePool)
	card.Status = types.Status_trial
	k.SetCard(ctx, cardId, card)
}
