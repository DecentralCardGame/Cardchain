package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/cardchain module sentinel errors
var (
	ErrInvalidSigner     = sdkerrors.Register(ModuleName, 1100, "expected gov account as only signer for proposal message")
	ErrCardDoesNotExist  = sdkerrors.Register(ModuleName, 1, "card does not exist")
	ErrCardAlreadyInSet  = sdkerrors.Register(ModuleName, 8, "Card already in set")
	ErrInvalidCardStatus = sdkerrors.Register(ModuleName, 13, "Invalid card-status")
	ErrUserDoesNotExist  = sdkerrors.Register(ModuleName, 5, "User does not exist")
	ErrCardobject        = sdkerrors.Register(ModuleName, 19, "Faulty cardobject")
	ErrImageSizeExceeded = sdkerrors.Register(ModuleName, 17, "Image too big! Max size is 500kb")
	ErrInvalidData       = sdkerrors.Register(ModuleName, 27, "Invalid data in transaction")
)
