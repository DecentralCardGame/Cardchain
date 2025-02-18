package keeper

import (
	"context"

	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) EncountersWithImage(goCtx context.Context, req *types.QueryEncountersWithImageRequest) (*types.QueryEncountersWithImageResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	var encountersWithImage []*types.EncounterWithImage

	encounters := k.Encounterk.GetAll(ctx)

	for _, encounter := range encounters {
		encountersWithImage = append(encountersWithImage, &types.EncounterWithImage{
			Encounter: encounter,
			Image:     string(k.Images.Get(ctx, encounter.ImageId).Image),
		})
	}

	return &types.QueryEncountersWithImageResponse{Encounters: encountersWithImage}, nil
}
