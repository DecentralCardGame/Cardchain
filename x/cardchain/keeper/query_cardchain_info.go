package keeper

import (
	"context"

	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) CardchainInfo(goCtx context.Context, req *types.QueryCardchainInfoRequest) (*types.QueryCardchainInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	return &types.QueryCardchainInfoResponse{
		CardAuctionPrice: *k.CardAuctionPrice.Get(ctx),
		ActiveSets:       k.GetActiveSets(ctx),
		CardsNumber:      k.CardK.GetNum(ctx),
		MatchesNumber:    k.MatchK.GetNum(ctx),
		SellOffersNumber: k.SellOfferK.GetNum(ctx),
		CouncilsNumber:   k.Councils.GetNum(ctx),
		LastCardModified: k.LastCardModified.Get(ctx).TimeStamp,
	}, nil
}
