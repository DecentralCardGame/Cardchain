package keeper

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/DecentralCardGame/cardchain/util"
	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/keywords"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type candidate struct {
	id    uint64
	votes int64
}

func (k Keeper) UpdateNerfLevels(ctx sdk.Context) {
	buffbois, nerfbois, _, banbois := k.GetOPandUPCards(ctx)
	k.NerfBuffCards(ctx, buffbois, true)
	k.NerfBuffCards(ctx, nerfbois, false)
	k.UpdateBanStatus(ctx, banbois)

	k.ResetAllVotes(ctx)
	k.RemoveExpiredVoteRights(ctx)
}

// NerfBuffCards Nerfes or buffs certain cards
// TODO maybe the whole auto balancing stuff should be moved into its own file
func (k Keeper) NerfBuffCards(ctx sdk.Context, cardIds []uint64, buff bool) {
	if len(cardIds) > 0 {
		k.SetLastCardModifiedNow(ctx)
	}

	for _, val := range cardIds {
		buffCard := k.CardK.Get(ctx, val)

		cardobj, err := keywords.Unmarshal(buffCard.Content)
		if err != nil {
			k.Logger().Error("error on card content:", err, "with card", buffCard.Content)
		}

		if buffCard.BalanceAnchor {
			continue
		}

		buffnerfCost := func(cost *cardobject.CastingCost) {
			update := *cost
			if buff {
				update -= 1
			} else {
				update += 1
			}
			// only apply the buffed/nerfed value if the new value validates without error
			if update.ValidateType(nil) == nil {
				*cost = update
			}
		}

		if cardobj.Action != nil {
			buffnerfCost(&cardobj.Action.CastingCost)
		}
		if cardobj.Entity != nil {
			buffnerfCost(&cardobj.Entity.CastingCost)
		}
		if cardobj.Place != nil {
			buffnerfCost(&cardobj.Place.CastingCost)
		}
		if cardobj.Headquarter != nil {
			updateHealth := cardobj.Headquarter.Health
			updateDelay := cardobj.Headquarter.Delay
			if buff {
				updateDelay -= 1
				updateHealth += 1
			} else {
				updateDelay += 1
				updateHealth -= 1
			}
			if updateDelay.ValidateType(nil) == nil {
				cardobj.Headquarter.Delay = updateDelay
			}
			if updateHealth.ValidateType(nil) == nil {
				cardobj.Headquarter.Health = updateHealth
			}
		}

		cardJSON, _ := json.Marshal(cardobj)
		buffCard.Content = cardJSON

		if buff {
			buffCard.Nerflevel -= 1
		} else {
			buffCard.Nerflevel += 1
		}

		k.CardK.Set(ctx, val, buffCard)
	}
}

// UpdateBanStatus Bans cards
func (k Keeper) UpdateBanStatus(ctx sdk.Context, newBannedIds []uint64) {
	var err error
	// go through all cards and find already marked cards
	iter := k.CardK.GetItemIterator(ctx)
	for ; iter.Valid(); iter.Next() {
		idx, gottenCard := iter.Value()
		if gottenCard.Status == types.CardStatus_bannedVerySoon {
			gottenUser, _ := k.GetUserFromString(ctx, gottenCard.Owner)

			// remove the card from the Cards store
			var emptyCard types.Card
			k.CardK.Set(ctx, idx, &emptyCard)

			// remove the card from the ownedCards of the owner
			gottenUser.OwnedPrototypes, err = util.PopItemFromArr(idx, gottenUser.OwnedPrototypes)
			if err == nil {
				k.SetUserFromUser(ctx, gottenUser)
			} else {
				k.Logger().Error(fmt.Sprintf("trying to delete card id: %d of owner %s but does not exist", idx, gottenUser.Addr))
			}
		} else if gottenCard.Status == types.CardStatus_bannedSoon {
			gottenCard.Status = types.CardStatus_bannedVerySoon
			k.CardK.Set(ctx, idx, gottenCard)
		}
	}

	// mark freshly banned cards
	for _, id := range newBannedIds {
		CardBan := k.CardK.Get(ctx, id)
		CardBan.Status = types.CardStatus_bannedSoon
		k.CardK.Set(ctx, id, CardBan)
	}

	if len(newBannedIds) > 0 {
		k.SetLastCardModifiedNow(ctx)
	}
}

// GetOPandUPCards Gets OP and UP cards
func (k Keeper) GetOPandUPCards(ctx sdk.Context) (buffbois []uint64, nerfbois []uint64, fairbois []uint64, banbois []uint64) {
	var OPcandidates []candidate
	var UPcandidates []candidate
	var IAcandidates []candidate

	//var votingResults VotingResults
	votingResults := types.NewVotingResults()

	var uUP float64 = 0
	var uOP float64 = 0

	// go through all cards and collect candidates
	iter := k.CardK.GetItemIterator(ctx)
	for ; iter.Valid(); iter.Next() {
		id, gottenCard := iter.Value()

		nettoOP := int64(gottenCard.OverpoweredVotes - gottenCard.FairEnoughVotes - gottenCard.UnderpoweredVotes)
		nettoUP := int64(gottenCard.UnderpoweredVotes - gottenCard.FairEnoughVotes - gottenCard.OverpoweredVotes)
		nettoIA := int64(gottenCard.InappropriateVotes - gottenCard.FairEnoughVotes - gottenCard.OverpoweredVotes - gottenCard.UnderpoweredVotes)

		votingResults.TotalFairEnoughVotes += gottenCard.FairEnoughVotes
		votingResults.TotalOverpoweredVotes += gottenCard.OverpoweredVotes
		votingResults.TotalUnderpoweredVotes += gottenCard.UnderpoweredVotes
		votingResults.TotalInappropriateVotes += gottenCard.InappropriateVotes
		votingResults.TotalVotes += gottenCard.FairEnoughVotes + gottenCard.OverpoweredVotes + gottenCard.UnderpoweredVotes + gottenCard.InappropriateVotes

		// all candidates are added to the results log
		if nettoIA > 0 || nettoOP > 0 || nettoUP > 0 {
			votingResults.CardResults = append(votingResults.CardResults, &types.VotingResult{
				CardId:             id,
				FairEnoughVotes:    gottenCard.FairEnoughVotes,
				OverpoweredVotes:   gottenCard.OverpoweredVotes,
				UnderpoweredVotes:  gottenCard.UnderpoweredVotes,
				InappropriateVotes: gottenCard.InappropriateVotes,
				Result:             "fair_enough",
			})

			// sort candidates into the specific arrays
			if nettoIA > 1 {
				IAcandidates = append(IAcandidates, candidate{id: id, votes: nettoIA})
			} else if nettoOP > 0 {
				uOP += float64(nettoOP)
				OPcandidates = append(OPcandidates, candidate{id: id, votes: nettoOP})
			} else if nettoUP > 0 {
				uUP += float64(nettoUP)
				UPcandidates = append(UPcandidates, candidate{id: id, votes: nettoUP})
			}
		}
	}

	// go through all OP candidates and calculate the cutoff value and collect all above this value
	if len(OPcandidates) > 0 {
		// µ is the average, so it must be divided by n, but we can do this only after all cards are counted
		uOP /= float64(len(OPcandidates))

		sort.Slice(OPcandidates, func(i, j int) bool {
			return OPcandidates[i].votes < OPcandidates[j].votes
		})

		var giniOPsum float64
		for i := 1; i <= len(OPcandidates); i++ {
			giniOPsum += float64(OPcandidates[i-1].votes) * float64(2*i-len(OPcandidates)-1)
		}

		giniOP := giniOPsum / float64(len(OPcandidates)*len(OPcandidates)) / uOP
		cutvalue := giniOP * float64(OPcandidates[len(OPcandidates)-1].votes)

		for i := 0; i < len(OPcandidates); i++ {
			if float64(OPcandidates[i].votes) > cutvalue {
				nerfbois = append(nerfbois, OPcandidates[i].id)
			} else {
				fairbois = append(fairbois, OPcandidates[i].id)
			}
		}
	}
	// go through all UP candidates and calculate the cutoff value and collect all above this value
	if len(UPcandidates) > 0 {
		uUP /= float64(len(UPcandidates))

		sort.Slice(UPcandidates, func(i, j int) bool {
			return UPcandidates[i].votes < UPcandidates[j].votes
		})

		var giniUPsum float64
		for i := 1; i <= len(UPcandidates); i++ {
			giniUPsum += float64(UPcandidates[i-1].votes) * float64(2*i-len(UPcandidates)-1)
		}

		giniUP := giniUPsum / float64(len(UPcandidates)*len(UPcandidates)) / uUP
		cutvalue := giniUP * float64(UPcandidates[len(UPcandidates)-1].votes)

		for i := 0; i < len(UPcandidates); i++ {
			if float64(UPcandidates[i].votes) > cutvalue {
				buffbois = append(buffbois, UPcandidates[i].id)
			} else {
				fairbois = append(fairbois, UPcandidates[i].id)
			}
		}
	}
	// go through all IA candidates and collect them (there is no cutoff here)
	if len(IAcandidates) > 0 {
		for i := 0; i < len(IAcandidates); i++ {
			banbois = append(banbois, IAcandidates[i].id)
		}
	}

	// add the result to the voting log
	allBois := [][]uint64{buffbois, nerfbois, banbois}
	boisCodes := []string{"buff", "nerf", "ban"}

	for i := 0; i < len(votingResults.CardResults); i++ {
		for idx, boisCode := range boisCodes {
			for _, bois := range allBois[idx] {
				if votingResults.CardResults[i].CardId == bois {
					votingResults.CardResults[i].Result = boisCode
				}
			}
		}
	}

	// and save the log
	k.LastVotingResults.Set(ctx, &votingResults)

	return
}
