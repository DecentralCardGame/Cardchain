package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/featureflag module sentinel errors
var (
	ErrFlagUnregisterd = sdkerrors.Register(ModuleName, 1, "Flag is not registered")
)
