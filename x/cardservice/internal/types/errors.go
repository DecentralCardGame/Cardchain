package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrCardDoesNotExist 				= sdkerrors.Register(ModuleName, 1, "card does not exist")
	ErrVoterHasNoVotingRights 	= sdkerrors.Register(ModuleName, 2, "The voter doesn't have any voting rights")
	ErrVoteRightIsExpired 			= sdkerrors.Register(ModuleName, 3, "The right to vote on the card has expired")
)
