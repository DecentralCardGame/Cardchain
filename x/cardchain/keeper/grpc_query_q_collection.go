package keeper

import (
	"context"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) QCollection(goCtx context.Context, req *types.QueryQCollectionRequest) (*types.OutpCollection, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	collection := k.Collections.Get(ctx, req.CollectionId)

	image := k.Images.Get(ctx, collection.ArtworkId)

	return &types.OutpCollection{
		Name:         collection.Name,
		Cards:        collection.Cards,
		Artist:       collection.Artist,
		StoryWriter:  collection.StoryWriter,
		Contributors: collection.Contributors,
		Story:        collection.Story,
		Artwork:      string(image.Image),
		Status:       collection.Status,
		TimeStamp:    collection.TimeStamp,
	}, nil
}
