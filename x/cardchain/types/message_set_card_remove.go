package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetCardRemove{}

func NewMsgSetCardRemove(creator string, setId uint64, cardId uint64) *MsgSetCardRemove {
	return &MsgSetCardRemove{
		Creator: creator,
		SetId:   setId,
		CardId:  cardId,
	}
}

func (msg *MsgSetCardRemove) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
