package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ProductDetailsAll(ctx context.Context, req *types.QueryAllProductDetailsRequest) (*types.QueryAllProductDetailsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var productDetailss []types.ProductDetails

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	productDetailsStore := prefix.NewStore(store, types.KeyPrefix(types.ProductDetailsKey))

	pageRes, err := query.Paginate(productDetailsStore, req.Pagination, func(key []byte, value []byte) error {
		var productDetails types.ProductDetails
		if err := k.cdc.Unmarshal(value, &productDetails); err != nil {
			return err
		}

		productDetailss = append(productDetailss, productDetails)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllProductDetailsResponse{ProductDetails: productDetailss, Pagination: pageRes}, nil
}

func (k Keeper) ProductDetails(ctx context.Context, req *types.QueryGetProductDetailsRequest) (*types.QueryGetProductDetailsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	productDetails, found := k.GetProductDetails(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetProductDetailsResponse{ProductDetails: productDetails}, nil
}
