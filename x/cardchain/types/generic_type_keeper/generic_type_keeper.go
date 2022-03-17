package generic_type_keeper

import(
  sdk "github.com/cosmos/cosmos-sdk/types"
  "github.com/cosmos/cosmos-sdk/codec"
)

type GenericTypeKeeper[T codec.ProtoMarshaler] struct {
  Key sdk.StoreKey
  cdc codec.BinaryCodec
}

func NewGTK[T codec.ProtoMarshaler](key sdk.StoreKey, cdc codec.BinaryCodec) GenericTypeKeeper[T] {
  gtk := GenericTypeKeeper[T] {
    Key: key,
    cdc: cdc,
  }
  return gtk
}

// Get Gets an object from store
func (gtk GenericTypeKeeper[T]) Get(ctx sdk.Context, id uint64) (gotten T) {
  store := ctx.KVStore(gtk.Key)
	bz := store.Get(sdk.Uint64ToBigEndian(id))

	gtk.cdc.MustUnmarshal(bz, gotten)
	return
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
func (gtk GenericTypeKeeper[T]) GetAll(ctx sdk.Context) (all []*T) {
	iterator := gtk.GetIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {

		var gotten T
		gtk.cdc.MustUnmarshal(iterator.Value(), gotten)

		all = append(all, &gotten)
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
