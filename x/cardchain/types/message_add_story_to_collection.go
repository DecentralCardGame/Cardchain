package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAddStoryToCollection = "add_story_to_collection"

var _ sdk.Msg = &MsgAddStoryToCollection{}

func NewMsgAddStoryToCollection(creator string, collectionId uint64, story string) *MsgAddStoryToCollection {
	return &MsgAddStoryToCollection{
		Creator:      creator,
		CollectionId: collectionId,
		Story:        story,
	}
}

func (msg *MsgAddStoryToCollection) Route() string {
	return RouterKey
}

func (msg *MsgAddStoryToCollection) Type() string {
	return TypeMsgAddStoryToCollection
}

func (msg *MsgAddStoryToCollection) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddStoryToCollection) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddStoryToCollection) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
