package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCardArtistChange{}

func NewMsgCardArtistChange(creator string, cardId uint64, artist string) *MsgCardArtistChange {
	return &MsgCardArtistChange{
		Creator: creator,
		CardId:  cardId,
		Artist:  artist,
	}
}

func (msg *MsgCardArtistChange) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
