package cardservice

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	DefaultCodespace       sdk.CodespaceType = "cardservice"
	VoterHasNoVotingRights sdk.CodeType      = 101
	VoteRightIsExpired     sdk.CodeType      = 102
)

// ErrNoVotingRights indicates that the voter does not have any voting rights
func ErrNoVotingRights() sdk.Error {
	return sdk.NewError(DefaultCodespace, VoterHasNoVotingRights, "The voter doesn't have any voting rights")
}

// ErrVoteRightHasExpired indicates that the right to vote on a card has expired
func ErrVoteRightHasExpired() sdk.Error {
	return sdk.NewError(DefaultCodespace, VoteRightIsExpired, "The right to vote on the card has expired")
}
