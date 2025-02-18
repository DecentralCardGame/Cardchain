package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSet{}

func NewMsgSet(authority string, module string, name string, value bool) *MsgSet {
	return &MsgSet{
		Authority: authority,
		Module:    module,
		Name:      name,
		Value:     value,
	}
}

func (msg *MsgSet) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid authority address (%s)", err)
	}
	return nil
}
