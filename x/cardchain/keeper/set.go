package keeper

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/types/errors"

	sdkerrors "cosmossdk.io/errors"
	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/keywords"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// CollectSetFee Collects a fee from a user
func (k Keeper) CollectSetFee(ctx sdk.Context, price sdk.Coin, creator string) error {
	err := k.BurnCoinsFromString(ctx, creator, sdk.Coins{price})
	if err != nil {
		return err
	}
	k.AddPoolCredits(ctx, PublicPoolKey, price)
	return nil
}

// CollectSetConributionFee Is a wrapper for CollectSetFee with contributionfee
func (k Keeper) CollectSetConributionFee(ctx sdk.Context, creator string) error {
	return k.CollectSetFee(ctx, sdk.NewInt64Coin("ucredits", 1000000), creator)
}

// CollectSetCreationFee Is a wrapper for CollectSetFee with creationfee
func (k Keeper) CollectSetCreationFee(ctx sdk.Context, creator string) error {
	return k.CollectSetFee(ctx, k.GetParams(ctx).SetCreationFee, creator)
}

// GetAllSetContributors Returns an array of all contributors of a set in their respective frequencies
func (k Keeper) GetAllSetContributors(ctx sdk.Context, set types.Set) []string {
	contribs := []string{set.StoryWriter, set.StoryWriter, set.Artist, set.Artist, set.Contributors[0], set.Contributors[0], set.Contributors[0], set.Contributors[0]}
	for _, cardId := range set.Cards {
		var card = k.Cards.Get(ctx, cardId)
		if card.Owner != "" {
			contribs = append(contribs, card.Owner, card.Artist)
		}
	}
	return contribs
}

// GetActiveSets Return a list of all active sets ids
func (k Keeper) GetActiveSets(ctx sdk.Context) (activeSets []uint64) {
	iter := k.Sets.GetItemIterator(ctx)
	for ; iter.Valid(); iter.Next() {
		idx, set := iter.Value()
		if set.Status == types.CStatus_active {
			activeSets = append(activeSets, idx)
		}
	}
	return
}

func (k Keeper) GetRarityDistribution(ctx sdk.Context, set types.Set, setSize uint64) (dist [2][4]uint64, err error) {
	var (
		unCommons, rares, commons, others, commonsAll, unCommonsAll, raresAll uint64
	)

	unCommonsAll = uint64(setSize / 3)
	raresAll = uint64(setSize / 3)
	commonsAll = uint64(setSize - raresAll - unCommonsAll)

	for _, cardId := range set.Cards {
		cardobj, err := keywords.Unmarshal(k.Cards.Get(ctx, cardId).Content)
		if err != nil {
			return dist, sdkerrors.Wrap(types.ErrCardobject, err.Error())
		}
		rarity, err := GetCardRarity(cardobj)
		if err != nil {
			return dist, sdkerrors.Wrap(types.ErrCardobject, err.Error())
		}
		if rarity == nil {
			others++
		} else {
			switch *rarity {
			case cardobject.Rarity("COMMON"):
				commons++
			case cardobject.Rarity("UNCOMMON"):
				unCommons++
			case cardobject.Rarity("RARE"):
				rares++
			default:
				others++
			}
		}
	}

	return [2][4]uint64{
		[4]uint64{commons, unCommons, rares, others},
		[4]uint64{commonsAll, unCommonsAll, raresAll, 0},
	}, nil
}

func GetCardRarity(card *keywords.Card) (*cardobject.Rarity, error) {
	if card.Action != nil {
		return card.Action.Rarity, nil
	} else if card.Place != nil {
		return card.Place.Rarity, nil
	} else if card.Entity != nil {
		return card.Entity.Rarity, nil
	} else if card.Headquarter != nil {
		return card.Headquarter.Rarity, nil
	}
	return nil, fmt.Errorf("no card-attributes")
}

func checkSetEditable(set *types.Set, user string) error {
	if len(set.Contributors) == 0 {
		return sdkerrors.Wrap(types.ErrUninitializedType, "Set not initialized")
	}

	if user != set.Contributors[0] {
		return sdkerrors.Wrap(errors.ErrUnauthorized, "Invalid creator")
	}

	if set.Status != types.CStatus_design {
		return types.ErrSetNotInDesign
	}
	return nil
}
