package generic_type_keeper

import (
	"fmt"
	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type ItemIterator[T codec.ProtoMarshaler] struct {
	iter sdk.Iterator
	gtk  *GenericTypeKeeper[T]
	idx  *uint64
}

func (i ItemIterator[T]) Valid() bool {
	return i.iter.Valid()
}

func (i ItemIterator[T]) Next() {
	i.iter.Next()
	*i.idx++
}

func (i ItemIterator[T]) Value() (uint64, T) {
	var gotten T = i.gtk.getEmpty()
	i.gtk.cdc.MustUnmarshal(i.iter.Value(), gotten)

	return *i.idx, gotten
}

// GetEmpty Just returns a empty object of a certain type
func GetEmpty[A any]() *A {
	var obj A
	return &obj
}

type GenericTypeKeeper[T codec.ProtoMarshaler] struct {
	Key      sdk.StoreKey
  InternalKey      sdk.StoreKey
	cdc      codec.BinaryCodec
	name     string
	getEmpty func() T // This is needed because codec.ProtoMarshaler always refers to a pointer, but for cdc.Unmarshal to work the passed pointer can't be nil, but when initializing a pointer it's nil
}

// NewGTK Returns a new GenericTypeKeeper
func NewGTK[T codec.ProtoMarshaler](key sdk.StoreKey, internalKey sdk.StoreKey, cdc codec.BinaryCodec, getEmpty func() T) GenericTypeKeeper[T] {
	gtk := GenericTypeKeeper[T]{
		Key:      key,
    InternalKey: internalKey,
		cdc:      cdc,
		name:     fmt.Sprintf("%T", getEmpty()),
		getEmpty: getEmpty,
	}
	return gtk
}

// Get Gets an object from store
func (gtk GenericTypeKeeper[T]) Get(ctx sdk.Context, id uint64) T {
	store := ctx.KVStore(gtk.Key)
	bz := store.Get(sdk.Uint64ToBigEndian(id))

	gotten := gtk.getEmpty()
	gtk.cdc.MustUnmarshal(bz, gotten)
	return gotten
}

// Set Sets an object in store
func (gtk GenericTypeKeeper[T]) Set(ctx sdk.Context, id uint64, new T) {
	store := ctx.KVStore(gtk.Key)
	store.Set(sdk.Uint64ToBigEndian(id), gtk.cdc.MustMarshal(new))
	num := gtk.GetNum(ctx)
	if id == num {
		gtk.SetNum(ctx, num+1)
	}
}

// GetCardAuctionPrice Returns the current price of the card scheme auction
func (gtk GenericTypeKeeper[T]) GetNum(ctx sdk.Context) uint64 {
	store := ctx.KVStore(gtk.InternalKey)
	bz := store.Get([]byte(gtk.name))
	var num types.Num
	gtk.cdc.MustUnmarshal(bz, &num)
	return num.Num
}

// SetCardAuctionPrice Sets the current price of the card scheme auction
func (gtk GenericTypeKeeper[T]) SetNum(ctx sdk.Context, _num uint64) {
	store := ctx.KVStore(gtk.InternalKey)
	var num = types.Num{Num: _num}
	store.Set([]byte(gtk.name), gtk.cdc.MustMarshal(&num))
}

// GetIterator Returns an iterator for all objects
func (gtk GenericTypeKeeper[T]) GetIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(gtk.Key)
	return sdk.KVStorePrefixIterator(store, nil)
}

// GetAll Gets all objs from store
func (gtk GenericTypeKeeper[T]) GetAll(ctx sdk.Context) (all []T) {
	iterator := gtk.GetIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {

		var gotten T = gtk.getEmpty()
		gtk.cdc.MustUnmarshal(iterator.Value(), gotten)

		all = append(all, gotten)
	}
	return
}

// GetAll Gets all objs from store
func (gtk GenericTypeKeeper[T]) GetItemIterator(ctx sdk.Context) ItemIterator[T] {
	iterator := gtk.GetIterator(ctx)
  var num uint64 = 0
	return ItemIterator[T]{
		iter: iterator,
		gtk:  &gtk,
    idx: &num,
	}
}

// GetNumber Gets the number of all objs in store
func (gtk GenericTypeKeeper[T]) GetNumber(ctx sdk.Context) (id uint64) {
	iterator := gtk.GetIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {
		id++
	}
	return
}
