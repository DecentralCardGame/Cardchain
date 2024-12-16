package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgProfileBioSet{}

func NewMsgProfileBioSet(creator string, bio string) *MsgProfileBioSet {
	return &MsgProfileBioSet{
		Creator: creator,
		Bio:     bio,
	}
}

func (msg *MsgProfileBioSet) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
