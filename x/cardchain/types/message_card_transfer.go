package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCardTransfer{}

func NewMsgCardTransfer(creator string, cardId uint64, receiver string) *MsgCardTransfer {
	return &MsgCardTransfer{
		Creator:  creator,
		CardId:   cardId,
		Receiver: receiver,
	}
}

func (msg *MsgCardTransfer) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
