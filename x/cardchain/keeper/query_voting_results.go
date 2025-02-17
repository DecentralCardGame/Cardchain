package keeper

import (
	"context"

	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) VotingResults(goCtx context.Context, req *types.QueryVotingResultsRequest) (*types.QueryVotingResultsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	votingResults := k.LastVotingResults.Get(ctx)

	return &types.QueryVotingResultsResponse{LastVotingResults: votingResults}, nil
}
