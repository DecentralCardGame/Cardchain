package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgReportMatch = "report_match"

var _ sdk.Msg = &MsgReportMatch{}

func NewMsgReportMatch(creator string, playerA string, playerB string, cardsA []uint64, cardsB []uint64, outcome Outcome) *MsgReportMatch {
	return &MsgReportMatch{
		Creator: creator,
		PlayerA: playerA,
		PlayerB: playerB,
		CardsA: cardsA,
		CardsB: cardsB,
		Outcome: outcome,
	}
}

func (msg *MsgReportMatch) Route() string {
	return RouterKey
}

func (msg *MsgReportMatch) Type() string {
	return TypeMsgReportMatch
}

func (msg *MsgReportMatch) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgReportMatch) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgReportMatch) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
