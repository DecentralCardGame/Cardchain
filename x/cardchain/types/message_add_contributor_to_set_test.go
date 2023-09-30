package types

import (
	"testing"

	"github.com/DecentralCardGame/Cardchain/testutil/sample"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgAddContributorToSet_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgAddContributorToSet
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgAddContributorToSet{
				Creator: "invalid_address",
			},
			err: errors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgAddContributorToSet{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
