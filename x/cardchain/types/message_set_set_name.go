package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetSetName = "set_set_name"

var _ sdk.Msg = &MsgSetSetName{}

func NewMsgSetSetName(creator string, setId uint64, name string) *MsgSetSetName {
	return &MsgSetSetName{
		Creator: creator,
		SetId:   setId,
		Name:    name,
	}
}

func (msg *MsgSetSetName) Route() string {
	return RouterKey
}

func (msg *MsgSetSetName) Type() string {
	return TypeMsgSetSetName
}

func (msg *MsgSetSetName) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetSetName) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetSetName) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
