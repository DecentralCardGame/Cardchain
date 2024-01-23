package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgChangeAlias = "change_alias"

var _ sdk.Msg = &MsgChangeAlias{}

func NewMsgChangeAlias(creator string, alias string) *MsgChangeAlias {
	return &MsgChangeAlias{
		Creator: creator,
		Alias:   alias,
	}
}

func (msg *MsgChangeAlias) Route() string {
	return RouterKey
}

func (msg *MsgChangeAlias) Type() string {
	return TypeMsgChangeAlias
}

func (msg *MsgChangeAlias) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgChangeAlias) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgChangeAlias) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
