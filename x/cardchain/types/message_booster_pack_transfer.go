package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgBoosterPackTransfer{}

func NewMsgBoosterPackTransfer(creator string, boosterPackId uint64, receiver string) *MsgBoosterPackTransfer {
	return &MsgBoosterPackTransfer{
		Creator:       creator,
		BoosterPackId: boosterPackId,
		Receiver:      receiver,
	}
}

func (msg *MsgBoosterPackTransfer) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
