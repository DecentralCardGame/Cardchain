package keeper

import (
	"context"

	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) SellOffer(goCtx context.Context, req *types.QuerySellOfferRequest) (*types.QuerySellOfferResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	sellOffer := k.SellOfferK.Get(ctx, req.SellOfferId)

	return &types.QuerySellOfferResponse{SellOffer: sellOffer}, nil
}
