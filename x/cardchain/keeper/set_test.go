package keeper_test

import (
	"testing"

	testkeeper "github.com/DecentralCardGame/Cardchain/testutil/keeper"
	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	// sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestSet(t *testing.T) {
	k, ctx := testkeeper.CardchainKeeper(t)
	setUpCard(ctx, k)
	set := types.Set{
		Name:         "Coll set",
		Cards:        []uint64{0},
		Artist:       "cosmos15kq043zhu0wjyuw9av0auft06y3v2kxss862qf",
		Contributors: []string{"cosmos15kq043zhu0wjyuw9av0auft06y3v2kxss862qf", "cosmos15kq043zhu0wjyuw9av0auft06y3v2kxss862qf", "cosmos15kq043zhu0wjyuw9av0auft06y3v2kxss862qf"},
		StoryWriter:  "cosmos15kq043zhu0wjyuw9av0auft06y3v2kxss862qf",
		Story:        "Bla bli blub",
		Status:       types.CStatus_active,
		TimeStamp:    100,
	}

	var contribs []string
	for i := 0; i < 10; i++ {
		contribs = append(contribs, "cosmos15kq043zhu0wjyuw9av0auft06y3v2kxss862qf")
	}

	k.Sets.Set(ctx, 0, &set)

	require.EqualValues(t, set, *k.Sets.Get(ctx, 0))
	require.EqualValues(t, []*types.Set{&set}, k.Sets.GetAll(ctx))
	require.EqualValues(t, 1, k.Sets.GetNumber(ctx))
	require.EqualValues(t, []uint64{0}, k.GetActiveSets(ctx))
	require.EqualValues(t, contribs, k.GetAllSetContributors(ctx, set))
}
