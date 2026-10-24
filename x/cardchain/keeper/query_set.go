package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Set(goCtx context.Context, req *types.QuerySetRequest) (*types.QuerySetResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	set := k.SetK.Get(ctx, req.SetId)
	if set == nil {
		return nil, errorsmod.Wrap(errors.ErrInvalidRequest, "setId does not represent a set")
	}

	image := k.Images.Get(ctx, set.ArtworkId)

	return &types.QuerySetResponse{Set: &types.SetWithArtwork{
		Set:     *set,
		Artwork: image.Image,
	}}, nil
}
