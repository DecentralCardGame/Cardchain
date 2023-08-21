package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgConfirmMatch = "confirm_match"

var _ sdk.Msg = &MsgConfirmMatch{}

func NewMsgConfirmMatch(creator string, matchId uint64, outcome Outcome) *MsgConfirmMatch {
	return &MsgConfirmMatch{
		Creator: creator,
		MatchId: matchId,
		Outcome: outcome,
	}
}

func (msg *MsgConfirmMatch) Route() string {
	return RouterKey
}

func (msg *MsgConfirmMatch) Type() string {
	return TypeMsgConfirmMatch
}

func (msg *MsgConfirmMatch) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgConfirmMatch) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgConfirmMatch) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(errors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
