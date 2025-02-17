package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgEncounterCreate{}

func NewMsgEncounterCreate(creator string, name string, drawlist []uint64, parameters map[string]string, image []byte) *MsgEncounterCreate {
	return &MsgEncounterCreate{
		Creator:    creator,
		Name:       name,
		Drawlist:   drawlist,
		Parameters: parameters,
		Image:      image,
	}
}

func (msg *MsgEncounterCreate) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
