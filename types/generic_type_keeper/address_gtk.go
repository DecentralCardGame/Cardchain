package generic_type_keeper

import (
	"bytes"

	"cosmossdk.io/core/store"
	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/gogoproto/proto"
)

type GenericAddressTypeKeeper[T proto.Message] struct {
	GenericTypeKeeper[T, sdk.AccAddress]
}

func (gtk GenericAddressTypeKeeper[T]) NormalizeAddress(address []byte) sdk.AccAddress {
	return bytes.TrimPrefix(address, append(types.KeyPrefix(gtk.valueKey()), []byte("/")...))
}

func NewAddressGTK[T proto.Message](key string, storeService store.KVStoreService, cdc codec.BinaryCodec, getEmpty func() T) GenericAddressTypeKeeper[T] {
	gtk := GenericAddressTypeKeeper[T]{
		GenericTypeKeeper[T, sdk.AccAddress]{
			BaseKeeper: BaseKeeper[T]{
				key:          key,
				cdc:          cdc,
				storeService: storeService,
				getEmpty:     getEmpty,
			},
			keyConverter: func(bz []byte, address sdk.AccAddress) []byte {
				return append(bz, address...)
			},
		},
	}
	return gtk
}
