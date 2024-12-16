package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgMatchReport{}

func NewMsgMatchReport(creator string, matchId uint64, playedCardsA []uint64, playedCardsB []uint64, outcome int32) *MsgMatchReport {
	return &MsgMatchReport{
		Creator:      creator,
		MatchId:      matchId,
		PlayedCardsA: playedCardsA,
		PlayedCardsB: playedCardsB,
		Outcome:      outcome,
	}
}

func (msg *MsgMatchReport) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
