package keeper

import (
	"context"
	"slices"

	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Matches(goCtx context.Context, req *types.QueryMatchesRequest) (*types.QueryMatchesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var (
		matchesList []*types.Match
		matchIds    []uint64
	)

	ctx := sdk.UnwrapSDKContext(goCtx)

	iter := k.matches.GetItemIterator(ctx)
	for ; iter.Valid(); iter.Next() {
		allUsersInMatch := true
		allCardsInMatch := true

		// Checks for timestamp
		idx, match := iter.Value()
		if req.TimestampUp != 0 || req.TimestampDown != 0 {
			if !(req.TimestampDown <= match.Timestamp && match.Timestamp <= req.TimestampUp) {
				continue
			}
		}

		// Checks for outcome
		if req.Outcome != types.Outcome_Undefined {
			if req.Outcome != match.Outcome {
				continue
			}
		}

		// Checks for users contained in the match
		for _, user := range req.ContainsUsers {
			if !slices.Contains([]string{match.PlayerA.Addr, match.PlayerB.Addr}, user) {
				allUsersInMatch = false
			}
		}
		if !allUsersInMatch {
			continue
		}

		// Checks for card contained in the match
		for _, card := range req.CardsPlayed {
			if !slices.Contains(append(match.PlayerA.PlayedCards, match.PlayerB.PlayedCards...), card) {
				allCardsInMatch = false
			}
		}
		if !allCardsInMatch {
			continue
		}

		// Checks for reporter
		if req.Reporter != "" {
			if match.Reporter != req.Reporter {
				continue
			}
		}

		matchIds = append(matchIds, idx)
		matchesList = append(matchesList, match)
	}

	return &types.QueryMatchesResponse{Matches: matchesList, MatchIds: matchIds}, nil
}
