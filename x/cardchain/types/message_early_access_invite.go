package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgEarlyAccessInvite{}

func NewMsgEarlyAccessInvite(creator string, user string) *MsgEarlyAccessInvite {
	return &MsgEarlyAccessInvite{
		Creator: creator,
		User:    user,
	}
}

func (msg *MsgEarlyAccessInvite) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
