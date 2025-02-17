package keeper

import (
	"context"

	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) User(goCtx context.Context, req *types.QueryUserRequest) (*types.QueryUserResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	user, err := k.GetUserFromString(ctx, req.Address)
	if err != nil {
		return nil, err
	}
	return &types.QueryUserResponse{User: user.User}, nil
}
