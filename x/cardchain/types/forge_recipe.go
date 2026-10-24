package types

import (
	"encoding/json"
	"fmt"
	"io"
	"slices"
	"sort"
	"strings"
	"unicode/utf8"
)

const (
	ForgeRecipeVersion     uint32 = 1
	MaxRunDeckCards               = 128
	MaxForgeRecipesPerCard        = 64
	MaxForgeRecipeBytes           = 16 * 1024
	MaxRunDeckRecipeBytes         = 256 * 1024
	MaxForgeEffects               = 64
)

// Format 1 follows gameclient b21ddc48's durable recipe codec, not the current
// base-card grammar. In particular it preserves legacy effect adapters.
var forgeEffectFields = map[string]string{
	"Anthem": "Tag", "Arm": "Target Amount", "Armor": "Target Amount", "Bounce": "Target", "Burn": "Target Amount",
	"Choose": "Type Zone", "Count": "Tag", "CountDust": "CardType", "CountForces": "Power", "CountPower": "Target",
	"Discard": "", "DiscountAction": "Amount", "DiscountEntity": "Amount", "DiscountPlace": "Amount",
	"DrawAction": "Tag", "DrawEntity": "Tag", "DrawPlace": "Tag", "Enrage": "Target", "Fortify": "Amount",
	"Grind": "Target Amount", "Grow": "GrowthAmount", "Harm": "Target Amount", "Hasten": "Target", "Heal": "Target",
	"Insight": "WisdomAmount", "Kill": "Target", "Laneswap": "Target", "Mill": "Amount Player", "Produce": "ManaAmount",
	"Reassemble": "Amount", "RecoverAction": "", "RecoverEntity": "", "RecoverPlace": "", "Repair": "Target Amount",
	"Resurrect": "Amount", "Sacrifice": "Target", "Silence": "Target Player", "Spawn": "Token Amount",
	"Strengthen": "Target Amount", "Suffer": "Target Amount", "Train": "Target Amount", "Void": "Target", "Withdraw": "Target",
}
var forgeAbilityFields = map[string]string{
	"Arrival": "Effects", "Avenge": "Effects", "Battlecry": "Effects", "Channel": "Effects", "DiscardPay": "ManaAmount Effects",
	"Dismantle": "Target Effects", "Dissolve": "VoidAmount Effects", "Furious": "Effects", "Loot": "Effects", "OnConstruction": "Effects",
	"OnDeath": "Effects", "OnSpawn": "Effects", "Pay": "ManaAmount Effects", "Periodic": "Effects", "Tribute": "Target Effects",
}
var forgeEnums = map[string]string{
	"Target": "DEFAULT ALL THIS RANDOM TARGET TOKEN CHOSENACTION CHOSENENTITY CHOSENPLACE",
	"Tag":    "DEFAULT AIR ALCHEMIST ANIMAL ARTIFACT ASSASSIN BEAST BELIEVER BOT BUILDING COLOSSUS DISASTER DOMESTIC DRAGON EARTH EPIC EVENT FACTORY FIRE HUMAN KNOWLEDGE LANDSCAPE PLANT REINFORCEMENT WIZARD SHRINE SPIRIT TACTIC TECHNOCRAT UNDEAD VEHICLE WARRIOR WATER WEAPON WORKER SPAWNED GOBLIN ELF",
	"Type":   "ENTITY ACTION PLACE HEADQUARTER DEFAULT", "CardType": "ENTITY ACTION PLACE HEADQUARTER DEFAULT",
	"Zone": "ATTACK PLACE QUERY DUSTPILE HAND DECK VOID HQ TOKEN FIELD", "Player": "YOU OPPONENT", "IntVariable": "X Y Z",
}

func forgeEnum(value any, field string) error {
	s, ok := value.(string)
	if !ok || !slices.Contains(strings.Fields(forgeEnums[field]), s) || s == "" {
		return fmt.Errorf("invalid %s", field)
	}
	return nil
}
func forgeNumber(value any, max int64) error {
	n, ok := value.(json.Number)
	if !ok {
		return fmt.Errorf("expected integer")
	}
	i, err := n.Int64()
	if err != nil || i < 0 || i > max {
		return fmt.Errorf("integer outside 0..%d", max)
	}
	return nil
}
func forgeObject(value any, allowed string) (map[string]any, error) {
	m, ok := value.(map[string]any)
	if !ok {
		return nil, fmt.Errorf("expected object")
	}
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		if !slices.Contains(strings.Fields(allowed), k) || k == "" {
			return nil, fmt.Errorf("unknown field %q", k)
		}
	}
	return m, nil
}
func forgeIntValue(value any) error {
	m, err := forgeObject(value, "SimpleIntValue IntVariable")
	if err != nil {
		return err
	}
	if variable, ok := m["IntVariable"]; ok {
		if err := forgeEnum(variable, "IntVariable"); err != nil {
			return err
		}
		if n, ok := m["SimpleIntValue"]; ok {
			return forgeNumber(n, 0)
		}
		return nil
	}
	return forgeNumber(m["SimpleIntValue"], 255)
}
func forgeAlternative(value any, definitions map[string]string) (string, map[string]any, error) {
	m, ok := value.(map[string]any)
	if !ok || len(m) != 1 {
		return "", nil, fmt.Errorf("expected exactly one alternative")
	}
	for name, payload := range m {
		fields, ok := definitions[name]
		if !ok {
			return "", nil, fmt.Errorf("unsupported alternative %q", name)
		}
		obj, err := forgeObject(payload, fields)
		if err != nil {
			return "", nil, err
		}
		for _, f := range strings.Fields(fields) {
			if obj[f] == nil {
				return "", nil, fmt.Errorf("missing %s.%s", name, f)
			}
		}
		return name, obj, nil
	}
	panic("unreachable")
}
func validateForgeAlternative(value any, ability bool) (string, map[string]any, error) {
	defs := forgeEffectFields
	if ability {
		defs = forgeAbilityFields
	}
	name, m, err := forgeAlternative(value, defs)
	if err != nil {
		return "", nil, err
	}
	for _, f := range strings.Fields(defs[name]) {
		v := m[f]
		switch f {
		case "Effects":
			_, err = validateForgeEffects(v)
		case "Target", "Tag", "Type", "CardType", "Zone", "Player":
			err = forgeEnum(v, f)
		case "Power", "VoidAmount":
			err = forgeNumber(v, 255)
		case "Token":
			_, _, err = forgeAlternative(v, map[string]string{"Beast": "", "Bot": "", "Recruit": "", "Powerstone": "", "Avatar": "", "Spirit": ""})
		default:
			if name == "DiscardPay" && f == "ManaAmount" {
				err = forgeNumber(v, 255)
			} else {
				err = forgeIntValue(v)
			}
		}
		if err != nil {
			return "", nil, fmt.Errorf("%s.%s: %w", name, f, err)
		}
	}
	return name, m, nil
}
func validateForgeEffects(value any) ([]any, error) {
	effects, ok := value.([]any)
	if !ok || len(effects) > MaxForgeEffects {
		return nil, fmt.Errorf("invalid effects array")
	}
	for i, e := range effects {
		if _, _, err := validateForgeAlternative(e, false); err != nil {
			return nil, fmt.Errorf("effect %d: %w", i, err)
		}
	}
	return effects, nil
}

// Token-level parsing rejects duplicate members before encoding/json can drop
// them. UseNumber preserves exact integers; depth and bytes are bounded.
func readForgeJSON(d *json.Decoder, depth int) (any, error) {
	if depth > 32 {
		return nil, fmt.Errorf("recipe JSON too deep")
	}
	t, err := d.Token()
	if err != nil {
		return nil, err
	}
	delim, ok := t.(json.Delim)
	if !ok {
		return t, nil
	}
	switch delim {
	case '{':
		m := map[string]any{}
		seen := map[string]bool{}
		for d.More() {
			key, err := d.Token()
			if err != nil {
				return nil, err
			}
			k, ok := key.(string)
			if !ok {
				return nil, fmt.Errorf("invalid member")
			}
			lower := strings.ToLower(k)
			if seen[lower] {
				return nil, fmt.Errorf("duplicate member %q", k)
			}
			seen[lower] = true
			v, err := readForgeJSON(d, depth+1)
			if err != nil {
				return nil, err
			}
			m[k] = v
		}
		if _, err := d.Token(); err != nil {
			return nil, err
		}
		return m, nil
	case '[':
		arr := []any{}
		for d.More() {
			v, err := readForgeJSON(d, depth+1)
			if err != nil {
				return nil, err
			}
			arr = append(arr, v)
		}
		if _, err := d.Token(); err != nil {
			return nil, err
		}
		return arr, nil
	default:
		return nil, fmt.Errorf("unexpected delimiter")
	}
}

// CanonicalForgeRecipe validates a durable recipe and sorts object keys only.
// Array order, duplicates in arrays and gameplay values are never normalized away.
func CanonicalForgeRecipe(raw string) (string, error) {
	if len(raw) == 0 || len(raw) > MaxForgeRecipeBytes || !utf8.ValidString(raw) {
		return "", fmt.Errorf("invalid recipe size or UTF-8")
	}
	d := json.NewDecoder(strings.NewReader(raw))
	d.UseNumber()
	value, err := readForgeJSON(d, 0)
	if err != nil {
		return "", err
	}
	if _, err = d.Token(); err != io.EOF {
		return "", fmt.Errorf("trailing JSON")
	}
	root, err := forgeObject(value, "FixedPackageId Generated")
	if err != nil {
		return "", err
	}
	fixed, generated := root["FixedPackageId"], root["Generated"]
	if (fixed == nil) == (generated == nil) {
		return "", fmt.Errorf("recipe requires fixed OR generated")
	}
	if fixed != nil {
		id, ok := fixed.(string)
		if !ok || id == "" || len(id) > 128 || strings.TrimSpace(id) != id {
			return "", fmt.Errorf("invalid fixed package ID")
		}
	} else if err = validateGeneratedForge(generated); err != nil {
		return "", err
	}
	canonical, err := json.Marshal(value)
	if err != nil {
		return "", err
	}
	if len(canonical) > MaxForgeRecipeBytes {
		return "", fmt.Errorf("canonical recipe too large")
	}
	return string(canonical), nil
}
func validateGeneratedForge(value any) error {
	g, err := forgeObject(value, "Ability Effects AdditionalCost EffectParameters")
	if err != nil {
		return err
	}
	if (g["Ability"] == nil) == (g["Effects"] == nil) {
		return fmt.Errorf("generated recipe requires ability OR effects")
	}
	var effects []any
	if g["Ability"] != nil {
		_, a, e := validateForgeAlternative(g["Ability"], true)
		if e != nil {
			return e
		}
		effects = a["Effects"].([]any)
	} else {
		effects, err = validateForgeEffects(g["Effects"])
		if err != nil {
			return err
		}
	}
	if cost := g["AdditionalCost"]; cost != nil {
		_, c, e := forgeAlternative(cost, map[string]string{"DiscardCost": "Amount", "SacrificeCost": "Amount", "VoidCost": "Amount"})
		if e != nil {
			return e
		}
		if e = forgeNumber(c["Amount"], 255); e != nil {
			return e
		}
	}
	if g["EffectParameters"] == nil {
		return nil
	}
	adapters, ok := g["EffectParameters"].([]any)
	if !ok || len(adapters) > len(effects) {
		return fmt.Errorf("invalid adapters")
	}
	seen := map[int64]bool{}
	for _, value := range adapters {
		a, err := forgeObject(value, "EffectIndex Amount Zones")
		if err != nil {
			return err
		}
		if err = forgeNumber(a["EffectIndex"], int64(len(effects))-1); err != nil {
			return err
		}
		idx, _ := a["EffectIndex"].(json.Number).Int64()
		if seen[idx] {
			return fmt.Errorf("duplicate effect index")
		}
		seen[idx] = true
		name, payload, err := validateForgeAlternative(effects[idx], false)
		if err != nil {
			return err
		}
		if a["Amount"] == nil && a["Zones"] == nil {
			return fmt.Errorf("empty adapter")
		}
		if amount := a["Amount"]; amount != nil {
			if name != "Discard" && name != "Void" && name != "Sacrifice" {
				return fmt.Errorf("amount adapter incompatible with %s", name)
			}
			if err = forgeNumber(amount, 255); err != nil {
				return err
			}
		}
		if zones := a["Zones"]; zones != nil {
			z, ok := zones.([]any)
			if name != "Choose" || !ok || len(z) == 0 || len(z) > 16 {
				return fmt.Errorf("invalid zones adapter")
			}
			for _, zone := range z {
				if err = forgeEnum(zone, "Zone"); err != nil {
					return err
				}
			}
			if z[0] != payload["Zone"] {
				return fmt.Errorf("first zone differs from Choose zone")
			}
		}
	}
	return nil
}

func ValidateRunDeckCards(version uint32, cards []*RunDeckCard) error {
	if version != ForgeRecipeVersion {
		return fmt.Errorf("unsupported recipe format %d", version)
	}
	if len(cards) == 0 || len(cards) > MaxRunDeckCards {
		return fmt.Errorf("deck requires 1..%d cards", MaxRunDeckCards)
	}
	seen := map[uint64]bool{}
	total, canonicalTotal := 0, 0
	for _, card := range cards {
		if card == nil {
			return fmt.Errorf("null card")
		}
		if seen[card.InstanceId] {
			return fmt.Errorf("duplicate instance %d", card.InstanceId)
		}
		seen[card.InstanceId] = true
		if len(card.ForgeRecipes) > MaxForgeRecipesPerCard {
			return fmt.Errorf("too many upgrades")
		}
		for i, raw := range card.ForgeRecipes {
			total += len(raw)
			if total > MaxRunDeckRecipeBytes {
				return fmt.Errorf("deck recipes too large")
			}
			canonical, err := CanonicalForgeRecipe(raw)
			if err != nil {
				return fmt.Errorf("instance %d recipe %d: %w", card.InstanceId, i, err)
			}
			canonicalTotal += len(canonical)
			if canonicalTotal > MaxRunDeckRecipeBytes {
				return fmt.Errorf("canonical deck recipes too large")
			}
		}
	}
	return nil
}
