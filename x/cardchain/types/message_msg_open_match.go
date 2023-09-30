package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgOpenMatch = "msg_open_match"

var _ sdk.Msg = &MsgOpenMatch{}

func NewMsgOpenMatch(creator string) *MsgOpenMatch {
	return &MsgOpenMatch{
		Creator: creator,
	}
}

func (msg *MsgOpenMatch) Route() string {
	return RouterKey
}

func (msg *MsgOpenMatch) Type() string {
	return TypeMsgOpenMatch
}

func (msg *MsgOpenMatch) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgOpenMatch) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgOpenMatch) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
