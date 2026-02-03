package keeper

import (
	"context"

	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Encounters(goCtx context.Context, req *types.QueryEncountersRequest) (*types.QueryEncountersResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	var encounters []*types.Encounter
	for _, encounter := range k.Encounterk.GetAll(ctx) {
		if req.Owner == "" || encounter.Owner == req.Owner {
			encounters = append(encounters, encounter)
		}
	}

	return &types.QueryEncountersResponse{Encounters: encounters}, nil
}
