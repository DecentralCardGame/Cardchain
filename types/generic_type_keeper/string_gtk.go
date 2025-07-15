package generic_type_keeper

import (
	"cosmossdk.io/core/store"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/gogoproto/proto"
)

type GenericStringTypeKeeper[T proto.Message] struct {
	GenericTypeKeeper[T, string]
}

func NewStringGTK[T proto.Message](key string, storeService store.KVStoreService, cdc codec.BinaryCodec, getEmpty func() T) GenericStringTypeKeeper[T] {
	gtk := GenericStringTypeKeeper[T]{
		GenericTypeKeeper[T, string]{
			BaseKeeper: BaseKeeper[T]{
				key:          key,
				cdc:          cdc,
				storeService: storeService,
				getEmpty:     getEmpty,
			},
			keyConverter: func(bz []byte, key string) []byte {
				return append(bz, key...)
			},
		},
	}
	return gtk
}
