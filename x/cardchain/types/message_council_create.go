package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCouncilCreate{}

func NewMsgCouncilCreate(creator string, cardId uint64) *MsgCouncilCreate {
	return &MsgCouncilCreate{
		Creator: creator,
		CardId:  cardId,
	}
}

func (msg *MsgCouncilCreate) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
