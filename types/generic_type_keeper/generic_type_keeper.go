package generic_type_keeper

import (
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

// GetEmpty Just returns a empty object of a certain type
func GetEmpty[A any]() *A {
	var obj A
	return &obj
}

type GenericTypeKeeper[T proto.Message, K any] struct {
	storeService store.KVStoreService
	key          string
	cdc          codec.BinaryCodec
	getEmpty     func() T // This is needed because codec.ProtoMarshaler always refers to a pointer, but for cdc.Unmarshal to work the passed pointer can't be nil, but when initializing a pointer it's nil
	keyConverter KeyConver[K]
}

// NewGTK Returns a new GenericTypeKeeper
func NewGTK[T proto.Message, K any](key string, storeService store.KVStoreService, cdc codec.BinaryCodec, getEmpty func() T, keyConverter KeyConver[K]) GenericTypeKeeper[T, K] {
	gtk := GenericTypeKeeper[T, K]{
		key:          key,
		cdc:          cdc,
		getEmpty:     getEmpty,
		storeService: storeService,
		keyConverter: keyConverter,
	}
	return gtk
}

func (gtk GenericTypeKeeper[T, K]) valueKey() string {
	return gtk.key + "/value/"
}

func (gtk GenericTypeKeeper[T, K]) countKey() string {
	return gtk.key + "/count/"
}

func (gtk GenericTypeKeeper[T, K]) getIDBytes(id K) []byte {
	bz := types.KeyPrefix(gtk.valueKey())
	bz = append(bz, []byte("/")...)
	bz = gtk.keyConverter(bz, id)
	return bz
}

func (gtk GenericTypeKeeper[T, K]) getValueStore(ctx sdk.Context) prefix.Store {
	storeAdapter := runtime.KVStoreAdapter(gtk.storeService.OpenKVStore(ctx))
	return prefix.NewStore(storeAdapter, types.KeyPrefix(gtk.valueKey()))
}

func (gtk GenericTypeKeeper[T, K]) getCountStore(ctx sdk.Context) prefix.Store {
	storeAdapter := runtime.KVStoreAdapter(gtk.storeService.OpenKVStore(ctx))
	return prefix.NewStore(storeAdapter, []byte{})
}

// Get Gets an object from store
func (gtk GenericTypeKeeper[T, K]) Get(ctx sdk.Context, id K) T {
	store := gtk.getValueStore(ctx)
	bz := store.Get(gtk.getIDBytes(id))

	gotten := gtk.getEmpty()
	gtk.cdc.MustUnmarshal(bz, gotten)
	return gotten
}

// Set Sets an object in store
func (gtk GenericTypeKeeper[T, K]) Set(ctx sdk.Context, id K, new T) {
	store := gtk.getValueStore(ctx)
	store.Set(gtk.getIDBytes(id), gtk.cdc.MustMarshal(new))
}

// GetIterator Returns an iterator for all objects
func (gtk GenericTypeKeeper[T, K]) GetIterator(ctx sdk.Context) db.Iterator {
	store := gtk.getValueStore(ctx)
	return storetypes.KVStorePrefixIterator(store, []byte{})
}

// GetAll Gets all objs from store -- use GetItemIterator instead
func (gtk GenericTypeKeeper[T, K]) GetAll(ctx sdk.Context) (all []T) {
	iterator := gtk.GetIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {

		var gotten T = gtk.getEmpty()
		gtk.cdc.MustUnmarshal(iterator.Value(), gotten)

		all = append(all, gotten)
	}
	return
}

// GetItemIterator Gets all objs from store
func (gtk GenericTypeKeeper[T, K]) GetItemIterator(ctx sdk.Context) ItemIterator[T, K] {
	iterator := gtk.GetIterator(ctx)
	var num uint64 = 0
	return ItemIterator[T, K]{
		iter: iterator,
		gtk:  &gtk,
		idx:  &num,
	}
}

// GetNumber Gets the number of all objs in store
// Deprecated: don't use
func (gtk GenericTypeKeeper[T, K]) GetNumber(ctx sdk.Context) (id uint64) {
	iterator := gtk.GetIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {
		id++
	}
	return
}
