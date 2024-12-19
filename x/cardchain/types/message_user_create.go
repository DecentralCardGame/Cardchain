package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgUserCreate{}

func NewMsgUserCreate(creator string, newUser string, alias string) *MsgUserCreate {
	return &MsgUserCreate{
		Creator: creator,
		NewUser: newUser,
		Alias:   alias,
	}
}

func (msg *MsgUserCreate) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if err = checkAliasLimit(msg.Alias); err != nil {
		return err
	}

	return nil
}
