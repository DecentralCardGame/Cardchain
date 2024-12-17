package cardchain

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/DecentralCardGame/cardchain/x/cardchain/keeper"
	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the productDetails
	for _, elem := range genState.ProductDetailsList {
		k.SetProductDetails(ctx, elem)
	}

	// Set productDetails count
	k.SetProductDetailsCount(ctx, genState.ProductDetailsCount)
	// this line is used by starport scaffolding # genesis/module/init
	if err := k.SetParams(ctx, genState.Params); err != nil {
		panic(err)
	}
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.ProductDetailsList = k.GetAllProductDetails(ctx)
	genesis.ProductDetailsCount = k.GetProductDetailsCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
