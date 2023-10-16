package keeper

import (
	"context"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) QSet(goCtx context.Context, req *types.QueryQSetRequest) (*types.OutpSet, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	set := k.Sets.Get(ctx, req.SetId)

	image := k.Images.Get(ctx, set.ArtworkId)

	return &types.OutpSet{
		Name:         set.Name,
		Cards:        set.Cards,
		Artist:       set.Artist,
		StoryWriter:  set.StoryWriter,
		Contributors: set.Contributors,
		Story:        set.Story,
		Artwork:      string(image.Image),
		Status:       set.Status,
		TimeStamp:    set.TimeStamp,
	}, nil
}
