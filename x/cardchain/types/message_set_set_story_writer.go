package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetSetStoryWriter = "set_set_story_writer"

var _ sdk.Msg = &MsgSetSetStoryWriter{}

func NewMsgSetSetStoryWriter(creator string, setId uint64, storyWriter string) *MsgSetSetStoryWriter {
	return &MsgSetSetStoryWriter{
		Creator:     creator,
		SetId:       setId,
		StoryWriter: storyWriter,
	}
}

func (msg *MsgSetSetStoryWriter) Route() string {
	return RouterKey
}

func (msg *MsgSetSetStoryWriter) Type() string {
	return TypeMsgSetSetStoryWriter
}

func (msg *MsgSetSetStoryWriter) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetSetStoryWriter) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetSetStoryWriter) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(errors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
