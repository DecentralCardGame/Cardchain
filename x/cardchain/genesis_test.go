package cardchain_test

import (
	"testing"

	keepertest "github.com/DecentralCardGame/Cardchain/testutil/keeper"
	"github.com/DecentralCardGame/Cardchain/testutil/nullify"
	"github.com/DecentralCardGame/Cardchain/x/cardchain"
	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CardchainKeeper(t)
	cardchain.InitGenesis(ctx, *k, genesisState)
	got := cardchain.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
