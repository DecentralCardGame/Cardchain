package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/cardchain module sentinel errors
var (
	ErrSample                  = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrCardDoesNotExist        = sdkerrors.Register(ModuleName, 1, "card does not exist")
	ErrVoterHasNoVotingRights  = sdkerrors.Register(ModuleName, 2, "The voter doesn't have any voting rights")
	ErrVoteRightIsExpired      = sdkerrors.Register(ModuleName, 3, "The right to vote on the card has expired")
	ErrInvalidAccAddress       = sdkerrors.Register(ModuleName, 4, "Not able to convert Address to AccAddress")
	ErrUserDoesNotExist        = sdkerrors.Register(ModuleName, 5, "User does not exist")
	ErrCollectionNotInDesign   = sdkerrors.Register(ModuleName, 6, "Collection not in design")
	ErrCollectionSize          = sdkerrors.Register(ModuleName, 7, "Collection size is bad")
	ErrCardAlreadyInCollection = sdkerrors.Register(ModuleName, 8, "Card already in collection")
	ErrNoActiveCollection      = sdkerrors.Register(ModuleName, 9, "No active collection")
	ErrCardNotThere            = sdkerrors.Register(ModuleName, 10, "Card not there")
	ErrContributor             = sdkerrors.Register(ModuleName, 11, "Contributor error")
	ErrNoOpenSellOffer         = sdkerrors.Register(ModuleName, 12, "No open sell-offer")
	ErrInvalidCardStatus       = sdkerrors.Register(ModuleName, 13, "Invalid card-status")
	ErrInvalidUserStatus       = sdkerrors.Register(ModuleName, 14, "Invalid user-status")
)
