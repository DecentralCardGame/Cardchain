package cardchain

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	"github.com/DecentralCardGame/cardobject/keywords"
)

func validateGenesisCards(cards []*types.Card) error {
	for id, card := range cards {
		if card == nil {
			return fmt.Errorf("card %d: null record", id)
		}
		if len(card.Content) == 0 {
			continue
		} // Preserve existing empty ID slots.
		if err := validateGenesisCardContent(card.Content); err != nil {
			return fmt.Errorf("card %d: %w", id, err)
		}
	}
	return nil
}

func validateGenesisCardContent(content []byte) (err error) {
	// Cardobject can panic on a missing/ambiguous root. Report the card ID rather
	// than allowing invalid input to silently erase its abilities during import.
	defer func() {
		if recovered := recover(); recovered != nil {
			err = fmt.Errorf("card validation panic: %v", recovered)
		}
	}()
	var card keywords.Card
	decoder := json.NewDecoder(bytes.NewReader(content))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&card); err != nil {
		return err
	}
	var extra any
	if err := decoder.Decode(&extra); err != io.EOF {
		return fmt.Errorf("expected exactly one card JSON object")
	}
	return card.Validate()
}
