package keeper

import (
	"context"
	"encoding/json"
	"slices"
	"sort"
	"strconv"
	"strings"

	sdkerrors "cosmossdk.io/errors"
	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
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

func (k Keeper) Cards(goCtx context.Context, req *types.QueryCardsRequest) (*types.QueryCardsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var (
		cardsList []uint64
		results   []Result
	)

	ctx := sdk.UnwrapSDKContext(goCtx)

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
	checkClass := func(cardobjClass cardobject.Class) bool {
		if len(req.Class) == 0 {
			return true
		}
		if !req.MultiClassOnly {
			if bool(cardobjClass.Mysticism) && slices.Contains(req.Class, types.CardClass_mysticism) {
				return true
			}
			if bool(cardobjClass.Nature) && slices.Contains(req.Class, types.CardClass_nature) {
				return true
			}
			if bool(cardobjClass.Technology) && slices.Contains(req.Class, types.CardClass_technology) {
				return true
			}
			if bool(cardobjClass.Culture) && slices.Contains(req.Class, types.CardClass_culture) {
				return true
			}
			return false
		} else {
			if bool(cardobjClass.Mysticism) != slices.Contains(req.Class, types.CardClass_mysticism) {
				return false
			}
			if bool(cardobjClass.Nature) != slices.Contains(req.Class, types.CardClass_nature) {
				return false
			}
			if bool(cardobjClass.Technology) != slices.Contains(req.Class, types.CardClass_technology) {
				return false
			}
			if bool(cardobjClass.Culture) != slices.Contains(req.Class, types.CardClass_culture) {
				return false
			}
			return true
		}
	}

	iterator := k.cards.GetItemIterator(ctx)

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
		if len(req.Status) != 0 {
			if !slices.Contains(req.Status, gottenCard.Status) {
				continue
			}
		} else {
			// first skip all cards with irrelevant status
			if gottenCard.Status == types.CardStatus_none || gottenCard.Status == types.CardStatus_scheme {
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
		if req.NameContains != "" || len(req.CardType) != 0 || req.SortBy != "" || len(req.Class) != 0 || req.KeywordsContains != "" {
			cardobj, err := keywords.Unmarshal(gottenCard.Content)
			if err != nil {
				return nil, sdkerrors.Wrap(errors.ErrJSONMarshal, err.Error()+" cardid="+strconv.FormatUint(idx, 10))
			}

			if cardobj.Action != nil {
				if len(req.CardType) != 0 && !slices.Contains(req.CardType, types.CardType_action) {
					continue
				}
				if !checkClass(cardobj.Action.Class) {
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
				if len(req.CardType) != 0 && !slices.Contains(req.CardType, types.CardType_entity) {
					continue
				}
				if !checkClass(cardobj.Entity.Class) {
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
				if len(req.CardType) != 0 && !slices.Contains(req.CardType, types.CardType_headquarter) {
					continue
				}
				if !checkClass(cardobj.Headquarter.Class) {
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
				if len(req.CardType) != 0 && !slices.Contains(req.CardType, types.CardType_place) {
					continue
				}
				if !checkName(cardobj.Place.CardName) {
					continue
				}
				if !checkClass(cardobj.Place.Class) {
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

	return &types.QueryCardsResponse{}, nil
}
