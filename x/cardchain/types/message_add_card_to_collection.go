package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAddCardToCollection = "add_card_to_collection"

var _ sdk.Msg = &MsgAddCardToCollection{}

func NewMsgAddCardToCollection(creator string, collectionId uint64, cardId uint64) *MsgAddCardToCollection {
	return &MsgAddCardToCollection{
		Creator:      creator,
		CollectionId: collectionId,
		CardId:       cardId,
	}
}

func (msg *MsgAddCardToCollection) Route() string {
	return RouterKey
}

func (msg *MsgAddCardToCollection) Type() string {
	return TypeMsgAddCardToCollection
}

func (msg *MsgAddCardToCollection) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddCardToCollection) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddCardToCollection) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
