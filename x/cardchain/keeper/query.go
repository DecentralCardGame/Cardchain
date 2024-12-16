package keeper

import (
	"cardchain/x/cardchain/types"
)

var _ types.QueryServer = Keeper{}
