package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetFinalize{}

func NewMsgSetFinalize(creator string, setId uint64) *MsgSetFinalize {
	return &MsgSetFinalize{
		Creator: creator,
		SetId:   setId,
	}
}

func (msg *MsgSetFinalize) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
