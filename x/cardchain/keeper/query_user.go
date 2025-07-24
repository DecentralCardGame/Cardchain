package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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

	if user.Alias == "" {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "user doesnt exist")
	}

	return &types.QueryUserResponse{User: user.User}, nil
}
