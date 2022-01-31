package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAddArtwork = "add_artwork"

var _ sdk.Msg = &MsgAddArtwork{}

func NewMsgAddArtwork(creator string, cardId uint64, image []byte, fullArt bool) *MsgAddArtwork {
	return &MsgAddArtwork{
		Creator: creator,
		CardId: cardId,
		Image:   image,
		FullArt: fullArt,
	}
}

func (msg *MsgAddArtwork) Route() string {
	return RouterKey
}

func (msg *MsgAddArtwork) Type() string {
	return TypeMsgAddArtwork
}

func (msg *MsgAddArtwork) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddArtwork) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddArtwork) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
