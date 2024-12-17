package types_test

import (
	"testing"

	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	tests := []struct {
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

				ProductDetailsList: []types.ProductDetails{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				ProductDetailsCount: 2,
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated productDetails",
			genState: &types.GenesisState{
				ProductDetailsList: []types.ProductDetails{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid productDetails count",
			genState: &types.GenesisState{
				ProductDetailsList: []types.ProductDetails{
					{
						Id: 1,
					},
				},
				ProductDetailsCount: 0,
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	}
	for _, tc := range tests {
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
