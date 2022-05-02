package generic_type_keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"golang.org/x/exp/slices"
)

type KeywordedGenericTypeKeeper[T codec.ProtoMarshaler] struct {
	GenericTypeKeeper[T]
	KeyWords []string
}

// NewKGTK Returns a new KeywordedGenericTypeKeeper
func NewKGTK[T codec.ProtoMarshaler](key sdk.StoreKey, internalKey sdk.StoreKey, cdc codec.BinaryCodec, getEmpty func() T, keywords []string) KeywordedGenericTypeKeeper[T] {
	gtk := KeywordedGenericTypeKeeper[T]{
		NewGTK[T](key, internalKey, cdc, getEmpty),
		keywords,
	}
	return gtk
}

// Get Gets an object from store
func (gtk KeywordedGenericTypeKeeper[T]) Get(ctx sdk.Context, keyword string) T {
	if !slices.Contains(gtk.KeyWords, keyword) {
		panic("Unknown keyword: " + keyword)
	}

	store := ctx.KVStore(gtk.Key)
	bz := store.Get([]byte(keyword))

	gotten := gtk.getEmpty()
	gtk.cdc.MustUnmarshal(bz, gotten)
	return gotten
}

// Set Sets an object in store
func (gtk KeywordedGenericTypeKeeper[T]) Set(ctx sdk.Context, keyword string, new T) {
	if !slices.Contains(gtk.KeyWords, keyword) {
		panic("Unknown keyword: " + keyword)
	}

	store := ctx.KVStore(gtk.Key)
	store.Set([]byte(keyword), gtk.cdc.MustMarshal(new))
}

// GetAll Gets all objs from store
func (gtk KeywordedGenericTypeKeeper[T]) GetAll(ctx sdk.Context) (all []T) {
	for _, keyWord := range gtk.KeyWords {
		all = append(all, gtk.Get(ctx, keyWord))
	}
	return
}

// GetNumber Gets the number of all objs in store
func (gtk KeywordedGenericTypeKeeper[T]) GetNumber(ctx sdk.Context) (id uint64) {
	return uint64(len(gtk.KeyWords))
}
