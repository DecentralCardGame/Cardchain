package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRemoveContributorFromCollection = "remove_contributor_from_collection"

var _ sdk.Msg = &MsgRemoveContributorFromCollection{}

func NewMsgRemoveContributorFromCollection(creator string, collectionId uint64, user string) *MsgRemoveContributorFromCollection {
	return &MsgRemoveContributorFromCollection{
		Creator:      creator,
		CollectionId: collectionId,
		User:         user,
	}
}

func (msg *MsgRemoveContributorFromCollection) Route() string {
	return RouterKey
}

func (msg *MsgRemoveContributorFromCollection) Type() string {
	return TypeMsgRemoveContributorFromCollection
}

func (msg *MsgRemoveContributorFromCollection) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRemoveContributorFromCollection) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRemoveContributorFromCollection) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
