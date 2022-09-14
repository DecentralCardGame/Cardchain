package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetCollectionStoryWriter = "set_collection_story_writer"

var _ sdk.Msg = &MsgSetCollectionStoryWriter{}

func NewMsgSetCollectionStoryWriter(creator string, collectionId uint64, storyWriter string) *MsgSetCollectionStoryWriter {
	return &MsgSetCollectionStoryWriter{
		Creator:      creator,
		CollectionId: collectionId,
		StoryWriter:  storyWriter,
	}
}

func (msg *MsgSetCollectionStoryWriter) Route() string {
	return RouterKey
}

func (msg *MsgSetCollectionStoryWriter) Type() string {
	return TypeMsgSetCollectionStoryWriter
}

func (msg *MsgSetCollectionStoryWriter) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetCollectionStoryWriter) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetCollectionStoryWriter) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
