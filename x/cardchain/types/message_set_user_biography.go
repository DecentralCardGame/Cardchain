package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetUserBiography = "set_user_biography"

var _ sdk.Msg = &MsgSetUserBiography{}

func NewMsgSetUserBiography(creator string, biography string) *MsgSetUserBiography {
	return &MsgSetUserBiography{
		Creator:   creator,
		Biography: biography,
	}
}

func (msg *MsgSetUserBiography) Route() string {
	return RouterKey
}

func (msg *MsgSetUserBiography) Type() string {
	return TypeMsgSetUserBiography
}

func (msg *MsgSetUserBiography) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetUserBiography) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetUserBiography) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
