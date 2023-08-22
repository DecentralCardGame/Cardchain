package keeper

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"strconv"

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

	cardId, err := strconv.ParseUint(req.CardId, 10, 64)
	if err != nil {
		return nil, sdkerrors.Wrap(errors.ErrUnknownRequest, "could not parse cardId")
	}

	card := k.Cards.Get(ctx, cardId)
	if &card == nil {
		return nil, sdkerrors.Wrap(errors.ErrUnknownRequest, "cardId does not represent a card")
	}

	image := k.Images.Get(ctx, card.ImageId)
	sum := md5.Sum(image.Image)
	hash := hex.EncodeToString(sum[:])

	k.Logger(ctx).Info(fmt.Sprintf("%v", card.Content) + " " + string(card.Content))

	return &types.QueryQCardContentResponse{
		Content: string(card.Content),
		Hash:    hash,
	}, nil
}
