package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetCollectionName = "set_collection_name"

var _ sdk.Msg = &MsgSetCollectionName{}

func NewMsgSetCollectionName(creator string, collectionId uint64, name string) *MsgSetCollectionName {
	return &MsgSetCollectionName{
		Creator:      creator,
		CollectionId: collectionId,
		Name:         name,
	}
}

func (msg *MsgSetCollectionName) Route() string {
	return RouterKey
}

func (msg *MsgSetCollectionName) Type() string {
	return TypeMsgSetCollectionName
}

func (msg *MsgSetCollectionName) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetCollectionName) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetCollectionName) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
