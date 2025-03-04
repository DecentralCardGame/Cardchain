package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/cardchain module sentinel errors
var (
	ErrCardDoesNotExist       = sdkerrors.Register(ModuleName, 1, "card does not exist")
	ErrVoterHasNoVotingRights = sdkerrors.Register(ModuleName, 2, "The voter doesn't have any voting rights")
	ErrVoteRightIsExpired     = sdkerrors.Register(ModuleName, 3, "The right to vote on the card has expired")
	ErrInvalidAccAddress      = sdkerrors.Register(ModuleName, 4, "Not able to convert Address to AccAddress")
	ErrUserDoesNotExist       = sdkerrors.Register(ModuleName, 5, "User does not exist")
	ErrSetNotInDesign         = sdkerrors.Register(ModuleName, 6, "Set not in design")
	ErrSetSize                = sdkerrors.Register(ModuleName, 7, "Set size is bad")
	ErrCardAlreadyInSet       = sdkerrors.Register(ModuleName, 8, "Card already in set")
	ErrNoActiveSet            = sdkerrors.Register(ModuleName, 9, "No active set")
	ErrCardNotThere           = sdkerrors.Register(ModuleName, 10, "Card not there")
	ErrContributor            = sdkerrors.Register(ModuleName, 11, "Contributor error")
	ErrNoOpenSellOffer        = sdkerrors.Register(ModuleName, 12, "No open sell-offer")
	ErrInvalidCardStatus      = sdkerrors.Register(ModuleName, 13, "Invalid card-status")
	ErrInvalidUserStatus      = sdkerrors.Register(ModuleName, 14, "Invalid user-status")
	ErrBadReveal              = sdkerrors.Register(ModuleName, 15, "Reveal does not fit commit")
	ErrCouncilStatus          = sdkerrors.Register(ModuleName, 16, "Wrong council status")
	ErrImageSizeExceeded      = sdkerrors.Register(ModuleName, 17, "Image too big! Max size is 500kb")
	ErrConversion             = sdkerrors.Register(ModuleName, 18, "Unable to convert types")
	ErrCardobject             = sdkerrors.Register(ModuleName, 19, "Faulty cardobject")
	ErrBoosterPack            = sdkerrors.Register(ModuleName, 20, "Unable to open Boosterpack")
	ErrStringLength           = sdkerrors.Register(ModuleName, 21, "String literal too long")
	ErrUserAlreadyExists      = sdkerrors.Register(ModuleName, 22, "User already exists")
	ErrWaitingForPlayers      = sdkerrors.Register(ModuleName, 23, "Waiting for players")
	ErrUninitializedType      = sdkerrors.Register(ModuleName, 24, "Type not yet initialized")
	ErrFinalizeSet            = sdkerrors.Register(ModuleName, 25, "Set can't be finalized")
	InvalidAlias              = sdkerrors.Register(ModuleName, 26, "Invalid alias")
	ErrInvalidData            = sdkerrors.Register(ModuleName, 27, "Invalid data in transaction")
)
