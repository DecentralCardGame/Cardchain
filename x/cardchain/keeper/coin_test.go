package keeper_test

import (
	"testing"

	// testkeeper "github.com/DecentralCardGame/Cardchain/testutil/keeper"
	"github.com/DecentralCardGame/Cardchain/x/cardchain/keeper"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestCoins(t *testing.T) {
	// k, ctx := testkeeper.CardchainKeeper(t)
	coin1 := sdk.NewInt64Coin("ucredits", 10)
	coin2 := sdk.NewInt64Coin("ucredits", 20)
	coin3 := sdk.NewInt64Coin("ucredits", 15)

	// Mulcoin
	require.EqualValues(t, coin2, keeper.MulCoin(coin1, 2))
	require.NotEqualValues(t, coin2, keeper.MulCoin(coin1, 3))
	// QuoCoin
	require.EqualValues(t, coin1, keeper.QuoCoin(coin2, 2))
	require.NotEqualValues(t, coin1, keeper.QuoCoin(coin2, 3))
	// MulCoinFloat
	require.EqualValues(t, coin3, keeper.MulCoinFloat(coin1, 1.5))
	require.NotEqualValues(t, coin3, keeper.MulCoinFloat(coin1, 2.5))

	// // MintCoinsToAddr won't be tested as of now, because of a panic somewhere
	// addr, err := sdk.AccAddressFromBech32("cosmos15kq043zhu0wjyuw9av0auft06y3v2kxss862qf")
	// if err != nil {
	// 	panic(err)
	// }
	//
	// err = k.MintCoinsToAddr(ctx, addr, sdk.Coins{coin1})
	// require.EqualValues(t, nil, err)
}
