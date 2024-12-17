package generic_type_keeper

import (
	"encoding/binary"

	"cosmossdk.io/core/store"
	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	db "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/gogoproto/proto"
)

// ItemIterator is a generic wrapper for sdk.KVStorePrefixIterator that provides unmarshaling of objects
type ItemIterator[T proto.Message] struct {
	iter storetypes.Iterator
	gtk  *GenericTypeKeeper[T]
	idx  *uint64
}

// Valid is a wrapper for sdk.KVStorePrefixIterator.Valid
func (i ItemIterator[T]) Valid() bool {
	return i.iter.Valid()
}

// Next is a wrapper for sdk.KVStorePrefixIterator.Next
func (i ItemIterator[T]) Next() {
	i.iter.Next()
	*i.idx++
}

// Value is a wrapper for sdk.KVStorePrefixIterator.Value but returns object and index
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

type GenericTypeKeeper[T proto.Message] struct {
	storeService store.KVStoreService
	key          string
	cdc          codec.BinaryCodec
	getEmpty     func() T // This is needed because codec.ProtoMarshaler always refers to a pointer, but for cdc.Unmarshal to work the passed pointer can't be nil, but when initializing a pointer it's nil
}

// NewGTK Returns a new GenericTypeKeeper
func NewGTK[T proto.Message](key string, storeService store.KVStoreService, cdc codec.BinaryCodec, getEmpty func() T) GenericTypeKeeper[T] {
	gtk := GenericTypeKeeper[T]{
		key:          key,
		cdc:          cdc,
		getEmpty:     getEmpty,
		storeService: storeService,
	}
	return gtk
}

func (gtk GenericTypeKeeper[T]) valueKey() string {
	return gtk.key + "/value/"
}

func (gtk GenericTypeKeeper[T]) countKey() string {
	return gtk.key + "/count/"
}

func (gtk GenericTypeKeeper[T]) getIDBytes(id uint64) []byte {
	bz := types.KeyPrefix(types.ProductDetailsKey)
	bz = append(bz, []byte("/")...)
	bz = binary.BigEndian.AppendUint64(bz, id)
	return bz
}

func (gtk GenericTypeKeeper[T]) getValueStore(ctx sdk.Context) prefix.Store {
	storeAdapter := runtime.KVStoreAdapter(gtk.storeService.OpenKVStore(ctx))
	return prefix.NewStore(storeAdapter, types.KeyPrefix(gtk.valueKey()))
}

func (gtk GenericTypeKeeper[T]) getCountStore(ctx sdk.Context) prefix.Store {
	storeAdapter := runtime.KVStoreAdapter(gtk.storeService.OpenKVStore(ctx))
	return prefix.NewStore(storeAdapter, []byte{})
}

// Get Gets an object from store
func (gtk GenericTypeKeeper[T]) Get(ctx sdk.Context, id uint64) T {
	store := gtk.getValueStore(ctx)
	bz := store.Get(gtk.getIDBytes(id))

	gotten := gtk.getEmpty()
	gtk.cdc.MustUnmarshal(bz, gotten)
	return gotten
}

// Set Sets an object in store
func (gtk GenericTypeKeeper[T]) Set(ctx sdk.Context, id uint64, new T) {
	store := gtk.getValueStore(ctx)
	store.Set(gtk.getIDBytes(id), gtk.cdc.MustMarshal(new))
	num := gtk.GetNum(ctx)
	if id == num {
		gtk.setNum(ctx, num+1)
	}
}

// Append Sets a new object
func (gtk GenericTypeKeeper[T]) Append(ctx sdk.Context, new T) (count uint64) {
	store := gtk.getValueStore(ctx)
	count = gtk.GetNum(ctx)

	store.Set(gtk.getIDBytes(count), gtk.cdc.MustMarshal(new))

	gtk.setNum(ctx, count+1)

	return
}

// GetNum Returns the number of items stored, way more performant than GetNumber
func (gtk GenericTypeKeeper[T]) GetNum(ctx sdk.Context) uint64 {
	store := gtk.getCountStore(ctx)
	byteKey := types.KeyPrefix(gtk.countKey())
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// setNum Sets thenumber of items stored
func (gtk GenericTypeKeeper[T]) setNum(ctx sdk.Context, count uint64) {
	store := gtk.getCountStore(ctx)
	byteKey := types.KeyPrefix(gtk.countKey())
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// GetIterator Returns an iterator for all objects
func (gtk GenericTypeKeeper[T]) GetIterator(ctx sdk.Context) db.Iterator {
	store := gtk.getValueStore(ctx)
	return storetypes.KVStorePrefixIterator(store, []byte{})
}

// GetAll Gets all objs from store -- use GetItemIterator instead
func (gtk GenericTypeKeeper[T]) GetAll(ctx sdk.Context) (all []T) {
	iterator := gtk.GetIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {

		var gotten T = gtk.getEmpty()
		gtk.cdc.MustUnmarshal(iterator.Value(), gotten)

		all = append(all, gotten)
	}
	return
}

// GetItemIterator Gets all objs from store
func (gtk GenericTypeKeeper[T]) GetItemIterator(ctx sdk.Context) ItemIterator[T] {
	iterator := gtk.GetIterator(ctx)
	var num uint64 = 0
	return ItemIterator[T]{
		iter: iterator,
		gtk:  &gtk,
		idx:  &num,
	}
}

// GetNumber Gets the number of all objs in store
// Deprecated: don't use
func (gtk GenericTypeKeeper[T]) GetNumber(ctx sdk.Context) (id uint64) {
	iterator := gtk.GetIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {
		id++
	}
	return
}
