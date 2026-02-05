package keeper

import (
	"fmt"
	"slices"

	"cosmossdk.io/core/store"
	sdkerrors "cosmossdk.io/errors"
	"cosmossdk.io/log"
	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	db "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/DecentralCardGame/cardchain/x/featureflag/types"
)

const flagStoreKey = "Flags"

type (
	Keeper struct {
		cdc          codec.BinaryCodec
		storeService store.KVStoreService
		logger       log.Logger

		// the address capable of executing a MsgUpdateParams message. Typically, this
		// should be the x/gov module account.
		authority string
		flagKeys  *[][2]string
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeService store.KVStoreService,
	logger log.Logger,
	authority string,

) Keeper {
	if _, err := sdk.AccAddressFromBech32(authority); err != nil {
		panic(fmt.Sprintf("invalid authority address: %s", authority))
	}

	return Keeper{
		cdc:          cdc,
		storeService: storeService,
		authority:    authority,
		logger:       logger,
		flagKeys:     &[][2]string{},
	}
}

// GetAuthority returns the module's authority.
func (k Keeper) GetAuthority() string {
	return k.authority
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
		k.SetFlag(ctx, GetKey(tup[0], tup[1]), types.Flag{
			Module: tup[0],
			Name:   tup[1],
			Set:    false,
		})
	}
}

func (k Keeper) GetModuleInstance(moduleName string, flagNames []string) ModuleInstance {
	return NewModuleInstance(k, moduleName, flagNames)
}

func GetKey(module string, name string) []byte {
	return []byte(module + "-" + name)
}

func (k Keeper) GetFlag(ctx sdk.Context, key []byte) (gottenFlag types.Flag) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(flagStoreKey))
	b := store.Get(key)
	k.cdc.MustUnmarshal(b, &gottenFlag)
	return
}

func (k Keeper) SetFlag(ctx sdk.Context, key []byte, flag types.Flag) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(flagStoreKey))
	store.Set(key, k.cdc.MustMarshal(&flag))
}

func (k Keeper) GetFlagsIterator(ctx sdk.Context) db.Iterator {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(flagStoreKey))
	return storetypes.KVStorePrefixIterator(store, []byte{})
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

// Logger returns a module-specific logger.
func (k Keeper) Logger() log.Logger {
	return k.logger.With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
