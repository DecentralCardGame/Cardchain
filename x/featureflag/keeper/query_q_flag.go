package keeper

import (
	"context"

	"github.com/DecentralCardGame/Cardchain/x/featureflag/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) QFlag(goCtx context.Context, req *types.QueryQFlagRequest) (*types.QueryQFlagResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	flag := k.GetFlag(ctx, GetKey(req.Module, req.Name))

	return &types.QueryQFlagResponse{Flag: &flag}, nil
}
