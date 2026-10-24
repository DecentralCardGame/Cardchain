package types

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestForgeRecipeCanonicalization(t *testing.T) {
	valid := []string{
		`{"FixedPackageId":"fixed.fire.1"}`,
		`{"Generated":{"Effects":[{"Sacrifice":{"Target":"THIS"}}],"EffectParameters":[{"EffectIndex":0,"Amount":2}]}}`,
		`{"Generated":{"Effects":[{"Choose":{"Type":"ENTITY","Zone":"FIELD"}}],"EffectParameters":[{"EffectIndex":0,"Zones":["FIELD","DUSTPILE"]}]}}`,
		`{"Generated":{"Ability":{"OnSpawn":{"Effects":[{"Burn":{"Amount":{"SimpleIntValue":0,"IntVariable":"X"},"Target":"RANDOM"}}]}},"AdditionalCost":{"VoidCost":{"Amount":1}}}}`,
	}
	for _, raw := range valid {
		canonical, err := CanonicalForgeRecipe(raw)
		require.NoError(t, err, raw)
		again, err := CanonicalForgeRecipe(canonical)
		require.NoError(t, err)
		require.Equal(t, canonical, again)
	}
	a, err := CanonicalForgeRecipe(`{"Generated":{"Effects":[{"Harm":{"Target":"ALL","Amount":{"SimpleIntValue":3}}}]}}`)
	require.NoError(t, err)
	b, err := CanonicalForgeRecipe(`{ "Generated": { "Effects": [ { "Harm": { "Amount": { "SimpleIntValue": 3 }, "Target": "ALL" } } ] } }`)
	require.NoError(t, err)
	require.Equal(t, a, b)
}

func TestForgeRecipeRejectsAmbiguityAndLoss(t *testing.T) {
	bad := []string{
		`{}`, `null`, `[]`, `{"FixedPackageId":""}`, `{"FixedPackageId":" a"}`, `{"FixedPackageId":"a","Generated":{"Effects":[]}}`,
		`{"FixedPackageId":"a","FixedPackageId":"b"}`, `{"FixedPackageId":"a","fixedpackageid":"b"}`, `{"FixedPackageId":"a","Price":3}`, `{"FixedPackageId":"a"} {}`,
		`{"Generated":{"Effects":[{"Sacrifice":{"Target":"THIS RANDOM"}}]}}`,
		`{"FixedPackageId":"a","FixedPackageId Generated":null}`,
		`{"Generated":{"Effects":[{"Unknown":{}}]}}`, `{"Generated":{"Effects":[{"Sacrifice":{}}]}}`,
		`{"Generated":{"Effects":[{"Harm":{"Target":"ALL","Amount":{"SimpleIntValue":256}}}]}}`,
		`{"Generated":{"Effects":[{"Harm":{"Target":"ALL","Amount":{"SimpleIntValue":1.5}}}]}}`,
		`{"Generated":{"Effects":[{"Harm":{"Target":"ALL","Amount":{"SimpleIntValue":1,"IntVariable":"X"}}}]}}`,
		`{"Generated":{"Effects":[{"Sacrifice":{"Target":"THIS","Ignored":2}}]}}`,
		`{"Generated":{"Effects":[{"Harm":{"Target":"ALL","Amount":{"SimpleIntValue":3}}}],"EffectParameters":[{"EffectIndex":0,"Amount":1}]}}`,
		`{"Generated":{"Effects":[{"Sacrifice":{"Target":"THIS"}}],"EffectParameters":[{"Amount":1}]}}`,
		`{"Generated":{"Effects":[{"Sacrifice":{"Target":"THIS"}}],"EffectParameters":[{"EffectIndex":1,"Amount":1}]}}`,
		`{"Generated":{"Effects":[{"Sacrifice":{"Target":"THIS"}}],"EffectParameters":[{"EffectIndex":0,"Amount":1},{"EffectIndex":0,"Amount":2}]}}`,
		`{"Generated":{"Effects":[{"Choose":{"Type":"ENTITY","Zone":"FIELD"}}],"EffectParameters":[{"EffectIndex":0,"Zones":["DECK"]}]}}`,
		strings.Repeat("[", 40) + strings.Repeat("]", 40), strings.Repeat("x", MaxForgeRecipeBytes+1),
	}
	for _, raw := range bad {
		_, err := CanonicalForgeRecipe(raw)
		require.Error(t, err, raw)
	}
}

func TestCanonicalRecipeAggregateLimit(t *testing.T) {
	raw := `{"FixedPackageId":"` + strings.Repeat("<", 128) + `"}`
	cards := make([]*RunDeckCard, 8)
	for i := range cards {
		cards[i] = &RunDeckCard{InstanceId: uint64(i), ForgeRecipes: make([]string, 64)}
		for j := range cards[i].ForgeRecipes {
			cards[i].ForgeRecipes[j] = raw
		}
	}
	require.Less(t, len(raw)*8*64, MaxRunDeckRecipeBytes)
	require.ErrorContains(t, ValidateRunDeckCards(1, cards), "canonical deck recipes too large")
}

func TestRunDeckLimitsAndOccurrenceIdentity(t *testing.T) {
	card := &RunDeckCard{InstanceId: 0, BaseCardId: 8, ForgeRecipes: []string{`{"FixedPackageId":"a"}`, `{"FixedPackageId":"a"}`}}
	require.NoError(t, ValidateRunDeckCards(1, []*RunDeckCard{card, {InstanceId: 1, BaseCardId: 8}}))
	require.Error(t, ValidateRunDeckCards(2, []*RunDeckCard{card}))
	require.Error(t, ValidateRunDeckCards(1, nil))
	require.Error(t, ValidateRunDeckCards(1, []*RunDeckCard{nil}))
	require.Error(t, ValidateRunDeckCards(1, []*RunDeckCard{card, card}))
	require.Error(t, ValidateRunDeckCards(1, make([]*RunDeckCard, MaxRunDeckCards+1)))
	card.ForgeRecipes = make([]string, MaxForgeRecipesPerCard+1)
	require.Error(t, ValidateRunDeckCards(1, []*RunDeckCard{card}))
}
