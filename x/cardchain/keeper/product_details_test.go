package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/DecentralCardGame/cardchain/testutil/keeper"
	"github.com/DecentralCardGame/cardchain/testutil/nullify"
	"github.com/DecentralCardGame/cardchain/x/cardchain/keeper"
	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	"github.com/stretchr/testify/require"
)

func createNProductDetails(keeper keeper.Keeper, ctx context.Context, n int) []types.ProductDetails {
	items := make([]types.ProductDetails, n)
	for i := range items {
		items[i].Id = keeper.AppendProductDetails(ctx, items[i])
	}
	return items
}

func TestProductDetailsGet(t *testing.T) {
	keeper, ctx := keepertest.CardchainKeeper(t)
	items := createNProductDetails(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetProductDetails(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestProductDetailsRemove(t *testing.T) {
	keeper, ctx := keepertest.CardchainKeeper(t)
	items := createNProductDetails(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveProductDetails(ctx, item.Id)
		_, found := keeper.GetProductDetails(ctx, item.Id)
		require.False(t, found)
	}
}

func TestProductDetailsGetAll(t *testing.T) {
	keeper, ctx := keepertest.CardchainKeeper(t)
	items := createNProductDetails(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllProductDetails(ctx)),
	)
}

func TestProductDetailsCount(t *testing.T) {
	keeper, ctx := keepertest.CardchainKeeper(t)
	items := createNProductDetails(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetProductDetailsCount(ctx))
}
