package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateuser = "createuser"

var _ sdk.Msg = &MsgCreateuser{}

func NewMsgCreateuser(creator sdk.AccAddress, newuser sdk.AccAddress, alias string) *MsgCreateuser {
	return &MsgCreateuser{
		Creator: creator,
		NewUser: newuser,
		Alias:   alias,
	}
}

func (msg *MsgCreateuser) Route() string {
	return RouterKey
}

func (msg *MsgCreateuser) Type() string {
	return TypeMsgCreateuser
}

func (msg *MsgCreateuser) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Creator}
}

func (msg *MsgCreateuser) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateuser) ValidateBasic() error {
	if msg.NewUser.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.NewUser.String())
	}
	return nil
}
