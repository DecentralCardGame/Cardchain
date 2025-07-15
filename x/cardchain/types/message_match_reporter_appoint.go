package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgMatchReporterAppoint{}

func NewMsgMatchReporterAppoint(authority string, reporter string) *MsgMatchReporterAppoint {
	return &MsgMatchReporterAppoint{
		Authority: authority,
		Reporter:  reporter,
	}
}

func (msg *MsgMatchReporterAppoint) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
