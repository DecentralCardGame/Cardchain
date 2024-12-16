package keeper_test

import (
	"testing"

	testkeeper "github.com/DecentralCardGame/cardchain/testutil/keeper"
	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.CardchainKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
