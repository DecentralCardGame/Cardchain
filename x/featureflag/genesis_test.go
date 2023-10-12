package featureflag_test

import (
	"testing"

	keepertest "github.com/DecentralCardGame/Cardchain/testutil/keeper"
	"github.com/DecentralCardGame/Cardchain/testutil/nullify"
	"github.com/DecentralCardGame/Cardchain/x/featureflag"
	"github.com/DecentralCardGame/Cardchain/x/featureflag/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		FlagsList: []types.Flags{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		FlagsList: []types.Flags{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.FeatureflagKeeper(t)
	featureflag.InitGenesis(ctx, *k, genesisState)
	got := featureflag.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.FlagsList, got.FlagsList)
	require.ElementsMatch(t, genesisState.FlagsList, got.FlagsList)
	// this line is used by starport scaffolding # genesis/test/assert
}
