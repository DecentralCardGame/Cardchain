package cardchain_test

import (
	"testing"

	keepertest "github.com/DecentralCardGame/cardchain/testutil/keeper"
	"github.com/DecentralCardGame/cardchain/testutil/nullify"
	cardchain "github.com/DecentralCardGame/cardchain/x/cardchain/module"
	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		ProductDetailsList: []types.ProductDetails{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		ProductDetailsCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CardchainKeeper(t)
	cardchain.InitGenesis(ctx, k, genesisState)
	got := cardchain.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.ProductDetailsList, got.ProductDetailsList)
	require.Equal(t, genesisState.ProductDetailsCount, got.ProductDetailsCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
