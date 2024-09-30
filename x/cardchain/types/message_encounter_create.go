package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgEncounterCreate = "encounter_create"

var _ sdk.Msg = &MsgEncounterCreate{}

func NewMsgEncounterCreate(creator string) *MsgEncounterCreate {
	return &MsgEncounterCreate{
		Creator: creator,
	}
}

func (msg *MsgEncounterCreate) Route() string {
	return RouterKey
}

func (msg *MsgEncounterCreate) Type() string {
	return TypeMsgEncounterCreate
}

func (msg *MsgEncounterCreate) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgEncounterCreate) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgEncounterCreate) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
