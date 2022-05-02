package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRevealCouncilResponse = "reveal_council_response"

var _ sdk.Msg = &MsgRevealCouncilResponse{}

func NewMsgRevealCouncilResponse(creator string, response Response, secret string, councilId uint64) *MsgRevealCouncilResponse {
	return &MsgRevealCouncilResponse{
		Creator:   creator,
		Response:  response,
		Secret:    secret,
		CouncilId: councilId,
	}
}

func (msg *MsgRevealCouncilResponse) Route() string {
	return RouterKey
}

func (msg *MsgRevealCouncilResponse) Type() string {
	return TypeMsgRevealCouncilResponse
}

func (msg *MsgRevealCouncilResponse) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRevealCouncilResponse) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRevealCouncilResponse) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
