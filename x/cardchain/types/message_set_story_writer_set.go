package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetStoryWriterSet{}

func NewMsgSetStoryWriterSet(creator string, setId uint64, storyWriter string) *MsgSetStoryWriterSet {
	return &MsgSetStoryWriterSet{
		Creator:     creator,
		SetId:       setId,
		StoryWriter: storyWriter,
	}
}

func (msg *MsgSetStoryWriterSet) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
