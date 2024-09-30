package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgEncounterDo = "encounter_do"

var _ sdk.Msg = &MsgEncounterDo{}

func NewMsgEncounterDo(creator string, encounterId uint64) *MsgEncounterDo {
	return &MsgEncounterDo{
		Creator:     creator,
		EncounterId: encounterId,
	}
}

func (msg *MsgEncounterDo) Route() string {
	return RouterKey
}

func (msg *MsgEncounterDo) Type() string {
	return TypeMsgEncounterDo
}

func (msg *MsgEncounterDo) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgEncounterDo) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgEncounterDo) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
