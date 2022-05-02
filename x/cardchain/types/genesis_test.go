package types_test

import (
	"testing"

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

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

				Matches: []*types.Match{
					&someMatch,
					&someMatch,
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
