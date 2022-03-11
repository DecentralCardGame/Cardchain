package keeper

import (
	"context"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) QMatches(goCtx context.Context, req *types.QueryQMatchesRequest) (*types.QueryQMatchesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var matchesList []uint64
	allUsersInMatch := true
	allCardsInMatch := true

	ctx := sdk.UnwrapSDKContext(goCtx)

	matches := k.GetAllMatches(ctx)
	for idx, match := range matches {
		// Checks for timestamp
		if !req.Ignore.Timestamp {
			if !(req.TimestampDown < match.Timestamp && match.Timestamp < req.TimestampUp) {
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
			if !StringItemInList(user, []string{match.PlayerA, match.PlayerB}) {
				allUsersInMatch = false
			}
		}
		if !allUsersInMatch {
			continue
		}

		// Checks for card contained in the match
		for _, card := range req.CardsPlayed {
			if !UintItemInList(card, append(match.PlayerACards, match.PlayerBCards...)) {
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

		matchesList = append(matchesList, uint64(idx))
	}

	return &types.QueryQMatchesResponse{matchesList}, nil
}
