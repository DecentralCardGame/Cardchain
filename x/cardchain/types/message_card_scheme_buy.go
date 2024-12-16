package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCardSchemeBuy = "card_scheme_buy"

var _ sdk.Msg = &MsgCardSchemeBuy{}

func NewMsgCardSchemeBuy(creator string, bid sdk.Coin) *MsgCardSchemeBuy {
	return &MsgCardSchemeBuy{
		Creator: creator,
		Bid:     bid,
	}
}

func (msg *MsgCardSchemeBuy) Route() string {
	return RouterKey
}

func (msg *MsgCardSchemeBuy) Type() string {
	return TypeMsgCardSchemeBuy
}

func (msg *MsgCardSchemeBuy) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCardSchemeBuy) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCardSchemeBuy) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
