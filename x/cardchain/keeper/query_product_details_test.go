package keeper_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/DecentralCardGame/cardchain/testutil/keeper"
	"github.com/DecentralCardGame/cardchain/testutil/nullify"
	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
)

func TestProductDetailsQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.CardchainKeeper(t)
	msgs := createNProductDetails(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetProductDetailsRequest
		response *types.QueryGetProductDetailsResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetProductDetailsRequest{Id: msgs[0].Id},
			response: &types.QueryGetProductDetailsResponse{ProductDetails: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetProductDetailsRequest{Id: msgs[1].Id},
			response: &types.QueryGetProductDetailsResponse{ProductDetails: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetProductDetailsRequest{Id: uint64(len(msgs))},
			err:     sdkerrors.ErrKeyNotFound,
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.ProductDetails(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestProductDetailsQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.CardchainKeeper(t)
	msgs := createNProductDetails(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllProductDetailsRequest {
		return &types.QueryAllProductDetailsRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.ProductDetailsAll(ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.ProductDetails), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.ProductDetails),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.ProductDetailsAll(ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.ProductDetails), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.ProductDetails),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.ProductDetailsAll(ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.ProductDetails),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.ProductDetailsAll(ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
