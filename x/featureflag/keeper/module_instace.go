package keeper

import (
	sdkerrors "cosmossdk.io/errors"
	"github.com/DecentralCardGame/Cardchain/x/featureflag/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"golang.org/x/exp/slices"
)

type ModuleInstance struct {
	moduleName string
	k          Keeper
	names      []string
}

func NewModuleInstance(ctx sdk.Context, keeper Keeper, moduleName string, names []string) ModuleInstance {
	instance := ModuleInstance{
		moduleName: moduleName,
		names:      names,
		k:          keeper,
	}
	for _, name := range names {
		val, _ := instance.Get(ctx, name)
		_ = instance.set(ctx, name, val)
	}
	return instance
}

func (m ModuleInstance) Get(ctx sdk.Context, name string) (bool, error) {
	err := m.nameExists(name)
	if err != nil {
		return false, err
	}
	return m.k.GetFlag(ctx, getKey(m.moduleName, name)).Set, nil
}

func (m ModuleInstance) set(ctx sdk.Context, name string, isSet bool) error {
	err := m.nameExists(name)
	if err != nil {
		return err
	}
	m.k.SetFlag(ctx, getKey(m.moduleName, name), types.Flag{
		Module: m.moduleName,
		Set:    isSet,
	})
	return nil
}

func (m ModuleInstance) nameExists(name string) error {
	if slices.Contains(m.names, name) {
		return nil
	}

	return sdkerrors.Wrapf(types.ErrFlagUnregisterd, "Flag '%s'", name)
}
