package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgDonateToCard = "donate_to_card"

var _ sdk.Msg = &MsgDonateToCard{}

func NewMsgDonateToCard(creator string, cardId uint64, donator string, amount sdk.Coin) *MsgDonateToCard {
	return &MsgDonateToCard{
		Creator: creator,
		CardId:  cardId,
		Donator: donator,
		Amount:  amount,
	}
}

func (msg *MsgDonateToCard) Route() string {
	return RouterKey
}

func (msg *MsgDonateToCard) Type() string {
	return TypeMsgDonateToCard
}

func (msg *MsgDonateToCard) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDonateToCard) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDonateToCard) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
