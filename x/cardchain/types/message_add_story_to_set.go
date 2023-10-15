package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAddStoryToSet = "add_story_to_set"

var _ sdk.Msg = &MsgAddStoryToSet{}

func NewMsgAddStoryToSet(creator string, setId uint64, story string) *MsgAddStoryToSet {
	return &MsgAddStoryToSet{
		Creator: creator,
		SetId:   setId,
		Story:   story,
	}
}

func (msg *MsgAddStoryToSet) Route() string {
	return RouterKey
}

func (msg *MsgAddStoryToSet) Type() string {
	return TypeMsgAddStoryToSet
}

func (msg *MsgAddStoryToSet) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddStoryToSet) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddStoryToSet) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(errors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
