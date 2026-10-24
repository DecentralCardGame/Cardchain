package generic_type_keeper

import (
	"cosmossdk.io/core/store"
	"cosmossdk.io/store/prefix"
	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/gogoproto/proto"
)

// GetEmpty Just returns a empty object of a certain type
func GetEmpty[A any]() *A {
	var obj A
	return &obj
}

type BaseKeeper[T proto.Message] struct {
	storeService store.KVStoreService
	key          string
	cdc          codec.BinaryCodec
	getEmpty     func() T
}

func (bk BaseKeeper[T]) valueKey() string {
	return bk.key + "/value/"
}

func (bk BaseKeeper[T]) getValueStore(ctx sdk.Context) prefix.Store {
	storeAdapter := runtime.KVStoreAdapter(bk.storeService.OpenKVStore(ctx))
	return prefix.NewStore(storeAdapter, types.KeyPrefix(bk.valueKey()))
}
