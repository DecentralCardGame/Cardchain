package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/DecentralCardGame/cardchain/testutil/keeper"
	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.CardchainKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
