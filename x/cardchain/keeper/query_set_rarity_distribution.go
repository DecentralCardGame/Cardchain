package keeper

import (
	"context"

	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) SetRarityDistribution(goCtx context.Context, req *types.QuerySetRarityDistributionRequest) (*types.QuerySetRarityDistributionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	setSize := k.GetParams(ctx).SetSize
	set := k.SetK.Get(ctx, req.SetId)
	dist, err := k.GetRarityDistribution(ctx, *set, uint32(setSize))

	return &types.QuerySetRarityDistributionResponse{Current: dist[0][:], Wanted: dist[1][:]}, err
}
