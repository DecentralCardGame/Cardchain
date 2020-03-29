package types

import (
	"fmt"
)

type GenesisState struct {
	CardRecords []Card `json:"cards_records"`
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
