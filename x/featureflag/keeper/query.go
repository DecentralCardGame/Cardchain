package keeper

import (
	"github.com/DecentralCardGame/Cardchain/x/featureflag/types"
)

var _ types.QueryServer = Keeper{}
