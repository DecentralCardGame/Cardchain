package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSellOfferCreate{}

func NewMsgSellOfferCreate(creator string, cardId uint64, price sdk.Coin) *MsgSellOfferCreate {
	return &MsgSellOfferCreate{
		Creator: creator,
		CardId:  cardId,
		Price:   price,
	}
}

func (msg *MsgSellOfferCreate) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
