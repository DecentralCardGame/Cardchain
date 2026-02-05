package keeper

import (
	"context"

	"github.com/DecentralCardGame/cardchain/x/featureflag/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Flags(goCtx context.Context, req *types.QueryFlagsRequest) (*types.QueryFlagsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	flags, _ := k.GetAllFlags(ctx)

	return &types.QueryFlagsResponse{Flags: flags}, nil
}
