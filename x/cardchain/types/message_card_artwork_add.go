package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCardArtworkAdd{}

func NewMsgCardArtworkAdd(creator string, cardId uint64, image []byte, fullArt bool) *MsgCardArtworkAdd {
	return &MsgCardArtworkAdd{
		Creator: creator,
		CardId:  cardId,
		Image:   image,
		FullArt: fullArt,
	}
}

func (msg *MsgCardArtworkAdd) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
