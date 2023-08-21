package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgMsgOpenMatch = "msg_open_match"

var _ sdk.Msg = &MsgMsgOpenMatch{}

func NewMsgMsgOpenMatch(creator string) *MsgMsgOpenMatch {
	return &MsgMsgOpenMatch{
		Creator: creator,
	}
}

func (msg *MsgMsgOpenMatch) Route() string {
	return RouterKey
}

func (msg *MsgMsgOpenMatch) Type() string {
	return TypeMsgMsgOpenMatch
}

func (msg *MsgMsgOpenMatch) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgMsgOpenMatch) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgMsgOpenMatch) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
