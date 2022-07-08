package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgTransferBoosterPack = "transfer_booster_pack"

var _ sdk.Msg = &MsgTransferBoosterPack{}

func NewMsgTransferBoosterPack(creator string, boosterPackId uint64, receiver string) *MsgTransferBoosterPack {
	return &MsgTransferBoosterPack{
		Creator:       creator,
		BoosterPackId: boosterPackId,
		Receiver:      receiver,
	}
}

func (msg *MsgTransferBoosterPack) Route() string {
	return RouterKey
}

func (msg *MsgTransferBoosterPack) Type() string {
	return TypeMsgTransferBoosterPack
}

func (msg *MsgTransferBoosterPack) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgTransferBoosterPack) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgTransferBoosterPack) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
