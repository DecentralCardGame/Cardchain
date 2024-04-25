package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgDisinviteEarlyAccess = "disinvite_early_access"

var _ sdk.Msg = &MsgDisinviteEarlyAccess{}

func NewMsgDisinviteEarlyAccess(creator string, user string) *MsgDisinviteEarlyAccess {
	return &MsgDisinviteEarlyAccess{
		Creator: creator,
		User:    user,
	}
}

func (msg *MsgDisinviteEarlyAccess) Route() string {
	return RouterKey
}

func (msg *MsgDisinviteEarlyAccess) Type() string {
	return TypeMsgDisinviteEarlyAccess
}

func (msg *MsgDisinviteEarlyAccess) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDisinviteEarlyAccess) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDisinviteEarlyAccess) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
