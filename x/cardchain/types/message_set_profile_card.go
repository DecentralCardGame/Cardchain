package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetProfileCard = "set_profile_card"

var _ sdk.Msg = &MsgSetProfileCard{}

func NewMsgSetProfileCard(creator string, cardId uint64) *MsgSetProfileCard {
	return &MsgSetProfileCard{
		Creator: creator,
		CardId:  cardId,
	}
}

func (msg *MsgSetProfileCard) Route() string {
	return RouterKey
}

func (msg *MsgSetProfileCard) Type() string {
	return TypeMsgSetProfileCard
}

func (msg *MsgSetProfileCard) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetProfileCard) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetProfileCard) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
