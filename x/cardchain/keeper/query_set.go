package keeper

import (
	"context"

	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Set(goCtx context.Context, req *types.QuerySetRequest) (*types.QuerySetResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	set := k.sets.Get(ctx, req.SetId)
	image := k.images.Get(ctx, set.ArtworkId)

	return &types.QuerySetResponse{Set: &types.SetWithArtwork{
		Set:     set,
		Artwork: image.Image,
	}}, nil
}
