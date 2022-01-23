package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateuser = "createuser"

var _ sdk.Msg = &MsgCreateuser{}

func NewMsgCreateuser(creator string, newuser string, alias string) *MsgCreateuser {
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

func (msg *MsgCreateuser) GetNewUserAddr() sdk.AccAddress {
	user, err := sdk.AccAddressFromBech32(msg.NewUser)
	if err != nil {
		panic(err)
	}
	return user
}

func (msg *MsgCreateuser) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateuser) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateuser) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
