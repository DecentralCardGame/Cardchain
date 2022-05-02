package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgBuyCard = "buy_card"

var _ sdk.Msg = &MsgBuyCard{}

func NewMsgBuyCard(creator string, sellOfferId uint64) *MsgBuyCard {
	return &MsgBuyCard{
		Creator:     creator,
		SellOfferId: sellOfferId,
	}
}

func (msg *MsgBuyCard) Route() string {
	return RouterKey
}

func (msg *MsgBuyCard) Type() string {
	return TypeMsgBuyCard
}

func (msg *MsgBuyCard) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgBuyCard) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgBuyCard) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
