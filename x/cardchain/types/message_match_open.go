package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgMatchOpen{}

func NewMsgMatchOpen(creator string, playerA string, playerB string, playerAdeck []uint64, playerBdeck []uint64) *MsgMatchOpen {
	return &MsgMatchOpen{
		Creator:     creator,
		PlayerA:     playerA,
		PlayerB:     playerB,
		PlayerAdeck: playerAdeck,
		PlayerBdeck: playerBdeck,
	}
}

func (msg *MsgMatchOpen) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
