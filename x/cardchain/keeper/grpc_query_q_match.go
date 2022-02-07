package keeper

import (
	"context"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) QMatch(goCtx context.Context, req *types.QueryQMatchRequest) (*types.QueryQMatchResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	match := k.GetMatch(ctx, req.MatchId)

	return &types.QueryQMatchResponse{
		match.Timestamp,
		match.Reporter,
		match.PlayerA,
		match.PlayerB,
		match.Outcome,
	}, nil
}
