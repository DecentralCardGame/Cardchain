package keeper

import (
	"context"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"golang.org/x/exp/slices"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) QSets(goCtx context.Context, req *types.QueryQSetsRequest) (*types.QueryQSetsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var (
		setIds        []uint64
		allUsersInSet bool = true
		allCardsInSet bool = true
	)

	ctx := sdk.UnwrapSDKContext(goCtx)

	iter := k.Sets.GetItemIterator(ctx)
	for ; iter.Valid(); iter.Next() {
		idx, set := iter.Value()

		// Checks for status
		if !req.IgnoreStatus {
			if req.Status != set.Status {
				continue
			}
		}

		// Checks for users contained in the contributors
		for _, user := range req.Contributors {
			if !slices.Contains(set.Contributors, user) {
				allUsersInSet = false
			}
		}
		if !allUsersInSet {
			continue
		}

		// Checks for card contained in the set
		for _, card := range req.ContainsCards {
			if !slices.Contains(set.Cards, card) {
				allCardsInSet = false
			}
		}
		if !allCardsInSet {
			continue
		}

		if req.Owner != "" {
			if req.Owner != set.Contributors[0] {
				continue
			}
		}

		setIds = append(setIds, idx)
	}

	return &types.QueryQSetsResponse{SetIds: setIds}, nil
}
