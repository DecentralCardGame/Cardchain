package cardservice

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

type GenesisState struct {
	CardRecords []Card `json:"whois_records"`
}

func NewGenesisState(cardsRecords []Card) GenesisState {
	return GenesisState{CardRecords: nil}
}

func ValidateGenesis(data GenesisState) error {
	for _, record := range data.CardRecords {
		if record.Owner == nil {
			return fmt.Errorf("invalid CardRecord: Owner: %s. Error: Missing Owner", record.Owner)
		}
		if record.Content == nil {
			return fmt.Errorf("invalid CardRecord: Content: %s. Error: Missing Content", record.Content)
		}
		if record.Status == "" {
			return fmt.Errorf("invalid CardRecord: Status: %s. Error: Missing Status", record.Status)
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
		currId := lastId + 1                      // iterate it by 1

		keeper.SetLastCardSchemeId(ctx, currId)
		keeper.SetCard(ctx, currId, record)
		keeper.AddOwnedCardScheme(ctx, currId, record.Owner)
	}
	return []abci.ValidatorUpdate{}
}

func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	records := k.GetAllCards(ctx)
	return GenesisState{CardRecords: records}
}
