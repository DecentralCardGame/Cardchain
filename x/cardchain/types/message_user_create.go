package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUserCreate = "user_create"

var _ sdk.Msg = &MsgUserCreate{}

func NewMsgUserCreate(creator string, newUser string, alias string) *MsgUserCreate {
	return &MsgUserCreate{
		Creator: creator,
		NewUser: newUser,
		Alias:   alias,
	}
}

func (msg *MsgUserCreate) Route() string {
	return RouterKey
}

func (msg *MsgUserCreate) Type() string {
	return TypeMsgUserCreate
}

func (msg *MsgUserCreate) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUserCreate) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUserCreate) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
