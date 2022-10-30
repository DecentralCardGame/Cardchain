package keeper

import (
	"context"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"golang.org/x/exp/slices"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) QCouncils(goCtx context.Context, req *types.QueryQCouncilsRequest) (*types.QueryQCouncilsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	var (
		councilIds           []uint64
		councilsList         []*types.Council
		allUsersInCollection bool
	)

	if req.Ignore == nil {
		newIgnore := types.NewIgnoreCouncils()
		req.Ignore = &newIgnore
	}

	iter := k.Councils.GetItemIterator(ctx)
	for ; iter.Valid(); iter.Next() {
		idx, council := iter.Value()
		allUsersInCollection = true

		// Checks for seller
		if req.Creator != "" {
			card := k.Cards.Get(ctx, council.CardId)
			if req.Creator != card.Owner {
				continue
			}
		}

		// Checks for users contained in the contributors
		for _, user := range req.Voters {
			if !slices.Contains(council.Voters, user) {
				allUsersInCollection = false
			}
		}
		if !allUsersInCollection {
			continue
		}

		// Checks for Status
		if !req.Ignore.Status {
			if req.Status != council.Status {
				continue
			}
		}

		// Checks for Card
		if !req.Ignore.Card {
			if req.Card != council.CardId {
				continue
			}
		}

		councilsList = append(councilsList, council)
		councilIds = append(councilIds, idx)
	}

	return &types.QueryQCouncilsResponse{councilIds, councilsList}, nil
}
