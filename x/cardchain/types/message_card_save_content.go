package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCardSaveContent = "card_save_content"

var _ sdk.Msg = &MsgCardSaveContent{}

func NewMsgCardSaveContent(creator string, cardId uint64, content string, notes string, artist string, balanceAnchor bool) *MsgCardSaveContent {
	return &MsgCardSaveContent{
		Creator:       creator,
		CardId:        cardId,
		Content:       content,
		Notes:         notes,
		Artist:        artist,
		BalanceAnchor: balanceAnchor,
	}
}

func (msg *MsgCardSaveContent) Route() string {
	return RouterKey
}

func (msg *MsgCardSaveContent) Type() string {
	return TypeMsgCardSaveContent
}

func (msg *MsgCardSaveContent) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCardSaveContent) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCardSaveContent) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
