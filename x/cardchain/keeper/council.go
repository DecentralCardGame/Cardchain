package keeper

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"sort"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetResponseHash generates a hash form a response and a secret
func GetResponseHash(response types.Response, secret string) string {
	clearResponse := response.String() + secret
	hashResponse := sha256.Sum256([]byte(clearResponse))
	hashStringResponse := hex.EncodeToString(hashResponse[:])
	return hashStringResponse
}

// GetCouncil Gets a certain council from store
func (k Keeper) GetCouncil(ctx sdk.Context, councilId uint64) (gottenCouncil types.Council) {
	store := ctx.KVStore(k.CouncilsStoreKey)
	bz := store.Get(sdk.Uint64ToBigEndian(councilId))

	k.cdc.MustUnmarshal(bz, &gottenCouncil)
	return
}

// SetCouncil Sets a certain council in store
func (k Keeper) SetCouncil(ctx sdk.Context, councilId uint64, newCouncil types.Council) {
	store := ctx.KVStore(k.CouncilsStoreKey)
	store.Set(sdk.Uint64ToBigEndian(councilId), k.cdc.MustMarshal(&newCouncil))
}

// GetCouncilsIterator Returns an interator for all councils
func (k Keeper) GetCouncilsIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.CouncilsStoreKey)
	return sdk.KVStorePrefixIterator(store, nil)
}

// GetAllCouncils Gets all councils from store
func (k Keeper) GetAllCouncils(ctx sdk.Context) (allCouncils []*types.Council) {
	iterator := k.GetCouncilsIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {

		var gottenCouncil types.Council
		k.cdc.MustUnmarshal(iterator.Value(), &gottenCouncil)

		allCouncils = append(allCouncils, &gottenCouncil)
	}
	return
}

// GetCouncilsNumber Gets the number of all existing councils
func (k Keeper) GetCouncilsNumber(ctx sdk.Context) (councilId uint64) {
	iterator := k.GetCouncilsIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {
		councilId++
	}
	return
}

// GetCouncilPartiedVoters Sorts the voters of a council in approvers, deniers and suggestors
func GetCouncilPartiedVoters(responses []*types.WrapClearResponse) (approvers []string, deniers []string, suggestors []string) {
	for _, response := range responses {
		switch response.Response {
		case types.Response_Yes:
			approvers = append(approvers, response.User)
		case types.Response_No:
			deniers = append(deniers, response.User)
		case types.Response_Suggestion:
			suggestors = append(suggestors, response.User)
		}
	}
	return
}

// TryEvaluate Tries to evaluate the council decision
func (k Keeper) TryEvaluate(ctx sdk.Context, council types.Council) (types.Council, error) {
	collateralDeposit := k.GetParams(ctx).CollateralDeposit
	bounty := MulCoin(collateralDeposit, 2)
	votePool := MulCoin(collateralDeposit, 5)

	if len(council.ClearResponses) == 5 {
		approvers, deniers, suggestors := GetCouncilPartiedVoters(council.ClearResponses)
		nrYes, nrNo, nrSuggestion := len(approvers), len(deniers), len(suggestors)

		if nrNo == nrYes || nrSuggestion > nrNo || nrSuggestion > nrYes {
			council.Status = types.CouncelingStatus_suggestionsMade
			return council, nil
		} else if nrNo > nrYes {
			for _, user := range deniers {
				err := k.MintCoinsToString(ctx, user, sdk.Coins{bounty})
				if err != nil {
					return council, err
				}
				council.Treasury = council.Treasury.Sub(bounty)
			}
			k.AddPoolCredits(ctx, PublicPoolKey, council.Treasury)
			council.Treasury = council.Treasury.Sub(council.Treasury)
			council.Status = types.CouncelingStatus_councilClosed
		} else if nrNo < nrYes {
			card := k.GetCard(ctx, council.CardId)
			card.ResetVotes()
			card.VotePool = card.VotePool.Add(votePool)
			card.Status = types.Status_trial
			k.SetCard(ctx, council.CardId, card)
			council.TrialStart = uint64(ctx.BlockHeight())
			council.Treasury = council.Treasury.Sub(votePool)
		}
		for _, addr := range council.Voters {
			user, err := k.GetUserFromString(ctx, addr)
			if err != nil {
				return council, err
			}
			user.CouncilStatus = types.CouncilStatus_available
			user.Cards = append(user.Cards, council.CardId)
			k.SetUserFromUser(ctx, user)
		}
	}
	return council, nil
}

// CheckTrial Checks for councils that shouldn't be in trial anymore
func (k Keeper) CheckTrial(ctx sdk.Context) error {
	collateralDeposit := k.GetParams(ctx).CollateralDeposit
	for idx, council := range k.GetAllCouncils(ctx) {
		if council.Status == types.CouncelingStatus_revealed {
			if council.TrialStart+k.GetParams(ctx).TrialPeriod <= uint64(ctx.BlockHeight()) {
				card := k.GetCard(ctx, council.CardId)

				var (
					group []string
					amt                       int64
				)

				approvers, deniers, _ := GetCouncilPartiedVoters(council.ClearResponses)

				votes := []uint64{card.FairEnoughVotes, card.OverpoweredVotes, card.UnderpoweredVotes, card.InappropriateVotes}
				sort.Slice(votes, func(i, j int) bool {
					return votes[i] < votes[j]
				})
				if votes[len(votes)-1] == 0 {
					council.TrialStart = uint64(ctx.BlockHeight())
					k.SetCouncil(ctx, uint64(idx), *council)
					continue
				}
				if card.FairEnoughVotes == votes[len(votes)-1] {
					card.Status = types.Status_permanent
					group = approvers
					amt = 2
				} else {
					card.Status = types.Status_prototype
					group = deniers
					amt = 3
				}

				k.Logger(ctx).Debug(fmt.Sprintf(":: Card Set to %s", card.Status.String()))

				bounty := MulCoin(collateralDeposit, amt)
				for _, user := range group {
					err := k.MintCoinsToString(ctx, user, sdk.Coins{bounty})
					if err != nil {
						return nil
					}
					council.Treasury = council.Treasury.Sub(bounty)
				}
				k.AddPoolCredits(ctx, PublicPoolKey, council.Treasury)
				council.Treasury = council.Treasury.Sub(council.Treasury)

				incentive := QuoCoin(card.VotePool, int64(len(card.Voters)))
				for _, user := range card.Voters {
					err := k.MintCoinsToString(ctx, user, sdk.Coins{incentive})
					if err != nil {
						return nil
					}
				}
				card.VotePool = card.VotePool.Sub(card.VotePool)
				council.Status = types.CouncelingStatus_councilClosed

				k.SetCouncil(ctx, uint64(idx), *council)
				k.SetCard(ctx, council.CardId, card)
			}
		}
	}
	return nil
}
