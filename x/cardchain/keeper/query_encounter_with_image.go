package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) EncounterWithImage(goCtx context.Context, req *types.QueryEncounterWithImageRequest) (*types.QueryEncounterWithImageResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	encounter := k.Encounterk.Get(ctx, req.EncounterId)
	if encounter == nil {
		return nil, errorsmod.Wrap(errors.ErrInvalidRequest, "encounterId does not represent an encounter")
	}

	return &types.QueryEncounterWithImageResponse{Encounter: &types.EncounterWithImage{
		Encounter: *encounter,
		Image:     string(k.Images.Get(ctx, encounter.ImageId).Image),
	}}, nil
}
