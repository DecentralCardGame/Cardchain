package keeper

import (
	"context"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) QCollection(goCtx context.Context, req *types.QueryQCollectionRequest) (*types.QueryQCollectionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	collection := k.GetCollection(ctx, req.CollectionId)

	return &types.QueryQCollectionResponse{
		collection.Name,
		collection.Cards,
		collection.Contributors,
		collection.Story,
		collection.Artwork,
		collection.Status,
	}, nil
}
