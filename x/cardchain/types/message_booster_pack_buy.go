package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgBoosterPackBuy{}

func NewMsgBoosterPackBuy(creator string, setId uint64) *MsgBoosterPackBuy {
	return &MsgBoosterPackBuy{
		Creator: creator,
		SetId:   setId,
	}
}

func (msg *MsgBoosterPackBuy) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
