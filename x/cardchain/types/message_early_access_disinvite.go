package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgEarlyAccessDisinvite{}

func NewMsgEarlyAccessDisinvite(creator string, user string) *MsgEarlyAccessDisinvite {
	return &MsgEarlyAccessDisinvite{
		Creator: creator,
		User:    user,
	}
}

func (msg *MsgEarlyAccessDisinvite) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
