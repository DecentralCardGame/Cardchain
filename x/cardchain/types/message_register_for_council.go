package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRegisterForCouncil = "register_for_council"

var _ sdk.Msg = &MsgRegisterForCouncil{}

func NewMsgRegisterForCouncil(creator string) *MsgRegisterForCouncil {
	return &MsgRegisterForCouncil{
		Creator: creator,
	}
}

func (msg *MsgRegisterForCouncil) Route() string {
	return RouterKey
}

func (msg *MsgRegisterForCouncil) Type() string {
	return TypeMsgRegisterForCouncil
}

func (msg *MsgRegisterForCouncil) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRegisterForCouncil) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRegisterForCouncil) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(errors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
