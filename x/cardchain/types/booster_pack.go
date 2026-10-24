package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewBoosterPack(ctx sdk.Context, setId uint64, raritiesPerPack []uint64, dropRatiosPerPack []uint64) BoosterPack {
	return BoosterPack{
		SetId:             setId,
		RaritiesPerPack:   raritiesPerPack,
		TimeStamp:         ctx.BlockHeight(),
		DropRatiosPerPack: dropRatiosPerPack,
	}
}
