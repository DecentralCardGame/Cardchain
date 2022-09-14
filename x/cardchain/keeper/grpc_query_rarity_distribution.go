package keeper

import (
	"context"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) RarityDistribution(goCtx context.Context, req *types.QueryRarityDistributionRequest) (*types.QueryRarityDistributionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	collectionSize := k.GetParams(ctx).CollectionSize
	collection := k.Collections.Get(ctx, req.CollectionId)
	dist, err := k.GetRarityDistribution(ctx, *collection, collectionSize)

	return &types.QueryRarityDistributionResponse{dist[0][:], dist[1][:]}, err
}
