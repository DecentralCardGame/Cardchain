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

	setSize := k.GetParams(ctx).SetSize
	set := k.Sets.Get(ctx, req.SetId)
	dist, err := k.GetRarityDistribution(ctx, *set, setSize)

	return &types.QueryRarityDistributionResponse{Current: dist[0][:], Wanted: dist[1][:]}, err
}
