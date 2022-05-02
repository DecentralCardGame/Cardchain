package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRewokeCouncilRegistration = "rewoke_council_registration"

var _ sdk.Msg = &MsgRewokeCouncilRegistration{}

func NewMsgRewokeCouncilRegistration(creator string) *MsgRewokeCouncilRegistration {
	return &MsgRewokeCouncilRegistration{
		Creator: creator,
	}
}

func (msg *MsgRewokeCouncilRegistration) Route() string {
	return RouterKey
}

func (msg *MsgRewokeCouncilRegistration) Type() string {
	return TypeMsgRewokeCouncilRegistration
}

func (msg *MsgRewokeCouncilRegistration) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRewokeCouncilRegistration) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRewokeCouncilRegistration) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
