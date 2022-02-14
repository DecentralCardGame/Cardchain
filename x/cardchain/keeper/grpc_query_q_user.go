package keeper

import (
	"context"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) QUser(goCtx context.Context, req *types.QueryQUserRequest) (*types.QueryQUserResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	user, err := k.GetUserFromString(ctx, req.Address)
	if err != nil {
		return nil, err
	}

	return &types.QueryQUserResponse{
		user.Alias,
		user.OwnedCardSchemes,
		user.OwnedPrototypes,
		user.Cards,
		user.VoteRights,
		user.CouncilStatus,
		user.ReportMatches,
	}, nil
}
