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

func (k Keeper) CardContent(goCtx context.Context, req *types.QueryCardContentRequest) (*types.QueryCardContentResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	cardContent, err := k.getCardContentFromId(ctx, req.CardId)

	return &types.QueryCardContentResponse{CardContent: cardContent}, err
}

func (k Keeper) getCardContentFromId(ctx sdk.Context, id uint64) (resp *types.CardContent, err error) {
	card := k.CardK.Get(ctx, id)
	if card == nil {
		return nil, errorsmod.Wrap(errors.ErrUnknownRequest, "cardId does not represent a card")
	}

	image := k.Images.Get(ctx, card.ImageId)

	return &types.CardContent{
		Content: string(card.Content),
		Hash:    image.GetHash(),
	}, nil
}
