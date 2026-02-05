package generic_type_keeper

import (
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/gogoproto/proto"
)

// ItemIterator is a generic wrapper for sdk.KVStorePrefixIterator that provides unmarshaling of objects
type ItemIterator[T proto.Message, K any] struct {
	iter storetypes.Iterator
	gtk  *GenericTypeKeeper[T, K]
	idx  *uint64
}

// Valid is a wrapper for sdk.KVStorePrefixIterator.Valid
func (i ItemIterator[T, K]) Valid() bool {
	return i.iter.Valid()
}

// Next is a wrapper for sdk.KVStorePrefixIterator.Next
func (i ItemIterator[T, K]) Next() {
	i.iter.Next()
	*i.idx++
}

// Value is a wrapper for sdk.KVStorePrefixIterator.Value but returns object and index
func (i ItemIterator[T, K]) Value() (uint64, T) {
	var gotten T = i.gtk.getEmpty()
	i.gtk.cdc.MustUnmarshal(i.iter.Value(), gotten)

	return *i.idx, gotten
}
