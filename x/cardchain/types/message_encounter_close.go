package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgEncounterClose{}

func NewMsgEncounterClose(creator string, encounterId uint64, user string, won bool) *MsgEncounterClose {
	return &MsgEncounterClose{
		Creator:     creator,
		EncounterId: encounterId,
		User:        user,
		Won:         won,
	}
}

func (msg *MsgEncounterClose) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
