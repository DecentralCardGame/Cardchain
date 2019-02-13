package cardservice

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/bank"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Keeper maintains the link to data storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	coinKeeper bank.Keeper

	cardsStoreKey  sdk.StoreKey // Unexposed key to access card store from sdk.Context
	usersStoreKey  sdk.StoreKey // Unexposed key to access user store from sdk.Context
	internalStoreKey sdk.StoreKey // Unexposed key to access internal variables from sdk.Context

	cdc *codec.Codec // The wire codec for binary encoding/decoding.
}

// NewKeeper creates new instances of the nameservice Keeper
func NewKeeper(coinKeeper bank.Keeper, cardsStoreKey sdk.StoreKey, usersStoreKey sdk.StoreKey, internalStoreKey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{
		coinKeeper:     	coinKeeper,
		cardsStoreKey:  	cardsStoreKey,
		usersStoreKey: 		usersStoreKey,
		internalStoreKey: internalStoreKey,
		cdc:            	cdc,
	}
}

// GetLastCardScheme - get the current id of the last bought card scheme
func (k Keeper) GetLastCardSchemeId(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.internalStoreKey)
	bz := store.Get([]byte("lastCardScheme"))
	return binary.BigEndian.Uint64(bz)
}

// SetLastCardScheme - sets the current id of the last bought card scheme
func (k Keeper) SetLastCardSchemeId(ctx sdk.Context, lastId uint64) {
	store := ctx.KVStore(k.internalStoreKey)
	store.Set([]byte("lastCardScheme"), sdk.Uint64ToBigEndian(lastId))
}

// returns the current price of the card scheme auction
func (k Keeper) GetCardAuctionPrice(ctx sdk.Context) sdk.Coin {
		store := ctx.KVStore(k.internalStoreKey)
		bz := store.Get([]byte("currentCardSchemeAuctionPrice"))
		var price sdk.Coin
		k.cdc.MustUnmarshalBinaryBare(bz, &price)
		return price
}

// sets the current price of the card scheme auction
func (k Keeper) SetCardAuctionPrice(ctx sdk.Context, price sdk.Coin) {
	store := ctx.KVStore(k.internalStoreKey)
	store.Set([]byte("currentCardSchemeAuctionPrice"), k.cdc.MustMarshalBinaryBare(price))
}

// gets the credits in the public pool
func (k Keeper) GetPublicPoolCredits(ctx sdk.Context) sdk.Coin {
		store := ctx.KVStore(k.internalStoreKey)
		bz := store.Get([]byte("publicPoolCredits"))
		var amount sdk.Coin
		k.cdc.MustUnmarshalBinaryBare(bz, &amount)
		return amount
}

// sets the credits in the public pool
func (k Keeper) SetPublicPoolCredits(ctx sdk.Context, price sdk.Coin) {
	store := ctx.KVStore(k.internalStoreKey)
	store.Set([]byte("publicPoolCredits"), k.cdc.MustMarshalBinaryBare(price))
}

// adds or subtracts credits from the public pool
func (k Keeper) DeltaPublicPoolCredits(ctx sdk.Context, delta sdk.Coin) {
	store := ctx.KVStore(k.internalStoreKey)
	bz := store.Get([]byte("publicPoolCredits"))
	var amount sdk.Coin
	k.cdc.MustUnmarshalBinaryBare(bz, &amount)
	newAmount := amount.Plus(delta)
	store.Set([]byte("publicPoolCredits"), k.cdc.MustMarshalBinaryBare(newAmount))
}

func (k Keeper) GetCard(ctx sdk.Context, cardId uint64) Card {
	store := ctx.KVStore(k.cardsStoreKey)
	bz := store.Get(sdk.Uint64ToBigEndian(cardId))

	var gottenCard Card
	k.cdc.MustUnmarshalBinaryBare(bz, &gottenCard)
	return gottenCard
}

func (k Keeper) SetCard(ctx sdk.Context, cardId uint64, newCard Card) {
	store := ctx.KVStore(k.cardsStoreKey)
	store.Set(sdk.Uint64ToBigEndian(cardId), k.cdc.MustMarshalBinaryBare(newCard))
}

func (k Keeper) GetUser(ctx sdk.Context, user sdk.AccAddress) User {
	store := ctx.KVStore(k.usersStoreKey)
	bz := store.Get(user)

	var gottenUser User
	k.cdc.MustUnmarshalBinaryBare(bz, &gottenUser)
	return gottenUser
}

func (k Keeper) SetUser(ctx sdk.Context, address sdk.AccAddress, userData User) {
	store := ctx.KVStore(k.usersStoreKey)
	store.Set(address, k.cdc.MustMarshalBinaryBare(userData))
}

func (k Keeper) GetVoteRights(ctx sdk.Context, voter sdk.AccAddress) []VoteRight {
	store := ctx.KVStore(k.usersStoreKey)
	bz := store.Get(k.cdc.MustMarshalBinaryBare(voter))
	var voteRights []VoteRight
	k.cdc.MustUnmarshalBinaryBare(bz, &voteRights)
	return voteRights
}

func (k Keeper) SetVoteRights(ctx sdk.Context, voteRights []VoteRight, voter sdk.AccAddress) {
	store := ctx.KVStore(k.usersStoreKey)
	store.Set(k.cdc.MustMarshalBinaryBare(voter), k.cdc.MustMarshalBinaryBare(voteRights))
}
