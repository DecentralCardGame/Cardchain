package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCommitCouncilResponse = "commit_council_response"

var _ sdk.Msg = &MsgCommitCouncilResponse{}

func NewMsgCommitCouncilResponse(creator string, response string, councilId uint64) *MsgCommitCouncilResponse {
	return &MsgCommitCouncilResponse{
		Creator:  creator,
		Response: response,
		CouncilId: councilId,
	}
}

func (msg *MsgCommitCouncilResponse) Route() string {
	return RouterKey
}

func (msg *MsgCommitCouncilResponse) Type() string {
	return TypeMsgCommitCouncilResponse
}

func (msg *MsgCommitCouncilResponse) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCommitCouncilResponse) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCommitCouncilResponse) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
