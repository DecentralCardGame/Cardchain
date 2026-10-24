package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/gogoproto/proto"
)

func ValidateRunDeckAddress(value string) error {
	addr, err := sdk.AccAddressFromBech32(value)
	if err != nil || addr.String() != value {
		return fmt.Errorf("invalid canonical account address")
	}
	return nil
}
func ValidateRunDeckIdentity(owner, reporter, ref string) error {
	if err := ValidateRunDeckAddress(owner); err != nil {
		return err
	}
	if err := ValidateRunDeckAddress(reporter); err != nil {
		return err
	}
	if len(ref) == 0 || len(ref) > 128 || strings.TrimSpace(ref) != ref {
		return fmt.Errorf("run reference requires 1..128 characters")
	}
	for _, c := range ref {
		if c < 33 || c > 126 {
			return fmt.Errorf("run reference must be printable ASCII without whitespace")
		}
	}
	return nil
}

func ValidateRunDeckGenesis(state *RunDeckGenesis) error {
	if state == nil {
		return nil
	}
	completed := map[string]bool{}
	for i, deck := range state.Decks {
		if deck == nil || deck.Id != uint64(i) {
			return fmt.Errorf("deck IDs must be contiguous from zero")
		}
		if err := ValidateRunDeckIdentity(deck.Owner, deck.Reporter, deck.RunRef); err != nil {
			return err
		}
		key := deck.Reporter + "\x00" + deck.RunRef
		if completed[key] {
			return fmt.Errorf("duplicate completed run")
		}
		completed[key] = true
		if err := ValidateRunDeckCards(deck.RecipeFormatVersion, deck.Cards); err != nil {
			return fmt.Errorf("deck %d: %w", i, err)
		}
		for _, card := range deck.Cards {
			for _, raw := range card.ForgeRecipes {
				canonical, _ := CanonicalForgeRecipe(raw)
				if canonical != raw {
					return fmt.Errorf("noncanonical stored recipe in deck %d", i)
				}
			}
		}
		if deck.CreatedHeight < 0 || deck.ConsumedHeight < 0 {
			return fmt.Errorf("invalid heights")
		}
		if !deck.Consumed && (deck.ConsumedHeight != 0 || deck.AllStarId != 0) {
			return fmt.Errorf("unconsumed deck has consumption metadata")
		}
		if deck.Consumed {
			if deck.ConsumedHeight < deck.CreatedHeight || deck.AllStarId >= uint64(len(state.AllStars)) {
				return fmt.Errorf("invalid consumption provenance")
			}
			star := state.AllStars[deck.AllStarId]
			if star == nil || star.SourceDeckId != deck.Id {
				return fmt.Errorf("consumption has no matching All-Star")
			}
		}
	}
	for i, star := range state.AllStars {
		if star == nil || star.Id != uint64(i) || star.SourceDeckId >= uint64(len(state.Decks)) || star.Card == nil {
			return fmt.Errorf("invalid All-Star ID or source")
		}
		deck := state.Decks[star.SourceDeckId]
		if !deck.Consumed || deck.AllStarId != star.Id || star.Owner != deck.Owner || star.RecipeFormatVersion != deck.RecipeFormatVersion || star.CreatedHeight != deck.ConsumedHeight {
			return fmt.Errorf("inconsistent All-Star provenance")
		}
		matched := false
		for _, card := range deck.Cards {
			if card.InstanceId == star.Card.InstanceId {
				matched = proto.Equal(card, star.Card)
				break
			}
		}
		if !matched {
			return fmt.Errorf("All-Star differs from original occurrence")
		}
	}
	return nil
}

func (gs GenesisState) ValidateRunDeckReferences() error {
	if err := ValidateRunDeckGenesis(gs.RunDecks); err != nil {
		return err
	}
	if gs.RunDecks == nil {
		return nil
	}
	users := map[string]bool{}
	for i, a := range gs.Addresses {
		if i < len(gs.Users) && gs.Users[i] != nil {
			users[a] = true
		}
	}
	for _, deck := range gs.RunDecks.Decks {
		if !users[deck.Owner] || !users[deck.Reporter] {
			return fmt.Errorf("deck owner/reporter not in genesis users")
		}
		for _, card := range deck.Cards {
			if card.BaseCardId >= uint64(len(gs.CardRecords)) {
				return fmt.Errorf("deck references missing base card")
			}
			base := gs.CardRecords[card.BaseCardId]
			if base == nil || base.Status == CardStatus_none || len(base.Content) == 0 {
				return fmt.Errorf("deck references empty base card")
			}
		}
	}
	return nil
}
