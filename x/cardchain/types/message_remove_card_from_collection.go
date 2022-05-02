package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRemoveCardFromCollection = "remove_card_from_collection"

var _ sdk.Msg = &MsgRemoveCardFromCollection{}

func NewMsgRemoveCardFromCollection(creator string, collectionId uint64, cardId uint64) *MsgRemoveCardFromCollection {
	return &MsgRemoveCardFromCollection{
		Creator:      creator,
		CollectionId: collectionId,
		CardId:       cardId,
	}
}

func (msg *MsgRemoveCardFromCollection) Route() string {
	return RouterKey
}

func (msg *MsgRemoveCardFromCollection) Type() string {
	return TypeMsgRemoveCardFromCollection
}

func (msg *MsgRemoveCardFromCollection) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRemoveCardFromCollection) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRemoveCardFromCollection) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
