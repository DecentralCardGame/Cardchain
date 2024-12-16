package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetArtworkAdd{}

func NewMsgSetArtworkAdd(creator string, setId uint64, image []byte) *MsgSetArtworkAdd {
	return &MsgSetArtworkAdd{
		Creator: creator,
		SetId:   setId,
		Image:   image,
	}
}

func (msg *MsgSetArtworkAdd) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
