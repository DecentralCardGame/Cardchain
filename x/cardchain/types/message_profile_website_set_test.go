package types

import (
	"testing"

	"github.com/DecentralCardGame/cardchain/testutil/sample"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgProfileWebsiteSet_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgProfileWebsiteSet
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgProfileWebsiteSet{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgProfileWebsiteSet{
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
