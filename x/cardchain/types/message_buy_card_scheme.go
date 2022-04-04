package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgBuyCardScheme = "buy_card_scheme"

var _ sdk.Msg = &MsgBuyCardScheme{}

func NewMsgBuyCardScheme(creator string, bid uint64) *MsgBuyCardScheme {
	return &MsgBuyCardScheme{
		Creator: creator,
		Bid:     bid,
	}
}

func (msg *MsgBuyCardScheme) Route() string {
	return RouterKey
}

func (msg *MsgBuyCardScheme) Type() string {
	return TypeMsgBuyCardScheme
}

func (msg *MsgBuyCardScheme) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgBuyCardScheme) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgBuyCardScheme) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
