package generic_type_keeper

import (
	"encoding/binary"

	"cosmossdk.io/core/store"
	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/gogoproto/proto"
)

type GenericUint64TypeKeeper[T proto.Message] struct {
	GenericTypeKeeper[T, uint64]
}

func NewUintGTK[T proto.Message](key string, storeService store.KVStoreService, cdc codec.BinaryCodec, getEmpty func() T) GenericUint64TypeKeeper[T] {
	gtk := GenericUint64TypeKeeper[T]{
		GenericTypeKeeper[T, uint64]{
			key:          key,
			cdc:          cdc,
			getEmpty:     getEmpty,
			storeService: storeService,
			keyConverter: func(bz []byte, id uint64) []byte {
				return binary.BigEndian.AppendUint64(bz, id)
			},
		},
	}
	return gtk
}

// Append Sets a new object
func (gtk GenericUint64TypeKeeper[T]) Append(ctx sdk.Context, new T) (count uint64) {
	store := gtk.getValueStore(ctx)
	count = gtk.GetNum(ctx)

	store.Set(gtk.getIDBytes(count), gtk.cdc.MustMarshal(new))

	gtk.setNum(ctx, count+1)

	return
}

// GetNum Returns the number of items stored, way more performant than GetNumber
func (gtk GenericUint64TypeKeeper[T]) GetNum(ctx sdk.Context) uint64 {
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
func (gtk GenericUint64TypeKeeper[T]) setNum(ctx sdk.Context, count uint64) {
	store := gtk.getCountStore(ctx)
	byteKey := types.KeyPrefix(gtk.countKey())
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}
