package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateCollection = "create_collection"

var _ sdk.Msg = &MsgCreateCollection{}

func NewMsgCreateCollection(creator string, name string, artist string, storyWriter string, contributors []string) *MsgCreateCollection {
	return &MsgCreateCollection{
		Creator:      creator,
		Name:         name,
		Artist:				artist,
		StoryWriter:	storyWriter,
		Contributors: contributors,
	}
}

func (msg *MsgCreateCollection) Route() string {
	return RouterKey
}

func (msg *MsgCreateCollection) Type() string {
	return TypeMsgCreateCollection
}

func (msg *MsgCreateCollection) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateCollection) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateCollection) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
