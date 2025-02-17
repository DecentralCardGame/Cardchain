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

func (k Keeper) Card(goCtx context.Context, req *types.QueryCardRequest) (*types.QueryCardResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	card := k.CardK.Get(ctx, req.CardId)
	if card == nil {
		return nil, errorsmod.Wrap(errors.ErrUnknownRequest, "cardId does not represent a card")
	}

	image := k.Images.Get(ctx, card.ImageId)

	return &types.QueryCardResponse{Card: &types.CardWithImage{
		Card: card, Image: string(image.Image), Hash: image.GetHash(),
	}}, nil
}
