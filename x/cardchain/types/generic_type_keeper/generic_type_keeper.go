package generic_type_keeper

import(
  sdk "github.com/cosmos/cosmos-sdk/types"
  "github.com/cosmos/cosmos-sdk/codec"
)

// GetEmpty Just returns a empty object of a certain type
func GetEmpty [A any]() *A {
  var obj A
  return &obj
}

type GenericTypeKeeper[T codec.ProtoMarshaler] struct {
  Key sdk.StoreKey
  cdc codec.BinaryCodec
  getEmpty func() T  // This is needed because codec.ProtoMarshaler always refers to a pointer, but for cdc.Unmarshal to work the passed pointer can't be nil, but when initializing a pointer it's nil
}

// NewGTK Returns a new GenericTypeKeeper
func NewGTK[T codec.ProtoMarshaler](key sdk.StoreKey, cdc codec.BinaryCodec, getEmpty func() T) GenericTypeKeeper[T] {
  gtk := GenericTypeKeeper[T] {
    Key: key,
    cdc: cdc,
    getEmpty: getEmpty,
  }
  return gtk
}

// Get Gets an object from store
func (gtk GenericTypeKeeper[T]) Get(ctx sdk.Context, id uint64) (T) {
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

// GetNumber Gets the number of all objs in store
func (gtk GenericTypeKeeper[T]) GetNumber(ctx sdk.Context) (id uint64) {
	iterator := gtk.GetIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {
		id++
	}
	return
}
