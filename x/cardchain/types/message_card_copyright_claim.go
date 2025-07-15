package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCardCopyrightClaim{}

func NewMsgCardCopyrightClaim(authority string, cardId uint64) *MsgCardCopyrightClaim {
	return &MsgCardCopyrightClaim{
		Authority: authority,
		CardId:    cardId,
	}
}

func (msg *MsgCardCopyrightClaim) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid authority address (%s)", err)
	}
	return nil
}
