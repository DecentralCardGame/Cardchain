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

	namesStoreKey  sdk.StoreKey // Unexposed key to access name store from sdk.Context
	ownersStoreKey sdk.StoreKey // Unexposed key to access owners store from sdk.Context
	pricesStoreKey sdk.StoreKey // Unexposed key to access prices store from sdk.Context
	typesStoreKey  sdk.StoreKey // Unexposed key to access types store from sdk.Context

	internalStoreKey sdk.StoreKey // Unexposed key to access internal variables from sdk.Context

	cdc *codec.Codec // The wire codec for binary encoding/decoding.
}

// NewKeeper creates new instances of the nameservice Keeper
func NewKeeper(coinKeeper bank.Keeper, namesStoreKey sdk.StoreKey, ownersStoreKey sdk.StoreKey, priceStoreKey sdk.StoreKey, typesStoreKey sdk.StoreKey, internalStoreKey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{
		coinKeeper:     coinKeeper,
		namesStoreKey:  namesStoreKey,
		ownersStoreKey: ownersStoreKey,
		pricesStoreKey: priceStoreKey,
		typesStoreKey:  typesStoreKey,
		internalStoreKey: internalStoreKey,
		cdc:            cdc,
	}
}

// ResolveName - returns the string that the name resolves to
func (k Keeper) ResolveName(ctx sdk.Context, name string) string {
	store := ctx.KVStore(k.namesStoreKey)
	bz := store.Get([]byte(name))
	return string(bz)
}

// SetName - sets the value string that a name resolves to
func (k Keeper) SetName(ctx sdk.Context, name string, value string) {
	store := ctx.KVStore(k.namesStoreKey)
	store.Set([]byte(name), []byte(value))
}

// HasOwner - returns whether or not the name already has an owner
func (k Keeper) HasOwner(ctx sdk.Context, name string) bool {
	store := ctx.KVStore(k.ownersStoreKey)
	bz := store.Get([]byte(name))
	return bz != nil
}

// GetOwner - get the current owner of a name
func (k Keeper) GetOwner(ctx sdk.Context, name string) sdk.AccAddress {
	store := ctx.KVStore(k.ownersStoreKey)
	bz := store.Get([]byte(name))
	return bz
}

// SetOwner - sets the current owner of a name
func (k Keeper) SetOwner(ctx sdk.Context, name string, owner sdk.AccAddress) {
	store := ctx.KVStore(k.ownersStoreKey)
	store.Set([]byte(name), owner)
}

// GetPrice - gets the current price of a name.  If price doesn't exist yet, set to 1steak.
func (k Keeper) GetPrice(ctx sdk.Context, name string) sdk.Coins {
	if !k.HasOwner(ctx, name) {
		return sdk.Coins{sdk.NewInt64Coin("mycoin", 1)}
	}
	store := ctx.KVStore(k.pricesStoreKey)
	bz := store.Get([]byte(name))
	var price sdk.Coins
	// k.cdc.MustUnmarshalBinary(bz, &price)
	k.cdc.MustUnmarshalBinaryBare(bz, &price)
	return price
}

// SetPrice - sets the current price of a name
func (k Keeper) SetPrice(ctx sdk.Context, name string, price sdk.Coins) {
	store := ctx.KVStore(k.pricesStoreKey)
	store.Set([]byte(name), k.cdc.MustMarshalBinaryBare(price))
}


// GetOwner - get the current owner of a card
func (k Keeper) GetType(ctx sdk.Context, name string) string  {
	store := ctx.KVStore(k.typesStoreKey)
	bz := store.Get([]byte(name))
	return string(bz)
}

// SetPrice - sets the current type of a card
func (k Keeper) SetType(ctx sdk.Context, name string, cardtype string) {
	store := ctx.KVStore(k.typesStoreKey)
	store.Set([]byte(name), []byte(cardtype))
}

// GetLastCardScheme - get the current id of the last bought card scheme
func (k Keeper) GetLastCardScheme(ctx sdk.Context) string  {
	store := ctx.KVStore(k.internalStoreKey)
	bz := store.Get([]byte("lastCardScheme"))
	return string(bz)
}

// SetLastCardScheme - sets the current id of the last bought card scheme
func (k Keeper) SetLastCardScheme(ctx sdk.Context, lastId string) {
	store := ctx.KVStore(k.internalStoreKey)
	store.Set([]byte("lastCardScheme"), []byte(lastId))
}

func (k Keeper) GetCardAuctionPrice(ctx sdk.Context) uint64 {
		store := ctx.KVStore(k.internalStoreKey)
		bz := store.Get([]byte("currentCardSchemeAuctionPrice"))
		price := binary.BigEndian.Uint64(bz)
		return price
}

func (k Keeper) SetCardAuctionPrice(ctx sdk.Context, price uint64) {
	store := ctx.KVStore(k.internalStoreKey)
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, price)
	store.Set([]byte("currentCardSchemeAuctionPrice"), []byte(b))
}

//
func (k Keeper) DoubleCardSchemeAuctionPrice(ctx sdk.Context) {
	store := ctx.KVStore(k.internalStoreKey)
	bz := store.Get([]byte("currentCardSchemeAuctionPrice"))
	store.Set([]byte("lastCardScheme"), []byte(bz))
}
