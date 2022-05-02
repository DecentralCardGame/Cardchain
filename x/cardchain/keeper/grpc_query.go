package keeper

import (
	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
)

var _ types.QueryServer = Keeper{}
