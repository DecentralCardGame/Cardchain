package keeper

import (
	"fmt"

	"github.com/DecentralCardGame/Cardchain/x/featureflag/types"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/tendermint/tendermint/libs/log"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   storetypes.StoreKey
		memKey     storetypes.StoreKey
		paramstore paramtypes.Subspace
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,

) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramstore: ps,
	}
}

func (k Keeper) GetModuleInstance(ctx sdk.Context, moduleName string, flagNames []string) ModuleInstance {
	return NewModuleInstance(ctx, k, moduleName, flagNames)
}

func getKey(module string, name string) []byte {
	return []byte(module+"-"+name)
}

func (k Keeper) GetFlag(ctx sdk.Context, key []byte) (gottenFlag types.Flag) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(key)
	k.cdc.MustUnmarshal(bz, &gottenFlag)
	return
}

func (k Keeper) SetFlag(ctx sdk.Context, key []byte, flag types.Flag) {
	store := ctx.KVStore(k.storeKey)
	store.Set(key, k.cdc.MustMarshal(&flag))
}

func (k Keeper) GetFlagsIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, nil)
}

func (k Keeper) GetAllFlags(ctx sdk.Context) (allFlags []*types.Flag, allKeys []string) {
	iterator := k.GetFlagsIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {

		var gottenFlag types.Flag
		k.cdc.MustUnmarshal(iterator.Value(), &gottenFlag)

		allFlags = append(allFlags, &gottenFlag)
		allKeys = append(allKeys, string(iterator.Key()))
	}
	return
}


func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
