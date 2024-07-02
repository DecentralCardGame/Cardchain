package keeper

import (
	"context"
	"crypto/md5"
	"encoding/hex"

	"github.com/cosmos/cosmos-sdk/types/errors"

	sdkerrors "cosmossdk.io/errors"
	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) QCardContent(goCtx context.Context, req *types.QueryQCardContentRequest) (*types.QueryQCardContentResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	return k.GetContentResponseFromId(ctx, req.CardId)
}

func (k Keeper) GetContentResponseFromId(ctx sdk.Context, id uint64) (resp *types.QueryQCardContentResponse, err error) {
	card := k.Cards.Get(ctx, id)
	if &card == nil {
		return nil, sdkerrors.Wrap(errors.ErrUnknownRequest, "cardId does not represent a card")
	}

	image := k.Images.Get(ctx, card.ImageId)
	sum := md5.Sum(image.Image)
	hash := hex.EncodeToString(sum[:])

	return &types.QueryQCardContentResponse{
		Content: string(card.Content),
		Hash:    hash,
	}, nil
}
