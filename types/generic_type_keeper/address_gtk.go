package generic_type_keeper

import (
	"cosmossdk.io/core/store"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/gogoproto/proto"
)

type GenericAddressTypeKeeper[T proto.Message] struct {
	GenericTypeKeeper[T, sdk.AccAddress]
}

func NewAddressGTK[T proto.Message](key string, storeService store.KVStoreService, cdc codec.BinaryCodec, getEmpty func() T) GenericAddressTypeKeeper[T] {
	gtk := GenericAddressTypeKeeper[T]{
		GenericTypeKeeper[T, sdk.AccAddress]{
			key:          key,
			cdc:          cdc,
			getEmpty:     getEmpty,
			storeService: storeService,
			keyConverter: func(bz []byte, address sdk.AccAddress) []byte {
				return append(bz, address...)
			},
		},
	}
	return gtk
}
