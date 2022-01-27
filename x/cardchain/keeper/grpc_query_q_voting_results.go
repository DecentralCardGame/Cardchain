package keeper

import (
	"context"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) QVotingResults(goCtx context.Context, req *types.QueryQVotingResultsRequest) (*types.QueryQVotingResultsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	vResults := k.GetLastVotingResults(ctx)

	return &types.QueryQVotingResultsResponse{&vResults}, nil
}
