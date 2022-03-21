package keeper

import (
	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
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
	for idx, collection := range k.Collections.GetAll(ctx) {
		if collection.Status == types.CStatus_active {
			activeCollections = append(activeCollections, uint64(idx))
		}
	}
	return
}
