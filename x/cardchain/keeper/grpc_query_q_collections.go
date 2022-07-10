package keeper

import (
	"context"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) QCollections(goCtx context.Context, req *types.QueryQCollectionsRequest) (*types.QueryQCollectionsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var (
		collectionIds []uint64
	)

	ctx := sdk.UnwrapSDKContext(goCtx)

	iter := k.Collections.GetItemIterator(ctx)
	for ; iter.Valid(); iter.Next() {
		idx, collection := iter.Value()

		// Checks for status
		if !req.IgnoreStatus {
			if req.Status != collection.Status {
				continue
			}
		}

		collectionIds = append(collectionIds, idx)
	}

	return &types.QueryQCollectionsResponse{collectionIds}, nil
}
