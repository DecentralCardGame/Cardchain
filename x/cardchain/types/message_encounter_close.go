package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgEncounterClose = "encounter_close"

var _ sdk.Msg = &MsgEncounterClose{}

func NewMsgEncounterClose(creator string, encounterId uint64) *MsgEncounterClose {
	return &MsgEncounterClose{
		Creator:     creator,
		EncounterId: encounterId,
	}
}

func (msg *MsgEncounterClose) Route() string {
	return RouterKey
}

func (msg *MsgEncounterClose) Type() string {
	return TypeMsgEncounterClose
}

func (msg *MsgEncounterClose) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgEncounterClose) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgEncounterClose) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
