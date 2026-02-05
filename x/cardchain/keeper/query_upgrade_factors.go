package keeper

import (
	"context"

	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) UpgradeFactors(goCtx context.Context, req *types.QueryUpgradeFactorsRequest) (*types.QueryUpgradeFactorsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	return &types.QueryUpgradeFactorsResponse{UpgradeFactors: k.UpgradeFactorK.GetAll(ctx)}, nil
}
