package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCardSaveContent{}

func NewMsgCardSaveContent(creator string, cardId uint64, content string, notes string, artist string, balanceAnchor bool) *MsgCardSaveContent {
	return &MsgCardSaveContent{
		Creator:       creator,
		CardId:        cardId,
		Content:       content,
		Notes:         notes,
		Artist:        artist,
		BalanceAnchor: balanceAnchor,
	}
}

func (msg *MsgCardSaveContent) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
