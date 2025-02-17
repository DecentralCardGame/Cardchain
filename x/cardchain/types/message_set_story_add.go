package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetStoryAdd{}

func NewMsgSetStoryAdd(creator string, setId uint64, story string) *MsgSetStoryAdd {
	return &MsgSetStoryAdd{
		Creator: creator,
		SetId:   setId,
		Story:   story,
	}
}

func (msg *MsgSetStoryAdd) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
