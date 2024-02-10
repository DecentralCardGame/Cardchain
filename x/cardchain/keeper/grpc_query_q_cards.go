package keeper

import (
	"context"
	"encoding/json"
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"

	sdkerrors "cosmossdk.io/errors"
	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/keywords"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Result struct {
	Id    uint64
	Value string
	Num   int
}

func (k Keeper) QCards(goCtx context.Context, req *types.QueryQCardsRequest) (*types.QueryQCardsResponse, error) {
	var (
		cardsList []uint64
		results   []Result
	)

	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	fmt.Printf("---> Hier: %#v\n", req)

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
		fmt.Printf("%v, %v\n", cardobjClass, req.Classes)
		if len(req.Classes) == 0 {
			return true
		}
		if bool(cardobjClass.Mysticism) && slices.Contains(req.Classes, types.CardClass_mysticism) {
			return true
		}
		if bool(cardobjClass.Nature) && slices.Contains(req.Classes, types.CardClass_nature) {
			return true
		}
		if bool(cardobjClass.Technology) && slices.Contains(req.Classes, types.CardClass_technology) {
			return true
		}
		if bool(cardobjClass.Culture) && slices.Contains(req.Classes, types.CardClass_culture) {
			return true
		}
		return false
	}

	iterator := k.Cards.GetItemIterator(ctx)

	for ; iterator.Valid(); iterator.Next() {
		idx, gottenCard := iterator.Value()

		// filter for starterCards
		if req.OnlyStarterCard && !gottenCard.StarterCard {
			continue
		}

		// filter for balance anchors
		if req.OnlyBalanceAnchors && !gottenCard.BalanceAnchor {
			continue
		}

		// then check if a status constrain was given and skip the card if it has the wrong status
		if len(req.Statuses) != 0 {
			if !slices.Contains(req.Statuses, gottenCard.Status) {
				continue
			}
		} else {
			// first skip all cards with irrelevant status
			if gottenCard.Status == types.Status_none || gottenCard.Status == types.Status_scheme {
				continue
			}
		}

		// rarity
		if len(req.Rarities) != 0 && !slices.Contains(req.Rarities, gottenCard.Rarity) {
			continue
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
		if req.NameContains != "" || len(req.CardTypes) != 0 || req.SortBy != "" || len(req.Classes) != 0 || req.KeywordsContains != "" {
			cardobj, err := keywords.Unmarshal(gottenCard.Content)
			if err != nil {
				return nil, sdkerrors.Wrap(errors.ErrJSONMarshal, err.Error()+"cardid="+strconv.FormatUint(idx, 10))
			}

			if cardobj.Action != nil {
				if len(req.CardTypes) != 0 && !slices.Contains(req.CardTypes, types.CardType_action) {
					continue
				}
				if !checkClasses(cardobj.Action.Class) {
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
				if len(req.CardTypes) != 0 && !slices.Contains(req.CardTypes, types.CardType_entity) {
					continue
				}
				if !checkClasses(cardobj.Entity.Class) {
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
				if len(req.CardTypes) != 0 && !slices.Contains(req.CardTypes, types.CardType_headquarter) {
					continue
				}
				if !checkClasses(cardobj.Headquarter.Class) {
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
				if len(req.CardTypes) != 0 && !slices.Contains(req.CardTypes, types.CardType_place) {
					continue
				}
				if !checkName(cardobj.Place.CardName) {
					continue
				}
				if !checkClasses(cardobj.Place.Class) {
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

	return &types.QueryQCardsResponse{CardsList: cardsList}, nil
}
