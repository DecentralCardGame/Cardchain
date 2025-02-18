package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgEarlyAccessGrant{}

func NewMsgEarlyAccessGrant(authority string, users []string) *MsgEarlyAccessGrant {
	return &MsgEarlyAccessGrant{
		Authority: authority,
		Users:     users,
	}
}

func (msg *MsgEarlyAccessGrant) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid authority address (%s)", err)
	}
	return nil
}
