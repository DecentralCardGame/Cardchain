package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSaveCardContent = "save_card_content"

var _ sdk.Msg = &MsgSaveCardContent{}

func NewMsgSaveCardContent(creator string, cardId uint64, content []byte, notes string, artist string) *MsgSaveCardContent {
	return &MsgSaveCardContent{
		Creator: creator,
		CardId:  cardId,
		Content: content,
		Notes:   notes,
		Artist:  artist,
	}
}

func (msg *MsgSaveCardContent) Route() string {
	return RouterKey
}

func (msg *MsgSaveCardContent) Type() string {
	return TypeMsgSaveCardContent
}

func (msg *MsgSaveCardContent) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSaveCardContent) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSaveCardContent) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(errors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
