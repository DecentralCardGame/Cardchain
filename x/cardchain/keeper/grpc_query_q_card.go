package keeper

import (
	"context"
	"strconv"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) QCard(goCtx context.Context, req *types.QueryQCardRequest) (*types.QueryQCardResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// Start of query code
	cardId, error := strconv.ParseUint(req.CardId, 10, 64)
	if error != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "could not parse cardId")
	}

	card := types.CardNoB64FromCard(k.GetCard(ctx, cardId))
	if &card == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "cardId does not represent a card")
	}

	res, err := codec.MarshalJSONIndent(codec.NewLegacyAmino(), card)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return &types.QueryQCardResponse{res}, nil
}
