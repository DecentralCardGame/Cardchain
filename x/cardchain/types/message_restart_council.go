package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRestartCouncil = "restart_council"

var _ sdk.Msg = &MsgRestartCouncil{}

func NewMsgRestartCouncil(creator string, councilId uint64) *MsgRestartCouncil {
	return &MsgRestartCouncil{
		Creator:   creator,
		CouncilId: councilId,
	}
}

func (msg *MsgRestartCouncil) Route() string {
	return RouterKey
}

func (msg *MsgRestartCouncil) Type() string {
	return TypeMsgRestartCouncil
}

func (msg *MsgRestartCouncil) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRestartCouncil) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRestartCouncil) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
