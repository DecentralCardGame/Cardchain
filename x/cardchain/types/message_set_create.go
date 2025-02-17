package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetCreate{}

func NewMsgSetCreate(creator string, name string, artist string, storyWriter string, contributors []string) *MsgSetCreate {
	return &MsgSetCreate{
		Creator:      creator,
		Name:         name,
		Artist:       artist,
		StoryWriter:  storyWriter,
		Contributors: contributors,
	}
}

func (msg *MsgSetCreate) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
