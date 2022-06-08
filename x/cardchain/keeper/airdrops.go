package keeper

import (
	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// ClaimAirDrop claims a given airdrop
func (k Keeper) ClaimAirDrop(ctx sdk.Context, user *User, airdrop types.AirDrop) (ret bool) {
	params := k.GetParams(ctx)

	// Ensures AirDrops is set
	if user.AirDrops == nil {
		user.AirDrops = &types.AirDrops{}
	}

	if ctx.BlockHeight() > params.AirDropMaxBlockHeight {
		return
	}

  airdrops := []*bool{
    &user.AirDrops.Play, &user.AirDrops.Vote, &user.AirDrops.Create,
		&user.AirDrops.Buy, &user.AirDrops.User,
  }

  choosen := airdrops[int(airdrop)]
  if !*choosen {
		// Coin transfer happens here
		k.MintCoinsToAddr(ctx, user.Addr, sdk.Coins{params.AirDropValue})
    ret = true
    *choosen = true
  }

  return
}
