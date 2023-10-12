package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type ModuleInstance struct {
	moduleName string
	k          Keeper
}

func NewModuleInstance(keeper Keeper, moduleName string, names []string) ModuleInstance {
	instance := ModuleInstance{
		moduleName: moduleName,
		k:          keeper,
	}
	for _, name := range names {
		keeper.RegisterKey(moduleName, name)
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

func (m ModuleInstance) nameExists(name string) error {
	return m.k.FlagExists(m.moduleName, name)
}
