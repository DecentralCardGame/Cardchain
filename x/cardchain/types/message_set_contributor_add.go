package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetContributorAdd{}

func NewMsgSetContributorAdd(creator string, setId uint64, user string) *MsgSetContributorAdd {
	return &MsgSetContributorAdd{
		Creator: creator,
		SetId:   setId,
		User:    user,
	}
}

func (msg *MsgSetContributorAdd) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
