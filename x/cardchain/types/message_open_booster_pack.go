package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgOpenBoosterPack = "open_booster_pack"

var _ sdk.Msg = &MsgOpenBoosterPack{}

func NewMsgOpenBoosterPack(creator string, boosterPackId uint64) *MsgOpenBoosterPack {
	return &MsgOpenBoosterPack{
		Creator:       creator,
		BoosterPackId: boosterPackId,
	}
}

func (msg *MsgOpenBoosterPack) Route() string {
	return RouterKey
}

func (msg *MsgOpenBoosterPack) Type() string {
	return TypeMsgOpenBoosterPack
}

func (msg *MsgOpenBoosterPack) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgOpenBoosterPack) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgOpenBoosterPack) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
