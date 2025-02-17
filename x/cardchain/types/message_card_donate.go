package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCardDonate{}

func NewMsgCardDonate(creator string, cardId uint64, amount sdk.Coin) *MsgCardDonate {
	return &MsgCardDonate{
		Creator: creator,
		CardId:  cardId,
		Amount:  amount,
	}
}

func (msg *MsgCardDonate) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
