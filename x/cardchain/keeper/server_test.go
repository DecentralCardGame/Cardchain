package keeper_test

import (
	"testing"

	testkeeper "github.com/DecentralCardGame/Cardchain/testutil/keeper"
	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	"github.com/stretchr/testify/require"
)

func TestServers(t *testing.T) {
	k, ctx := testkeeper.CardchainKeeper(t)

	require.EqualValues(t, 0, k.GetServerId(ctx, "cosmos15kq043zhu0wjyuw9av0auft06y3v2kxss862qf"))
  require.EqualValues(t, types.Server{Reporter: "cosmos15kq043zhu0wjyuw9av0auft06y3v2kxss862qf"}, *k.Servers.Get(ctx, 0))

  k.ReportServerMatch(ctx, "cosmos15kq043zhu0wjyuw9av0auft06y3v2kxss862qf", 1, true)
  require.EqualValues(t, types.Server{Reporter: "cosmos15kq043zhu0wjyuw9av0auft06y3v2kxss862qf", ValidReports: 1}, *k.Servers.Get(ctx, 0))

  k.ReportServerMatch(ctx, "cosmos15kq043zhu0wjyuw9av0auft06y3v2kxss862qf", 2, false)
  require.EqualValues(t, types.Server{Reporter: "cosmos15kq043zhu0wjyuw9av0auft06y3v2kxss862qf", ValidReports: 1, InvalidReports: 2}, *k.Servers.Get(ctx, 0))
}
