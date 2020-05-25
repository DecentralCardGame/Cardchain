package keeper

import (
	"encoding/binary"
	//"encoding/json"
	"fmt"
	"strconv"
	"strings"

	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/DecentralCardGame/Cardchain/x/cardservice/internal/types"
)

// query endpoints supported by the cardservice Querier
const (
	QueryCard      	 	 = "card"
	QueryUser          = "user"
	QueryCards         = "cards"
	QueryVotableCards  = "votable-cards"
	QueryCardchainInfo = "cardchain-info"
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
			return queryCards(ctx, req, keeper)
		case QueryVotableCards:
			return queryVotableCards(ctx, path[1:], req, keeper)
		case QueryCardchainInfo:
			return queryCardchainInfo(ctx, req, keeper)
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

			//fmt.Println(string(res))
			//return nil
	card := keeper.GetCard(ctx, cardId)

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

// TODO this should be changed to query card ids
func queryCards(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) ([]byte, error) {
	var cardsList []uint64

	iterator := keeper.GetCardsIterator(ctx)

	for ; iterator.Valid(); iterator.Next() {

		var gottenCard types.Card
		keeper.cdc.MustUnmarshalBinaryBare(iterator.Value(), &gottenCard)

		cardsList = append(cardsList, binary.BigEndian.Uint64(iterator.Key()))
	}

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
	/*
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
	*/

	res, err := codec.MarshalJSONIndent(keeper.cdc, user.VoteRights)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func queryCardchainInfo(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) ([]byte, error) {
	info := keeper.GetCardAuctionPrice(ctx)

	res, err := codec.MarshalJSONIndent(keeper.cdc, info)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}
