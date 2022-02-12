package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSubmitCollectionProposal = "submit_collection_proposal"

var _ sdk.Msg = &MsgSubmitCollectionProposal{}

func NewMsgSubmitCollectionProposal(creator string, collectionId uint64) *MsgSubmitCollectionProposal {
	return &MsgSubmitCollectionProposal{
		Creator:      creator,
		CollectionId: collectionId,
	}
}

func (msg *MsgSubmitCollectionProposal) Route() string {
	return RouterKey
}

func (msg *MsgSubmitCollectionProposal) Type() string {
	return TypeMsgSubmitCollectionProposal
}

func (msg *MsgSubmitCollectionProposal) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSubmitCollectionProposal) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSubmitCollectionProposal) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
