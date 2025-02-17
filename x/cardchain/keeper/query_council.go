package keeper

import (
	"context"

	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Council(goCtx context.Context, req *types.QueryCouncilRequest) (*types.QueryCouncilResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	council := k.councils.Get(ctx, req.CouncilId)

	return &types.QueryCouncilResponse{Council: council}, nil
}
