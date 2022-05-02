package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgTransferCard = "transfer_card"

var _ sdk.Msg = &MsgTransferCard{}

func NewMsgTransferCard(creator string, cardId uint64, receiver string) *MsgTransferCard {
	return &MsgTransferCard{
		Creator:  creator,
		CardId:   cardId,
		Receiver: receiver,
	}
}

func (msg *MsgTransferCard) Route() string {
	return RouterKey
}

func (msg *MsgTransferCard) Type() string {
	return TypeMsgTransferCard
}

func (msg *MsgTransferCard) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgTransferCard) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgTransferCard) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
