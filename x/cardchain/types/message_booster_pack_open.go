package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgBoosterPackOpen{}

func NewMsgBoosterPackOpen(creator string, boosterPackId uint64) *MsgBoosterPackOpen {
	return &MsgBoosterPackOpen{
		Creator:       creator,
		BoosterPackId: boosterPackId,
	}
}

func (msg *MsgBoosterPackOpen) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
