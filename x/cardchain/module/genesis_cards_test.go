package cardchain

import (
	"bytes"
	"strings"
	"testing"

	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
)

const validCard = `{"Action":{"CardName":"Test","CastingCost":1,"Class":{"Mysticism":true},"Effects":[{"Sacrifice":{}}],"FlavourText":"","Tags":[]}}`

func TestValidateGenesisCardsPreservesContent(t *testing.T) {
	card := &types.Card{Content: []byte(validCard)}
	before := append([]byte(nil), card.Content...)
	if err := validateGenesisCards([]*types.Card{{}, card}); err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(before, card.Content) {
		t.Fatal("validation changed card content")
	}
}

func TestValidateGenesisCardsRejectsLossyOrMalformedContent(t *testing.T) {
	for _, content := range []string{
		strings.Replace(validCard, `"Sacrifice":{}`, `"Sacrifice":{"Target":"THIS"}`, 1),
		`{}`, `null`, `{broken`, validCard + ` {}`,
	} {
		card := &types.Card{Content: []byte(content)}
		if err := validateGenesisCards([]*types.Card{{}, card}); err == nil || !strings.Contains(err.Error(), "card 1:") {
			t.Fatalf("expected ID-specific error for %s, got %v", content, err)
		}
		if string(card.Content) != content {
			t.Fatal("invalid content was mutated")
		}
	}
	if err := validateGenesisCards([]*types.Card{nil}); err == nil {
		t.Fatal("null record accepted")
	}
}
