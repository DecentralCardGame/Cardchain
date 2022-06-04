package keeper

import (
	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) ClaimAirDrop(ctx sdk.Context, user *User, airdrop types.AirDrop) (ret bool) {
	// Ensures AirDrops is set
	if user.AirDrops == nil {
		user.AirDrops = &types.AirDrops{}
	}

  airdrops := []*bool{
    &user.AirDrops.Play, &user.AirDrops.Vote, &user.AirDrops.Create, &user.AirDrops.Buy,
  }

  choosen := airdrops[int(airdrop)]
  if !*choosen {
		// Coin transfer happens here
    ret = true
    *choosen = true
  }

  return
}
