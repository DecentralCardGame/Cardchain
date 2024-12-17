package generic_type_keeper

import (
	"slices"

	"cosmossdk.io/core/store"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/gogoproto/proto"
)

type KeywordedGenericTypeKeeper[T proto.Message] struct {
	GenericTypeKeeper[T]
	KeyWords []string
}

// NewKGTK Returns a new KeywordedGenericTypeKeeper
func NewKGTK[T proto.Message](key string, storeService store.KVStoreService, cdc codec.BinaryCodec, getEmpty func() T, keywords []string) KeywordedGenericTypeKeeper[T] {
	return KeywordedGenericTypeKeeper[T]{
		GenericTypeKeeper: NewGTK[T](key, storeService, cdc, getEmpty),
		KeyWords:          keywords,
	}
}

func (gtk KeywordedGenericTypeKeeper[T]) getUintID(keyword string) uint64 {
	id := slices.Index(gtk.KeyWords, keyword)
	if id == -1 {
		panic("Unknown keyword: " + keyword)
	}

	return uint64(id)
}

// Get Gets an object from store
func (gtk KeywordedGenericTypeKeeper[T]) Get(ctx sdk.Context, keyword string) T {
	return gtk.GenericTypeKeeper.Get(ctx, gtk.getUintID(keyword))
}

// Set Sets an object in store
func (gtk KeywordedGenericTypeKeeper[T]) Set(ctx sdk.Context, keyword string, new T) {
	gtk.GenericTypeKeeper.Set(ctx, gtk.getUintID(keyword), new)
}

// GetAll Gets all objs from store
func (gtk KeywordedGenericTypeKeeper[T]) GetAll(ctx sdk.Context) (all []T) {
	for _, keyWord := range gtk.KeyWords {
		all = append(all, gtk.Get(ctx, keyWord))
	}
	return
}
