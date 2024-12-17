package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgProfileAliasSet{}

func NewMsgProfileAliasSet(creator string, alias string) *MsgProfileAliasSet {
	return &MsgProfileAliasSet{
		Creator: creator,
		Alias:   alias,
	}
}

func (msg *MsgProfileAliasSet) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
