package cardservice

import (
	"strconv"
	"github.com/cosmos/cosmos-sdk/codec"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// query endpoints supported by the governance Querier
const (
	QueryResolve = "resolve"
	QueryWhois   = "whois"
)

// NewQuerier is the module level router for state queries
func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryResolve:
			return queryResolve(ctx, path[1:], req, keeper)
		case QueryWhois:
			return queryWhois(ctx, path[1:], req, keeper)
		default:
			return nil, sdk.ErrUnknownRequest("unknown cardservice query endpoint")
		}
	}
}

// CardJSON represents a card for json parsing
type CardJSON struct {
	Owner sdk.AccAddress			`json:"owner"`
	Content []byte						`json:"owner"`
	Status string							`json:"owner"`
	VotePool sdk.Coin					`json:"owner"`
	FairEnoughVotes uint64		`json:"owner"`
	OverpoweredVotes uint64		`json:"owner"`
	UnderpoweredVotes uint64	`json:"owner"`
	InappropriateVotes uint64	`json:"owner"`
	Nerflevel int64						`json:"owner"`
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

	bz, err2 := codec.MarshalJSONIndent(keeper.cdc,  card)
	if err2 != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
}

// Whois represents a name -> value lookup
type Whois struct {
	Value string         `json:"value"`
	Owner sdk.AccAddress `json:"owner"`
	Price sdk.Coins      `json:"price"`
	Type string      		 `json:"type"`
}

// nolint: unparam
func queryWhois(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) (res []byte, err sdk.Error) {
	//name := path[0]

	whois := Whois{}

	whois.Value = "" //keeper.ResolveName(ctx, name)
	whois.Owner = nil //keeper.GetOwner(ctx, name)
	//whois.Price = keeper.GetPrice(ctx, name)
	//whois.Type = keeper.GetType(ctx, name)

	bz, err2 := codec.MarshalJSONIndent(keeper.cdc, whois)
	if err2 != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
}
