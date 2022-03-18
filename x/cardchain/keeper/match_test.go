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
	match := types.Match{0, "cosmos15kq043zhu0wjyuw9av0auft06y3v2kxss862qf", "cosmos15kq043zhu0wjyuw9av0auft06y3v2kxss862qf", "cosmos15kq043zhu0wjyuw9av0auft06y3v2kxss862qf", []uint64{1, 2, 3}, []uint64{5,6,7}, types.Outcome_AWon}
	params := types.DefaultParams()

	k.SetParams(ctx, params)
	k.Matches.Set(ctx, 0, &match)
	k.Matches.Set(ctx, 1, &match)

	require.EqualValues(t, match, *k.Matches.Get(ctx, 0))
	require.EqualValues(t, []*types.Match{&match, &match}, k.Matches.GetAll(ctx))
	require.EqualValues(t, 2, k.Matches.GetNumber(ctx))
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
