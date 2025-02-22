package keeper

import (
	"context"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) QEncounters(goCtx context.Context, req *types.QueryQEncountersRequest) (*types.QueryQEncountersResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	encounters := k.Encounters.GetAll(ctx)

	return &types.QueryQEncountersResponse{Encounters: encounters}, nil
}
