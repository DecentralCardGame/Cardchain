package keeper

import (
	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	// sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) ClaimAirDrop(ctx sdk.Context, user *User, airdrop types.AirDrop) (ret bool) {
  airdrops := []*bool{
    &user.AirDrops.Play, &user.AirDrops.Vote, &user.AirDrops.Create, &user.AirDrops.Buy,
  }
  choosen := airdrops[int(airdrop)]
  if !*choosen {
    ret = true
    *choosen = true
  }
  // _ = airdrops
  return
}
