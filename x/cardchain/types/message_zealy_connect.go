package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgZealyConnect{}

func NewMsgZealyConnect(creator string, zealyId string) *MsgZealyConnect {
	return &MsgZealyConnect{
		Creator: creator,
		ZealyId: zealyId,
	}
}

func (msg *MsgZealyConnect) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
