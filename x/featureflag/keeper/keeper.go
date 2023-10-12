package keeper

import (
	"fmt"

	sdkerrors "cosmossdk.io/errors"
	"github.com/DecentralCardGame/Cardchain/x/featureflag/types"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/tendermint/tendermint/libs/log"
	"golang.org/x/exp/slices"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   storetypes.StoreKey
		memKey     storetypes.StoreKey
		paramstore paramtypes.Subspace
		flagKeys   *[][2]string
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
		flagKeys:   &[][2]string{},
	}
}

func (k Keeper) FlagExists(module string, name string) error {
	if slices.Contains(*k.flagKeys, [2]string{module, name}) {
		return nil
	}

	return sdkerrors.Wrapf(types.ErrFlagUnregisterd, "Flag '%s' for module '%s'", name, module)
}

func (k Keeper) RegisterKey(module string, name string) {
	*k.flagKeys = append(*k.flagKeys, [2]string{module, name})
}

func (k Keeper) InitAllStores(ctx sdk.Context) {
	for _, tup := range *k.flagKeys {
		k.SetFlag(ctx, getKey(tup[0], tup[1]), types.Flag{
			Module: tup[0],
			Name:   tup[1],
			Set:    false,
		})
	}
}

func (k Keeper) GetModuleInstance(moduleName string, flagNames []string) ModuleInstance {
	return NewModuleInstance(k, moduleName, flagNames)
}

func getKey(module string, name string) []byte {
	return []byte(module + "-" + name)
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
