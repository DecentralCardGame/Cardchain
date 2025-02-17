package keeper

import (
	"context"

	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) CardContents(goCtx context.Context, req *types.QueryCardContentsRequest) (*types.QueryCardContentsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	var contents []*types.CardContent

	for _, cardId := range req.CardIds {
		resp, err := k.getCardContentFromId(ctx, cardId)
		if err != nil {
			return nil, err
		}
		contents = append(contents, resp)
	}

	return &types.QueryCardContentsResponse{CardContents: contents}, nil
}
