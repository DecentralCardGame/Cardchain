package cardchain_test

import (
	"testing"

	keepertest "github.com/DecentralCardGame/Cardchain/testutil/keeper"
	"github.com/DecentralCardGame/Cardchain/testutil/nullify"
	"github.com/DecentralCardGame/Cardchain/x/cardchain"
	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	"github.com/stretchr/testify/require"
)

var someMatch = types.Match{
	100,
	"cc1cyezs5v34utk48l3mgm8v8ll2d286xhs7apu0d",
	"cc1cyezs5v34utk48l3mgm8v8ll2d286xhs7apu0d",
	"cc1cyezs5v34utk48l3mgm8v8ll2d286xhs7apu0d",
	types.Outcome_AWon,
}

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		Matches: []*types.Match{
			&someMatch,
			&someMatch,
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CardchainKeeper(t)
	cardchain.InitGenesis(ctx, *k, genesisState)
	got := cardchain.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.Matches, got.Matches)
	// this line is used by starport scaffolding # genesis/test/assert
}
