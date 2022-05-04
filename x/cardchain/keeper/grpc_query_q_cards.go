package keeper

import (
	"context"
	"encoding/json"
	"sort"
	"strings"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/keywords"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Result struct {
	Id    uint64
	Value string
	Num   int
}

func (k Keeper) QCards(goCtx context.Context, req *types.QueryQCardsRequest) (*types.QueryQCardsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// This one fixes escaping
	for _, arg := range []*string{&req.Owner, &req.CardType, &req.Classes, &req.SortBy, &req.NameContains, &req.KeywordsContains, &req.NotesContains} {
		if *arg == "\"\"" {
			*arg = ""
		}
	}

	var cardsList []uint64
	var results []Result

	checkName := func(cardName cardobject.CardName) bool {
		if req.NameContains != "" && !strings.Contains(strings.ToLower(string(cardName)), strings.ToLower(req.NameContains)) {
			return false
		}
		return true
	}
	checkAbilities := func(cardAbilities string) bool {
		if strings.Contains(strings.ToLower(cardAbilities), strings.ToLower(req.KeywordsContains)) {
			return true
		}
		return false
	}
	checkClasses := func(cardobjClass cardobject.Class) bool {
		if strings.Contains(req.Classes, "OR") {
			if bool(cardobjClass.Mysticism) && strings.Contains(req.Classes, "Mysticism") {
				return true
			}
			if bool(cardobjClass.Nature) == true && strings.Contains(req.Classes, "Nature") {
				return true
			}
			if bool(cardobjClass.Technology) && strings.Contains(req.Classes, "Technology") {
				return true
			}
			if bool(cardobjClass.Culture) && strings.Contains(req.Classes, "Culture") {
				return true
			}
			return false
		} else {
			if bool(cardobjClass.Mysticism) != strings.Contains(req.Classes, "Mysticism") {
				return false
			}
			if bool(cardobjClass.Nature) != strings.Contains(req.Classes, "Nature") {
				return false
			}
			if bool(cardobjClass.Technology) != strings.Contains(req.Classes, "Technology") {
				return false
			}
			if bool(cardobjClass.Culture) != strings.Contains(req.Classes, "Culture") {
				return false
			}
			return true
		}
	}

	iterator := k.Cards.GetItemIterator(ctx)

	for ; iterator.Valid(); iterator.Next() {
		idx, gottenCard := iterator.Value()

		// first skip all cards with irrelevant status
		if gottenCard.Status == types.Status_none || gottenCard.Status == types.Status_scheme {
			continue
		}
		// then check if a status constrain was given and skip the card if it has the wrong status
		if req.Status != types.Status_none {
			if gottenCard.Status != req.Status {
				continue
			}
		}
		// Checks for playability
		switch req.Playable {
			case types.QueryQCardsRequest_Yes:
				if gottenCard.Status != types.Status_trial && gottenCard.Status != types.Status_permanent {
					continue
				}
			case types.QueryQCardsRequest_No:
				if gottenCard.Status == types.Status_trial || gottenCard.Status == types.Status_permanent {
					continue
				}
		}
		// then check if an owner constrain was given and skip the card if it has the wrong owner
		if req.Owner != "" {
			if gottenCard.Owner != req.Owner {
				continue
			}
		}
		// then check if the notes should contain something and skip the card if it does not
		if req.NotesContains != "" {
			if !strings.Contains(gottenCard.Notes, req.NotesContains) {
				continue
			}
		}

		// lastly check if this is a special request and skip the card if it does not meet it
		if req.NameContains != "" || req.CardType != "" || req.SortBy != "" || req.Classes != "" || req.KeywordsContains != "" {
			cardobj, err := keywords.Unmarshal(gottenCard.Content)
			if err != nil {
				return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
			}

			if cardobj.Action != nil {
				if req.CardType != "" && req.CardType != "Action" {
					continue
				}
				if req.Classes != "" && !checkClasses(cardobj.Action.Class) {
					continue
				}
				if !checkName(cardobj.Action.CardName) {
					continue
				}
				if req.KeywordsContains != "" {
					jsonStr, _ := json.Marshal(cardobj.Action.Effects)
					if !checkAbilities(string(jsonStr)) {
						continue
					}
				}
				if req.SortBy == "Name" {
					results = append(results, Result{idx, string(cardobj.Action.CardName), 0})
				} else if req.SortBy == "CastingCost" {
					results = append(results, Result{idx, "", int(cardobj.Action.CastingCost)})
				} else if req.SortBy == "Id" {
					cardsList = append(cardsList, idx)
				}
			}
			if cardobj.Entity != nil {
				if req.CardType != "" && req.CardType != "Entity" {
					continue
				}
				if req.Classes != "" && !checkClasses(cardobj.Entity.Class) {
					continue
				}
				if !checkName(cardobj.Entity.CardName) {
					continue
				}
				if req.KeywordsContains != "" {
					jsonStr, _ := json.Marshal(cardobj.Entity.Abilities)
					if !checkAbilities(string(jsonStr)) {
						continue
					}
				}
				if req.SortBy == "Name" {
					results = append(results, Result{idx, string(cardobj.Entity.CardName), 0})
				} else if req.SortBy == "CastingCost" {
					results = append(results, Result{idx, "", int(cardobj.Entity.CastingCost)})
				} else if req.SortBy == "Id" {
					cardsList = append(cardsList, idx)
				}
			}
			if cardobj.Headquarter != nil {
				if req.CardType != "" && req.CardType != "Headquarter" {
					continue
				}
				if req.Classes != "" && !checkClasses(cardobj.Headquarter.Class) {
					continue
				}
				if !checkName(cardobj.Headquarter.CardName) {
					continue
				}
				if req.KeywordsContains != "" {
					jsonStr, _ := json.Marshal(cardobj.Headquarter.Abilities)
					if !checkAbilities(string(jsonStr)) {
						continue
					}
				}
				if req.SortBy == "Name" {
					results = append(results, Result{idx, string(cardobj.Headquarter.CardName), 0})
				} else if req.SortBy == "CastingCost" {
					results = append(results, Result{idx, "", int(cardobj.Headquarter.Delay)})
				} else if req.SortBy == "Id" {
					cardsList = append(cardsList, idx)
				}
			}
			if cardobj.Place != nil {
				if req.CardType != "" && req.CardType != "Place" {
					continue
				}
				if !checkName(cardobj.Place.CardName) {
					continue
				}
				if req.Classes != "" && !checkClasses(cardobj.Place.Class) {
					continue
				}
				if req.KeywordsContains != "" {
					jsonStr, _ := json.Marshal(cardobj.Place.Abilities)
					if !checkAbilities(string(jsonStr)) {
						continue
					}
				}
				if req.SortBy == "Name" {
					results = append(results, Result{idx, string(cardobj.Place.CardName), 0})
				} else if req.SortBy == "CastingCost" {
					results = append(results, Result{idx, "", int(cardobj.Place.CastingCost)})
				} else if req.SortBy == "Id" {
					cardsList = append(cardsList, idx)
				}
			}
		}
		// finally if all checks were passed and unsorted, add the card
		if req.SortBy == "" {
			cardsList = append(cardsList, idx)
		}
	}

	if req.SortBy == "Name" {
		sort.Slice(results[:], func(i, j int) bool {
			return results[i].Value < results[j].Value
		})
		for _, val := range results {
			cardsList = append(cardsList, val.Id)
		}
	}
	if req.SortBy == "CastingCost" {
		sort.Slice(results[:], func(i, j int) bool {
			return results[i].Num < results[j].Num
		})
		for _, val := range results {
			cardsList = append(cardsList, val.Id)
		}
	}

	return &types.QueryQCardsResponse{cardsList}, nil
}
