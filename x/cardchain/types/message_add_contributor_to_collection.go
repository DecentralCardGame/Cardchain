package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAddContributorToCollection = "add_contributor_to_collection"

var _ sdk.Msg = &MsgAddContributorToCollection{}

func NewMsgAddContributorToCollection(creator string, collectionId uint64, user string) *MsgAddContributorToCollection {
	return &MsgAddContributorToCollection{
		Creator:      creator,
		CollectionId: collectionId,
		User:         user,
	}
}

func (msg *MsgAddContributorToCollection) Route() string {
	return RouterKey
}

func (msg *MsgAddContributorToCollection) Type() string {
	return TypeMsgAddContributorToCollection
}

func (msg *MsgAddContributorToCollection) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddContributorToCollection) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddContributorToCollection) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
