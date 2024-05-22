package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgInviteEarlyAccess = "invite_early_access"

var _ sdk.Msg = &MsgInviteEarlyAccess{}

func NewMsgInviteEarlyAccess(creator string, user string) *MsgInviteEarlyAccess {
	return &MsgInviteEarlyAccess{
		Creator: creator,
		User:    user,
	}
}

func (msg *MsgInviteEarlyAccess) Route() string {
	return RouterKey
}

func (msg *MsgInviteEarlyAccess) Type() string {
	return TypeMsgInviteEarlyAccess
}

func (msg *MsgInviteEarlyAccess) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgInviteEarlyAccess) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgInviteEarlyAccess) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
