package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgEncounterDo{}

func NewMsgEncounterDo(creator string, encounterId uint64, user string) *MsgEncounterDo {
	return &MsgEncounterDo{
		Creator:     creator,
		EncounterId: encounterId,
		User:        user,
	}
}

func (msg *MsgEncounterDo) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
