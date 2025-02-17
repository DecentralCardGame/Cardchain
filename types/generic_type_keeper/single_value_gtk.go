package generic_type_keeper

import (
	"cosmossdk.io/core/store"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/gogoproto/proto"
)

func NewSingleValueGenericTypeKeeper[T proto.Message](key string, storeService store.KVStoreService, cdc codec.BinaryCodec, getEmpty func() T) SingleValueGenericTypeKeeper[T] {
	return SingleValueGenericTypeKeeper[T]{
		BaseKeeper: BaseKeeper[T]{
			key:          key,
			cdc:          cdc,
			storeService: storeService,
			getEmpty:     getEmpty,
		},
	}
}

type SingleValueGenericTypeKeeper[T proto.Message] struct {
	BaseKeeper[T]
}

// Get Gets an object from store
func (svgtk SingleValueGenericTypeKeeper[T]) Get(ctx sdk.Context) T {
	store := svgtk.getValueStore(ctx)
	bz := store.Get([]byte{1})

	gotten := svgtk.getEmpty()
	svgtk.cdc.MustUnmarshal(bz, gotten)
	return gotten
}

// Set Sets an object in store
func (svgtk SingleValueGenericTypeKeeper[T]) Set(ctx sdk.Context, new T) {
	store := svgtk.getValueStore(ctx)
	store.Set([]byte{1}, svgtk.cdc.MustMarshal(new))
}
