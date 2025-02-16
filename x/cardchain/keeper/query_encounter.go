package keeper

import (
	"context"

	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Encounter(goCtx context.Context, req *types.QueryEncounterRequest) (*types.QueryEncounterResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	encounter := k.encounters.Get(ctx, req.EncounterId)

	return &types.QueryEncounterResponse{Encounter: encounter}, nil
}
