package keeper_test

import (
	"testing"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/keeper"
  sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestCoins(t *testing.T) {
  // Mulcoin
  coin1 := sdk.NewInt64Coin("ucredits", 10)
  coin2 := sdk.NewInt64Coin("ucredits", 20)
  require.EqualValues(t, coin2, keeper.MulCoin(coin1, 2))
}
