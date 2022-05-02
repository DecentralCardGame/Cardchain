package keeper

import (
	"context"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) QSellOffer(goCtx context.Context, req *types.QueryQSellOfferRequest) (*types.SellOffer, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	sellOffer := k.SellOffers.Get(ctx, req.SellOfferId)

	return sellOffer, nil
}
