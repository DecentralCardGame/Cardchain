package featureflag_test

import (
	"testing"

	keepertest "github.com/DecentralCardGame/cardchain/testutil/keeper"
	"github.com/DecentralCardGame/cardchain/testutil/nullify"
	featureflag "github.com/DecentralCardGame/cardchain/x/featureflag/module"
	"github.com/DecentralCardGame/cardchain/x/featureflag/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.FeatureflagKeeper(t)
	featureflag.InitGenesis(ctx, k, genesisState)
	got := featureflag.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
