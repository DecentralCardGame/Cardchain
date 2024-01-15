package keeper

import (
	"context"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) QCardContents(goCtx context.Context, req *types.QueryQCardContentsRequest) (*types.QueryQCardContentsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	var resps []*types.QueryQCardContentResponse

	for _, cardId := range req.CardIds {
		resp, err := k.GetContentResponseFromId(ctx, cardId)
		if err != nil {
			return nil, err
		}
		resps = append(resps, resp)
	}

	return &types.QueryQCardContentsResponse{
		Cards: resps,
	}, nil
}
