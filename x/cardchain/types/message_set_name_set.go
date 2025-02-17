package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetNameSet{}

func NewMsgSetNameSet(creator string, setId uint64, name string) *MsgSetNameSet {
	return &MsgSetNameSet{
		Creator: creator,
		SetId:   setId,
		Name:    name,
	}
}

func (msg *MsgSetNameSet) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
