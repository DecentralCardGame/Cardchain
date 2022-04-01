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
	match := types.Match{
		0,
		"cosmos15kq043zhu0wjyuw9av0auft06y3v2kxss862qf",
		types.NewMatchPlayer("cosmos15kq043zhu0wjyuw9av0auft06y3v2kxss862qf", []uint64{1, 2, 3}),
		types.NewMatchPlayer("cosmos15kq043zhu0wjyuw9av0auft06y3v2kxss862qf", []uint64{5, 6, 7}),
		types.Outcome_AWon,
		false,
	}
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

	addrs, err := k.GetMatchAddresses(ctx, match)
	require.EqualValues(t, nil, err)
	require.EqualValues(t, addrs[0], addrs[1])
	require.EqualValues(t, 2, len(addrs))

	match.PlayerA.Addr = "abc"
	addrs, err = k.GetMatchAddresses(ctx, match)
	require.NotEqualValues(t, nil, err)

	outcome, err := k.GetOutcome(ctx, match)
	require.EqualValues(t, nil, err)
	require.EqualValues(t, types.Outcome_Aborted, outcome)

	match.PlayerA.Confirmed = true
	match.PlayerA.Outcome = types.Outcome_BWon
	outcome, err = k.GetOutcome(ctx, match)
	require.NotEqualValues(t, nil, err)

	match.PlayerB.Confirmed = true
	match.PlayerB.Outcome = types.Outcome_AWon
	outcome, err = k.GetOutcome(ctx, match)
	require.EqualValues(t, nil, err)
	require.EqualValues(t, types.Outcome_AWon, outcome)
}
