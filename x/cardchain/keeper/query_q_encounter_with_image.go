package keeper

import (
	"context"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) QEncounterWithImage(goCtx context.Context, req *types.QueryQEncounterWithImageRequest) (*types.QueryQEncounterWithImageResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	encounter := k.Encounters.Get(ctx, req.Id)

	return &types.QueryQEncounterWithImageResponse{Encounter: &types.EncounterWithImage{
		Encounter: encounter,
		Image:     string(k.Images.Get(ctx, encounter.ImageId).Image),
	}}, nil
}
