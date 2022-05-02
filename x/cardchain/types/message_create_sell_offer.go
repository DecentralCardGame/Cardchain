package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateSellOffer = "create_sell_offer"

var _ sdk.Msg = &MsgCreateSellOffer{}

func NewMsgCreateSellOffer(creator string, card uint64, price sdk.Coin) *MsgCreateSellOffer {
	return &MsgCreateSellOffer{
		Creator: creator,
		Card:    card,
		Price:   price,
	}
}

func (msg *MsgCreateSellOffer) Route() string {
	return RouterKey
}

func (msg *MsgCreateSellOffer) Type() string {
	return TypeMsgCreateSellOffer
}

func (msg *MsgCreateSellOffer) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateSellOffer) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateSellOffer) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
