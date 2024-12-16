package keeper

import (
	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
)

var _ types.QueryServer = Keeper{}
