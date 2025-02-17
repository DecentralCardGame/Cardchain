package types

import (
	"testing"

	"github.com/DecentralCardGame/cardchain/testutil/sample"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgEncounterCreate_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgEncounterCreate
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgEncounterCreate{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgEncounterCreate{
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
