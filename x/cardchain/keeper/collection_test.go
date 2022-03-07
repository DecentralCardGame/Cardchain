package keeper_test

import (
	"testing"

	testkeeper "github.com/DecentralCardGame/Cardchain/testutil/keeper"
	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
  // sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestCollection(t *testing.T) {
	k, ctx := testkeeper.CardchainKeeper(t)
  setUpCard(ctx, k)
	collection := types.Collection{
    Name: "Coll collection",
    Cards: []uint64{0},
    Artist: "cosmos15kq043zhu0wjyuw9av0auft06y3v2kxss862qf",
    Contributors: []string{"cosmos15kq043zhu0wjyuw9av0auft06y3v2kxss862qf", "cosmos15kq043zhu0wjyuw9av0auft06y3v2kxss862qf", "cosmos15kq043zhu0wjyuw9av0auft06y3v2kxss862qf"},
    StoryWriter: "cosmos15kq043zhu0wjyuw9av0auft06y3v2kxss862qf",
    Story: "Bla bli blub",
    Status: types.CStatus_active,
    TimeStamp: 100,
  }

  contribs := []string{}
  for i := 0; i < 10; i++ {
    contribs = append(contribs, "cosmos15kq043zhu0wjyuw9av0auft06y3v2kxss862qf")
  }

	k.SetCollection(ctx, 0, collection)

	require.EqualValues(t, collection, k.GetCollection(ctx, 0))
  require.EqualValues(t, []*types.Collection{&collection}, k.GetAllCollections(ctx))
  require.EqualValues(t, 1, k.GetCollectionsNumber(ctx))
  require.EqualValues(t, []uint64{0}, k.GetActiveCollections(ctx))
  require.EqualValues(t, contribs, k.GetAllCollectionContributors(ctx, collection))
}
