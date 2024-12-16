package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCardSchemeBuy{}

func NewMsgCardSchemeBuy(creator string, bid sdk.Coin) *MsgCardSchemeBuy {
	return &MsgCardSchemeBuy{
		Creator: creator,
		Bid:     bid,
	}
}

func (msg *MsgCardSchemeBuy) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
