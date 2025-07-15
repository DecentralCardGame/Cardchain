package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCouncilResponseReveal{}

func NewMsgCouncilResponseReveal(creator string, councilId uint64, reponse Response, secret string) *MsgCouncilResponseReveal {
	return &MsgCouncilResponseReveal{
		Creator:   creator,
		CouncilId: councilId,
		Response:  reponse,
		Secret:    secret,
	}
}

func (msg *MsgCouncilResponseReveal) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
