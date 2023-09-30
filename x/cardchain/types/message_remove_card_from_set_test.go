package types

import (
	"testing"

	"github.com/DecentralCardGame/Cardchain/testutil/sample"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgRemoveCardFromSet_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgRemoveCardFromSet
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgRemoveCardFromSet{
				Creator: "invalid_address",
			},
			err: errors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgRemoveCardFromSet{
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
