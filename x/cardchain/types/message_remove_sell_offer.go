package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRemoveSellOffer = "remove_sell_offer"

var _ sdk.Msg = &MsgRemoveSellOffer{}

func NewMsgRemoveSellOffer(creator string, sellOfferId uint64) *MsgRemoveSellOffer {
	return &MsgRemoveSellOffer{
		Creator:     creator,
		SellOfferId: sellOfferId,
	}
}

func (msg *MsgRemoveSellOffer) Route() string {
	return RouterKey
}

func (msg *MsgRemoveSellOffer) Type() string {
	return TypeMsgRemoveSellOffer
}

func (msg *MsgRemoveSellOffer) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRemoveSellOffer) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRemoveSellOffer) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
