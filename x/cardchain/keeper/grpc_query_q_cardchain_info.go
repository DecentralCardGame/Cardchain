package keeper

import (
	"context"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) QCardchainInfo(goCtx context.Context, req *types.QueryQCardchainInfoRequest) (*types.QueryQCardchainInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	price := k.GetCardAuctionPrice(ctx)

	return &types.QueryQCardchainInfoResponse{
		price,
		k.GetActiveCollections(ctx),
		k.Cards.GetNumber(ctx),
		k.GetMatchesNumber(ctx),
		k.GetSellOffersNumber(ctx),
		k.Councils.GetNumber(ctx),
	}, nil
}
