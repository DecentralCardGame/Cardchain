package keeper_test

import (
	"testing"

	testkeeper "github.com/DecentralCardGame/Cardchain/testutil/keeper"
	"github.com/DecentralCardGame/Cardchain/x/cardchain/keeper"
	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	// sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestIncentives(t *testing.T) {
	k, ctx := testkeeper.CardchainKeeper(t)
	runningAverage := types.RunningAverage{[]int64{1, 5, 7}}
	var clearRunningAverage types.RunningAverage
	params := types.DefaultParams()

	k.SetParams(ctx, params)
	k.RunningAverages.Set(ctx, keeper.Games24ValueKey, &runningAverage)

	require.EqualValues(t, runningAverage, *k.RunningAverages.Get(ctx, keeper.Games24ValueKey))
	require.EqualValues(t, []*types.RunningAverage{&runningAverage, &clearRunningAverage}, k.RunningAverages.GetAll(ctx))
	require.EqualValues(t, 13, k.GetGames(ctx))
	require.EqualValues(t, 1, k.GetVotes(ctx))
	require.EqualValues(t, float32(0.98484848484848484848484848484848), k.GetWinnerIncentives(ctx))
	require.EqualValues(t, float32(0.01515151515151515151515151515152), k.GetBalancerIncentives(ctx))
}
