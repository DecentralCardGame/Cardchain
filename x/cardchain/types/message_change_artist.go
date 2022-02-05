package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgChangeArtist = "change_artist"

var _ sdk.Msg = &MsgChangeArtist{}

func NewMsgChangeArtist(creator string, cardID uint64, artist string) *MsgChangeArtist {
	return &MsgChangeArtist{
		Creator: creator,
		CardID:  cardID,
		Artist:  artist,
	}
}

func (msg *MsgChangeArtist) Route() string {
	return RouterKey
}

func (msg *MsgChangeArtist) Type() string {
	return TypeMsgChangeArtist
}

func (msg *MsgChangeArtist) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgChangeArtist) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgChangeArtist) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
