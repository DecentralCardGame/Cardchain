package keeper

import (
	"github.com/DecentralCardGame/cardchain/x/featureflag/types"
)

var _ types.QueryServer = Keeper{}
