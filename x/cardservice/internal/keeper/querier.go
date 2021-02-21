package keeper

import (
	"encoding/binary"
	//"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"

	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/DecentralCardGame/Cardchain/x/cardservice/internal/types"
	"github.com/DecentralCardGame/cardobject"
)

// query endpoints supported by the cardservice Querier
const (
	QueryCard      	 	 = "card"
	QueryUser          = "user"
	QueryCards         = "cards"
	QueryCardSVG       = "cardsvg"
	QueryVotableCards  = "votable-cards"
	QueryCardchainInfo = "cardchain-info"
	QueryVotingResults = "cardchain-votingresults"
)

// NewQuerier is the module level router for state queries
func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err error) {
		switch path[0] {
		case QueryCard:
			return queryCard(ctx, path[1:], req, keeper)
		case QueryUser:
			return queryUser(ctx, path[1:], req, keeper)
		case QueryCards:
			return queryCards(ctx, path[1], path[2], path[3], path[4], path[5], path[6], req, keeper)
		case QueryCardSVG:
			return queryCardSVG(ctx, path[1:], req, keeper)
		case QueryVotableCards:
			return queryVotableCards(ctx, path[1:], req, keeper)
		case QueryCardchainInfo:
			return queryCardchainInfo(ctx, req, keeper)
		case QueryVotingResults:
			return queryVotingResults(ctx, req, keeper)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown cardservice query endpoint")
		}
	}
}

// nolint: unparam
func queryCard(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, error) {
	cardId, error := strconv.ParseUint(path[0], 10, 64)
	if error != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "could not parse cardId")
	}

	card := types.CardNoB64FromCard(keeper.GetCard(ctx, cardId))

	if &card == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "cardId does not represent a card")
	}

	res, err := codec.MarshalJSONIndent(keeper.cdc, card)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

// nolint: unparam
func queryCardSVG(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, error) {
	cardId, error := strconv.ParseUint(path[0], 10, 64)
	if error != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "could not parse cardId")
	}

	card := keeper.GetCard(ctx, cardId)

	if &card == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "cardId does not represent a card")
	}

	cardobj, err := cardobject.NewCardFromJson(string(card.Content))
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	if cardobj.Action != nil {
		return []byte(cardobj.Action.CardName), nil
	}
	if cardobj.Entity != nil {
		fmt.Println(cardobj.Entity.CardName)
		return []byte(cardobj.Entity.CardName), nil
	}
	if cardobj.Headquarter != nil {
		return []byte(cardobj.Headquarter.CardName), nil
	}
	if cardobj.Place != nil {
		return []byte(cardobj.Place.CardName), nil
	}

	res, err := codec.MarshalJSONIndent(keeper.cdc, card)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

// nolint: unparam
func queryUser(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, error) {
	address, error := sdk.AccAddressFromBech32(path[0])
	if error != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "could not parse user address")
	}

	user := keeper.GetUser(ctx, address)

	fmt.Println(user)

	res, err := codec.MarshalJSONIndent(keeper.cdc, user)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func queryCards(ctx sdk.Context, owner string, status string, cardType string, classes string, sortBy string, nameContains string, notesContains string, req abci.RequestQuery, keeper Keeper) ([]byte, error) {
	var cardsList []uint64

	type Result struct {
		Id uint64
		Value string
	}
	var results []Result

	iterator := keeper.GetCardsIterator(ctx)

	for ; iterator.Valid(); iterator.Next() {

		var gottenCard types.Card
		keeper.cdc.MustUnmarshalBinaryBare(iterator.Value(), &gottenCard)

		// first skip all cards with irrelevant status
		if gottenCard.Status == "" || gottenCard.Status == "scheme" {
			continue
		}
		// then check if a status constrain was given and skip the card if it has the wrong status
		if status != "" {
			if gottenCard.Status != status {
				continue
			}
		}
		// then check if an owner constrain was given and skip the card if it has the wrong owner
		if owner != "" {
			if gottenCard.Owner.String() != owner {
				continue
			}
		}
		// then check if the notes should contain something and skip the card if it does not
		if notesContains != "" {
			if !strings.Contains(gottenCard.Notes, notesContains) {
				continue
			}
		}

		// lastly check if the name should contain a certain string and skip the card if it does not
		if nameContains != "" || cardType != "" || sortBy != "" {
			cardobj, err := cardobject.NewCardFromJson(string(gottenCard.Content))
			if err != nil {
				return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
			}
			fmt.Println("sortby:", sortBy)


			if cardobj.Action != nil {
				if cardType != "" && cardType != "Action" {
					continue
				}
				if !strings.Contains(cardobj.Action.CardName, nameContains) {
					continue
				}
				if sortBy == "Name" {
					results = append(results, Result{binary.BigEndian.Uint64(iterator.Key()), cardobj.Action.CardName})
				} else if sortBy == "CastingCost" {
					results = append(results, Result{binary.BigEndian.Uint64(iterator.Key()), strconv.Itoa(cardobj.Action.CastingCost)})
				}
			}
			if cardobj.Entity != nil {
				if cardType != "" && cardType != "Entity" {
					continue
				}
				if !strings.Contains(cardobj.Entity.CardName, nameContains) {
					continue
				}
				if sortBy == "Name" {
					results = append(results, Result{binary.BigEndian.Uint64(iterator.Key()), cardobj.Entity.CardName})
				} else if sortBy == "CastingCost" {
					fmt.Println("castingcost:", cardobj.Entity.CastingCost)
					fmt.Println(strconv.Itoa(cardobj.Entity.CastingCost))
					results = append(results, Result{binary.BigEndian.Uint64(iterator.Key()), strconv.Itoa(cardobj.Entity.CastingCost)})
				}
			}
			if cardobj.Headquarter != nil {
				fmt.Println("cardobj:", cardobj.Headquarter)
				if cardType != "" && cardType != "Headquarter" {
					continue
				}
				if !strings.Contains(cardobj.Headquarter.CardName, nameContains) {
					continue
				}
				if sortBy == "Name" {
					results = append(results, Result{binary.BigEndian.Uint64(iterator.Key()), cardobj.Headquarter.CardName})
				} else if sortBy == "CastingCost" {
					results = append(results, Result{binary.BigEndian.Uint64(iterator.Key()), "0"})
				}
			}
			if cardobj.Place != nil {
				if cardType != "" && cardType != "Place" {
					continue
				}
				if !strings.Contains(cardobj.Place.CardName, nameContains) {
					continue
				}
				if sortBy == "Name" {
					results = append(results, Result{binary.BigEndian.Uint64(iterator.Key()), cardobj.Place.CardName})
				} else if sortBy == "CastingCost" {
					fmt.Println("castingcost:", cardobj.Place.CastingCost)
					results = append(results, Result{binary.BigEndian.Uint64(iterator.Key()), strconv.Itoa(cardobj.Place.CastingCost)})
				}
			}
		}
		// finally if all checks were passed and unsorted, add the card
		if sortBy == "" {
			cardsList = append(cardsList, binary.BigEndian.Uint64(iterator.Key()))
		}
	}

	fmt.Println(results)

	if sortBy	!= "" {
		fmt.Println(results)
		sort.Slice(results[:], func(i, j int) bool {
			return results[i].Value < results[j].Value
		})

		for _,val := range results {
			cardsList = append(cardsList, val.Id)
		}
	}
	fmt.Println(cardsList)

	res, err2 := codec.MarshalJSONIndent(keeper.cdc, cardsList)
	if err2 != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err2.Error())
	}

	return res, nil
}

// Query Result Payload for a cards query
type QueryResCards []string

// implement fmt.Stringer
func (n QueryResCards) String() string {
	return strings.Join(n[:], "\n")
}

// TODO wtf? remove endless comments?
func queryVotableCards(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, error) {
	address, error := sdk.AccAddressFromBech32(path[0])
	if error != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "could not parse user address")
	}

	user := keeper.GetUser(ctx, address)

	if user.Alias == "" {
		return []byte("{\"unregistered\": true}"), nil
	}

	if len(user.VoteRights) == 0 {
		return []byte("{\"noVoteRights\": true}"), nil
	} else {
		res, err := codec.MarshalJSONIndent(keeper.cdc, user.VoteRights)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
		}
		return res, nil
	}

}

func queryCardchainInfo(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) ([]byte, error) {
	info := keeper.GetCardAuctionPrice(ctx)

	res, err := codec.MarshalJSONIndent(keeper.cdc, info)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func queryVotingResults(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) ([]byte, error) {
	info := keeper.GetLastVotingResults(ctx)

	res, err := codec.MarshalJSONIndent(keeper.cdc, info)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}
