package featureflag

import (
	"github.com/DecentralCardGame/Cardchain/x/featureflag/keeper"
	"github.com/DecentralCardGame/Cardchain/x/featureflag/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
	for key, val := range genState.Flags {
		k.SetFlag(ctx, []byte(key), *val)
	}
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)
	flags, keys := k.GetAllFlags(ctx)
	for idx, key := range keys {
		genesis.Flags[key] = flags[idx]
	}

// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
