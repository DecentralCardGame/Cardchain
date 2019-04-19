package cardservice

import (
	"fmt"
	"strings"
	"strconv"

	"github.com/cosmos/cosmos-sdk/codec"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// query endpoints supported by the cardservice Querier
const (
	QueryResolve = "resolve"
	QueryWhois   = "whois"
	QueryNames   = "names"
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
			return queryNames(ctx, req, keeper)
		default:
			return nil, sdk.ErrUnknownRequest("unknown cardservice query endpoint")
		}
	}
}

// nolint: unparam
func queryResolve(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) (res []byte, err sdk.Error) {
	cardId, error := strconv.ParseUint(path[0], 10, 64);
	if error != nil {
		return nil, sdk.ErrUnknownRequest("could not parse cardId")
	}

	card := keeper.GetCard(ctx, cardId)

	if &card == nil {
		return []byte{}, sdk.ErrUnknownRequest("could not resolve name")
	}

	/* old code:
	bz, err2 := codec.MarshalJSONIndent(keeper.cdc, QueryResResolve{value})
	*/
	bz, err2 := codec.MarshalJSONIndent(keeper.cdc, card)
	if err2 != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
}

// TODO check if this can be removed
// Query Result Payload for a resolve query
type QueryResResolve struct {
	Value string `json:"value"`
}

// implement fmt.Stringer
func (r QueryResResolve) String() string {
	return r.Value
}

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

// implement fmt.Stringer
func (w Whois) String() string {
	return strings.TrimSpace(fmt.Sprintf(`Owner: %s
Value: %s
Price: %s`, w.Owner, w.Value, w.Price))
}

func queryNames(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) (res []byte, err sdk.Error) {
	var namesList QueryResNames

	iterator := keeper.GetNamesIterator(ctx)

	for ; iterator.Valid(); iterator.Next() {
		name := string(iterator.Key())
		namesList = append(namesList, name)
	}

	bz, err2 := codec.MarshalJSONIndent(keeper.cdc, namesList)
	if err2 != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
}

// Query Result Payload for a names query
type QueryResNames []string

// implement fmt.Stringer
func (n QueryResNames) String() string {
	return strings.Join(n[:], "\n")
}
