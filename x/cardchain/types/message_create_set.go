package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateSet = "create_set"

var _ sdk.Msg = &MsgCreateSet{}

func NewMsgCreateSet(creator string, name string, artist string, storyWriter string, contributors []string) *MsgCreateSet {
	return &MsgCreateSet{
		Creator:      creator,
		Name:         name,
		Artist:       artist,
		StoryWriter:  storyWriter,
		Contributors: contributors,
	}
}

func (msg *MsgCreateSet) Route() string {
	return RouterKey
}

func (msg *MsgCreateSet) Type() string {
	return TypeMsgCreateSet
}

func (msg *MsgCreateSet) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateSet) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateSet) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(errors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
