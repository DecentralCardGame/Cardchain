package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetContributorRemove{}

func NewMsgSetContributorRemove(creator string, setId uint64, user string) *MsgSetContributorRemove {
	return &MsgSetContributorRemove{
		Creator: creator,
		SetId:   setId,
		User:    user,
	}
}

func (msg *MsgSetContributorRemove) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
