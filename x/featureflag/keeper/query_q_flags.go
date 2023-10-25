package keeper

import (
	"context"

	"github.com/DecentralCardGame/Cardchain/x/featureflag/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) QFlags(goCtx context.Context, req *types.QueryQFlagsRequest) (*types.QueryQFlagsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	flags, _ := k.GetAllFlags(ctx)

	return &types.QueryQFlagsResponse{Flags: flags}, nil
}
