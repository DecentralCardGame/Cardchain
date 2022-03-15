package keeper_test

import (
	"testing"

	testkeeper "github.com/DecentralCardGame/Cardchain/testutil/keeper"
	// "github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	"github.com/DecentralCardGame/Cardchain/x/cardchain/keeper"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func SetUpPools(ctx sdk.Context, k keeper.Keeper) (sdk.Coin, sdk.Coin, sdk.Coin) {
	pool := sdk.NewInt64Coin("ucredits", 10)
	pool1 := sdk.NewInt64Coin("ucredits", 2000000)
	pool2 := sdk.NewInt64Coin("ucredits", 12)
	k.SetPool(ctx, keeper.PublicPoolKey, pool)
	k.SetPool(ctx, keeper.WinnersPoolKey, pool1)
	k.SetPool(ctx, keeper.BalancersPoolKey, pool2)

	return pool, pool1, pool2
}

func TestPools(t *testing.T) {
	k, ctx := testkeeper.CardchainKeeper(t)
	pool, pool1, pool2 := SetUpPools(ctx, *k)

	require.EqualValues(t, pool, k.GetPool(ctx, keeper.PublicPoolKey))
	require.EqualValues(t, []sdk.Coin{pool, pool1, pool2}, k.GetAllPools(ctx))

	k.AddPoolCredits(ctx, keeper.PublicPoolKey, sdk.NewInt64Coin("ucredits", 2))
	require.EqualValues(t, sdk.NewInt64Coin("ucredits", 12), k.GetPool(ctx, keeper.PublicPoolKey))

	k.SubPoolCredits(ctx, keeper.PublicPoolKey, sdk.NewInt64Coin("ucredits", 3))
	require.EqualValues(t, sdk.NewInt64Coin("ucredits", 9), k.GetPool(ctx, keeper.PublicPoolKey))
}
