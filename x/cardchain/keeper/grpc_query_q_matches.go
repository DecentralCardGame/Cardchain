package keeper

import (
	"context"
	"golang.org/x/exp/slices"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) QMatches(goCtx context.Context, req *types.QueryQMatchesRequest) (*types.QueryQMatchesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var (
		matchesList []*types.Match
		matchIds    []uint64
	)
	allUsersInMatch := true
	allCardsInMatch := true

	ctx := sdk.UnwrapSDKContext(goCtx)

	matches := k.Matches.GetAll(ctx)
	for idx, match := range matches {
		// Checks for timestamp
		if !req.Ignore.Timestamp {
			if !(req.TimestampDown <= match.Timestamp && match.Timestamp <= req.TimestampUp) {
				continue
			}
		}

		// Checks for outcome
		if !req.Ignore.Outcome {
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
		if !req.Ignore.Reporter {
			if match.Reporter != req.Reporter {
				continue
			}
		}

		matchIds = append(matchIds, uint64(idx))
		matchesList = append(matchesList, match)
	}

	return &types.QueryQMatchesResponse{matchIds, matchesList}, nil
}
