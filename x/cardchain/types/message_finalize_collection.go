package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgFinalizeCollection = "finalize_collection"

var _ sdk.Msg = &MsgFinalizeCollection{}

func NewMsgFinalizeCollection(creator string, collectionId uint64) *MsgFinalizeCollection {
	return &MsgFinalizeCollection{
		Creator:      creator,
		CollectionId: collectionId,
	}
}

func (msg *MsgFinalizeCollection) Route() string {
	return RouterKey
}

func (msg *MsgFinalizeCollection) Type() string {
	return TypeMsgFinalizeCollection
}

func (msg *MsgFinalizeCollection) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgFinalizeCollection) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgFinalizeCollection) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
