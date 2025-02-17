package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCouncilResponseCommit{}

func NewMsgCouncilResponseCommit(creator string, councilId uint64, reponse string, suggestion string) *MsgCouncilResponseCommit {
	return &MsgCouncilResponseCommit{
		Creator:    creator,
		CouncilId:  councilId,
		Response:   reponse,
		Suggestion: suggestion,
	}
}

func (msg *MsgCouncilResponseCommit) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
