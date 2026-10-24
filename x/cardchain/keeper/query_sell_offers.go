package keeper

import (
	"context"

	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) SellOffers(goCtx context.Context, req *types.QuerySellOffersRequest) (*types.QuerySellOffersResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	var (
		sellOfferIds   []uint64
		sellOffersList []*types.SellOffer
	)

	iter := k.SellOfferK.GetItemIterator(ctx)
	for ; iter.Valid(); iter.Next() {
		idx, sellOffer := iter.Value()
		// Checks for price
		if req.PriceUp.IsValid() && req.PriceUp.IsValid() {
			if !(sellOffer.Price.IsGTE(req.PriceDown) && req.PriceUp.IsGTE(sellOffer.Price)) {
				continue
			}
		}

		// Checks for seller
		if req.Seller != "" {
			if req.Seller != sellOffer.Seller {
				continue
			}
		}

		// Checks for buyer
		if req.Buyer != "" {
			if req.Buyer != sellOffer.Buyer {
				continue
			}
		}

		// Checks for Status
		if req.Status != types.SellOfferStatus_empty {
			if req.Status != sellOffer.Status {
				continue
			}
		}

		// Checks for Card
		if req.Card != 0 { // I assume CardId 0 will never be used i a meeaningfull way
			if req.Card != sellOffer.Card {
				continue
			}
		}

		sellOffersList = append(sellOffersList, sellOffer)
		sellOfferIds = append(sellOfferIds, idx)
	}

	return &types.QuerySellOffersResponse{SellOffers: sellOffersList, SellOfferIds: sellOfferIds}, nil
}
