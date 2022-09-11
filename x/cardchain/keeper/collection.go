package keeper

import (
	"errors"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/keywords"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// CollectCollectionFee Collects a fee from a user
func (k Keeper) CollectCollectionFee(ctx sdk.Context, price sdk.Coin, creator string) error {
	err := k.BurnCoinsFromString(ctx, creator, sdk.Coins{price})
	if err != nil {
		return err
	}
	k.AddPoolCredits(ctx, PublicPoolKey, price)
	return nil
}

// CollectCollectionConributionFee Is a wrapper for CollectCollectionFee with contributionfee
func (k Keeper) CollectCollectionConributionFee(ctx sdk.Context, creator string) error {
	return k.CollectCollectionFee(ctx, sdk.NewInt64Coin("ucredits", 1000000), creator)
}

// CollectCollectionCreationFee Is a wrapper for CollectCollectionFee with creationfee
func (k Keeper) CollectCollectionCreationFee(ctx sdk.Context, creator string) error {
	return k.CollectCollectionFee(ctx, k.GetParams(ctx).CollectionCreationFee, creator)
}

// GetAllCollectionContributors Returns an array of all contributors of a collection in their respective frequencies
func (k Keeper) GetAllCollectionContributors(ctx sdk.Context, collection types.Collection) []string {
	contribs := []string{collection.StoryWriter, collection.StoryWriter, collection.Artist, collection.Artist, collection.Contributors[0], collection.Contributors[0], collection.Contributors[0], collection.Contributors[0]}
	for _, cardId := range collection.Cards {
		var card = k.Cards.Get(ctx, cardId)
		if card.Owner != "" {
			contribs = append(contribs, card.Owner, card.Artist)
		}
	}
	return contribs
}

// GetActiveCollections Return a list of all active collections ids
func (k Keeper) GetActiveCollections(ctx sdk.Context) (activeCollections []uint64) {
	iter := k.Collections.GetItemIterator(ctx)
	for ; iter.Valid(); iter.Next() {
		idx, collection := iter.Value()
		if collection.Status == types.CStatus_active {
			activeCollections = append(activeCollections, idx)
		}
	}
	return
}

func (k Keeper) GetRarityDistribution(ctx sdk.Context, collection types.Collection, collectionSize uint64) (dist [2][4]uint64, err error) {
	var (
		unCommons, rares, commons, others, commonsAll, unCommonsAll, raresAll uint64
	)

	unCommonsAll = uint64(collectionSize / 3)
	raresAll = uint64(collectionSize / 3)
	commonsAll = uint64(collectionSize - raresAll - unCommonsAll)

	for _, cardId := range collection.Cards {
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
	return nil, errors.New("No card-attributes")
}
