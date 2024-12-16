package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCouncilRestart{}

func NewMsgCouncilRestart(creator string, councilId uint64) *MsgCouncilRestart {
	return &MsgCouncilRestart{
		Creator:   creator,
		CouncilId: councilId,
	}
}

func (msg *MsgCouncilRestart) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
