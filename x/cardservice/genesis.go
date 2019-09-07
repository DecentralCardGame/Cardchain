package cardservice

import (
	"fmt"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

type GenesisState struct {
	CardRecords []Card `json:"card_records"`
}

func NewGenesisState(cardRecords []Card) GenesisState {
	return GenesisState{CardRecords: nil}
}

func ValidateGenesis(data GenesisState) error {
	for _, record := range data.CardRecords {
		if record.Owner == nil {
			return fmt.Errorf("Error: Missing Owner")
		}
		if record.Content == nil {
			return fmt.Errorf("Error: Missing Content")
		}
		if record.Status == "" {
			return fmt.Errorf("Error: Missing Price")
		}
	}
	return nil
}

func DefaultGenesisState() GenesisState {
	return GenesisState{
		CardRecords: []Card{},
	}
}

func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) []abci.ValidatorUpdate {
	for _, record := range data.CardRecords {
		lastId := keeper.GetLastCardSchemeId(ctx) // first get last card bought id
		currId := lastId+1
		keeper.SetCard(ctx, currId, record)
	}
	return []abci.ValidatorUpdate{}
}

func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	var records []Card
	iterator := k.GetCardsIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {

		cardId, error := strconv.ParseUint(string(iterator.Key()), 10, 64)
		if error != nil {
			fmt.Errorf("Error: could not parse cardId: ", iterator.Key())
			return GenesisState{}
		}
		card := k.GetCard(ctx, cardId)
		records = append(records, card)

	}
	return GenesisState{CardRecords: records}
}
