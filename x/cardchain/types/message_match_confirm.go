package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgMatchConfirm{}

func NewMsgMatchConfirm(creator string, matchId uint64, outcome Outcome, votedCards []*SingleVote) *MsgMatchConfirm {
	return &MsgMatchConfirm{
		Creator:    creator,
		MatchId:    matchId,
		Outcome:    outcome,
		VotedCards: votedCards,
	}
}

func (msg *MsgMatchConfirm) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
