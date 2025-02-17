package keeper

import (
	"context"
	"slices"

	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Sets(goCtx context.Context, req *types.QuerySetsRequest) (*types.QuerySetsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	var (
		setIds        []uint64
		allUsersInSet bool = true
		allCardsInSet bool = true
	)

	iter := k.Setk.GetItemIterator(ctx)
	for ; iter.Valid(); iter.Next() {
		idx, set := iter.Value()

		// Checks for status
		if req.Status != types.SetStatus_undefined {
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

	return &types.QuerySetsResponse{SetIds: setIds}, nil
}
