package types

import (
	"testing"

	"github.com/DecentralCardGame/Cardchain/testutil/sample"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgAddStoryToSet_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgAddStoryToSet
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgAddStoryToSet{
				Creator: "invalid_address",
			},
		}, {
			name: "valid address",
			err:  errors.ErrInvalidAddress,
			msg: MsgAddStoryToSet{
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
