package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgReportMatch = "report_match"

var _ sdk.Msg = &MsgReportMatch{}

func NewMsgReportMatch(creator string, matchId uint64, cardsA []uint64, cardsB []uint64, outcome Outcome) *MsgReportMatch {
	return &MsgReportMatch{
		Creator:      creator,
		MatchId:      matchId,
		PlayedCardsA: cardsA,
		PlayedCardsB: cardsB,
		Outcome:      outcome,
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
		return sdkerrors.Wrapf(errors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
