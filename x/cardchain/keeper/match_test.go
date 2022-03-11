package keeper_test

import (
	"testing"

	testkeeper "github.com/DecentralCardGame/Cardchain/testutil/keeper"
	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestMatches(t *testing.T) {
	k, ctx := testkeeper.CardchainKeeper(t)
	setUpCard(ctx, k)
	SetUpPools(ctx, *k)
	match := types.Match{0, "cosmos15kq043zhu0wjyuw9av0auft06y3v2kxss862qf", "cosmos15kq043zhu0wjyuw9av0auft06y3v2kxss862qf", "cosmos15kq043zhu0wjyuw9av0auft06y3v2kxss862qf", types.Outcome_AWon}
	params := types.DefaultParams()

	k.SetParams(ctx, params)
	k.SetMatch(ctx, 0, match)
	k.SetMatch(ctx, 1, match)

	require.EqualValues(t, match, k.GetMatch(ctx, 0))
	require.EqualValues(t, []*types.Match{&match, &match}, k.GetAllMatches(ctx))
	require.EqualValues(t, 2, k.GetMatchesNumber(ctx))
	require.EqualValues(t, sdk.NewInt64Coin("ucredits", 2), k.GetMatchReward(ctx), k.GetMatchReward(ctx).Amount.Int64())

	amountA, amountB := k.CalculateMatchReward(ctx, types.Outcome_AWon)
	require.EqualValues(t, sdk.NewInt64Coin("ucredits", 2), amountA)
	require.EqualValues(t, sdk.NewInt64Coin("ucredits", 0), amountB)

	SetUpPools(ctx, *k)
	amountA, amountB = k.CalculateMatchReward(ctx, types.Outcome_BWon)
	require.EqualValues(t, sdk.NewInt64Coin("ucredits", 0), amountA)
	require.EqualValues(t, sdk.NewInt64Coin("ucredits", 2), amountB)

	SetUpPools(ctx, *k)
	amountA, amountB = k.CalculateMatchReward(ctx, types.Outcome_Draw)
	require.EqualValues(t, sdk.NewInt64Coin("ucredits", 1), amountA)
	require.EqualValues(t, sdk.NewInt64Coin("ucredits", 1), amountB)

	SetUpPools(ctx, *k)
	amountA, amountB = k.CalculateMatchReward(ctx, types.Outcome_Aborted)
	require.EqualValues(t, sdk.NewInt64Coin("ucredits", 0), amountA)
	require.EqualValues(t, sdk.NewInt64Coin("ucredits", 0), amountB)
}
