package keeper

import (
	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

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

// GetLastCardModified Returns the block height the last card was modified at
func (k Keeper) GetLastCardModified(ctx sdk.Context) (timeStamp types.TimeStamp) {
	store := ctx.KVStore(k.InternalStoreKey)
	bz := store.Get([]byte("lastCardModified"))
	k.cdc.MustUnmarshal(bz, &timeStamp)
	return
}

// SetLastCardModified Sets the block height the last card was modified at
func (k Keeper) SetLastCardModified(ctx sdk.Context, timeStamp types.TimeStamp) {
	store := ctx.KVStore(k.InternalStoreKey)
	store.Set([]byte("lastCardModified"), k.cdc.MustMarshal(&timeStamp))
}

// SetLastCardModifiedNow Sets the block height the last card was modified at, to now
func (k Keeper) SetLastCardModifiedNow(ctx sdk.Context) {
	timeStamp := types.TimeStamp{TimeStamp: uint64(ctx.BlockHeight())}
	k.SetLastCardModified(ctx, timeStamp)
}

// AddOwnedCardScheme Adds a cardscheme to a user
func (k Keeper) AddOwnedCardScheme(ctx sdk.Context, cardId uint64, address sdk.AccAddress) {
	user, _ := k.GetUser(ctx, address)
	user.OwnedCardSchemes = append(user.OwnedCardSchemes, cardId)
	k.SetUser(ctx, address, user)
}

// SetCardToTrial Sets a card to trial
func (k Keeper) SetCardToTrial(ctx sdk.Context, cardId uint64, votePool sdk.Coin) {
	card := k.Cards.Get(ctx, cardId)
	card.ResetVotes()
	card.VotePool = card.VotePool.Add(votePool)
	card.Status = types.Status_trial
	k.Cards.Set(ctx, cardId, card)
	k.SetLastCardModifiedNow(ctx)
}
