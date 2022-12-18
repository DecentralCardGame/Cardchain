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
		CardAuctionPrice:  price,
		ActiveCollections: k.GetActiveCollections(ctx),
		CardsNumber:       k.Cards.GetNum(ctx),
		MatchesNumber:     k.Matches.GetNum(ctx),
		SellOffersNumber:  k.SellOffers.GetNum(ctx),
		CouncilsNumber:    k.Councils.GetNum(ctx),
	}, nil
}
