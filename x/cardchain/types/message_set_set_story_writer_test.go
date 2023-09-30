package types

import (
	"testing"

	"github.com/DecentralCardGame/Cardchain/testutil/sample"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgSetSetStoryWriter_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgSetSetStoryWriter
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgSetSetStoryWriter{
				Creator: "invalid_address",
			},
			err: errors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgSetSetStoryWriter{
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
