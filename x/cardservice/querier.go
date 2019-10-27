package cardservice

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/codec"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// query endpoints supported by the cardservice Querier
const (
	QueryResolve      = "resolve"
	QueryWhois        = "whois"
	QueryNames        = "cards"
	QueryVotableCards = "votable-cards"
)

// NewQuerier is the module level router for state queries
func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryResolve:
			return queryResolve(ctx, path[1:], req, keeper)
		case QueryWhois:
			return queryWhois(ctx, path[1:], req, keeper)
		case QueryNames:
			return queryCards(ctx, req, keeper)
		case QueryVotableCards:
			return queryVotableCards(ctx, path[1:], req, keeper)
		case QueryRegisterUser:
			return queryRegisterUser(ctx, path[1:], req, keeper)
		default:
			return nil, sdk.ErrUnknownRequest("unknown cardservice query endpoint")
		}
	}
}

// nolint: unparam
func queryResolve(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) (res []byte, err sdk.Error) {
	cardId, error := strconv.ParseUint(path[0], 10, 64)
	if error != nil {
		return nil, sdk.ErrUnknownRequest("could not parse cardId")
	}

	card := keeper.GetCard(ctx, cardId)

	if &card == nil {
		return []byte{}, sdk.ErrUnknownRequest("cardId does not represent a card")
	}

	bz, err2 := codec.MarshalJSONIndent(keeper.cdc, card)
	if err2 != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
}

/*
// TODO check if this can be removed
// Query Result Payload for a resolve query
type QueryResResolve struct {
	Value string `json:"value"`
}

// implement fmt.Stringer
func (r QueryResResolve) String() string {
	return r.Value
}
*/

// nolint: unparam
func queryWhois(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) (res []byte, err sdk.Error) {
	address, error := sdk.AccAddressFromBech32(path[0])
	if error != nil {
		return nil, sdk.ErrUnknownRequest("could not parse user address")
	}

	user := keeper.GetUser(ctx, address)

	bz, err2 := codec.MarshalJSONIndent(keeper.cdc, user)
	if err2 != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
}

// nolint: unparam
func queryRegisterUser(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) (res []byte, err sdk.Error) {
	address, error := sdk.AccAddressFromBech32(path[0])
	if error != nil {
		return nil, sdk.ErrUnknownRequest("could not parse user address")
	}

	user := keeper.CreateUser(ctx, address)

	bz, err2 := codec.MarshalJSONIndent(keeper.cdc, user)
	if err2 != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
}

// implement fmt.Stringer
func (w Whois) String() string {
	return strings.TrimSpace(fmt.Sprintf(`Owner: %s
Value: %s
Price: %s`, w.Owner, w.Value, w.Price))
}

func queryCards(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) (res []byte, err sdk.Error) {
	var cardsList QueryResCards

	iterator := keeper.GetCardsIterator(ctx)

	for ; iterator.Valid(); iterator.Next() {

		var gottenCard Card
		keeper.cdc.MustUnmarshalBinaryBare(iterator.Value(), &gottenCard)

		// TODO check if json.Marshal is fair enough here
		b, err := json.Marshal(gottenCard)
		if err != nil {
			panic("could not marshal gottenCard to JSON")
		}

		cardsList = append(cardsList, string(b))
	}

	bz, err2 := codec.MarshalJSONIndent(keeper.cdc, cardsList)
	if err2 != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
}

// Query Result Payload for a names query
type QueryResCards []string

// implement fmt.Stringer
func (n QueryResCards) String() string {
	return strings.Join(n[:], "\n")
}

func queryVotableCards(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) (res []byte, err sdk.Error) {
	address, error := sdk.AccAddressFromBech32(path[0])
	if error != nil {
		return nil, sdk.ErrUnknownRequest("could not parse user address")
	}

	user := keeper.GetUser(ctx, address)

	var cardsList QueryResCards

	iterator := keeper.GetCardsIterator(ctx)

	for ; iterator.Valid(); iterator.Next() {

		var gottenCard Card
		cardId := binary.BigEndian.Uint64(iterator.Key())
		keeper.cdc.MustUnmarshalBinaryBare(iterator.Value(), &gottenCard)

		// TODO check if json.Marshal is fair enough here
		b, err := json.Marshal(gottenCard)
		if err != nil {
			panic("could not marshal gottenCard to JSON")
		}

		var voteRight = SearchVoteRights(cardId, user.VoteRights)
		if voteRight >= 0 {
			cardsList = append(cardsList, string(b))
		}
	}

	bz, err2 := codec.MarshalJSONIndent(keeper.cdc, cardsList)
	if err2 != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
}
