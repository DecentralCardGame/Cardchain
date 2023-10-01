package keeper

import (
	"github.com/cosmos/cosmos-sdk/types/errors"

	sdkerrors "cosmossdk.io/errors"
	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
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

func (k Keeper) GetRarityDistribution(ctx sdk.Context, set types.Set, setSize uint32) (dist [2][3]uint32, err error) {
	var (
		unCommons, rares, commons, commonsAll, unCommonsAll, raresAll uint32
	)

	unCommonsAll = uint32(setSize / 3)
	raresAll = uint32(setSize / 3)
	commonsAll = uint32(setSize - raresAll - unCommonsAll)

	for _, cardId := range set.Cards {
		card := k.Cards.Get(ctx, cardId)
		if err != nil {
			return dist, sdkerrors.Wrap(types.ErrCardobject, err.Error())
		}
		switch card.Rarity {
		case types.CardRarity_common, types.CardRarity_unique, types.CardRarity_exceptional:
			commons++
		case types.CardRarity_uncommon:
			unCommons++
		case types.CardRarity_rare:
			rares++
		default:
			return dist, sdkerrors.Wrapf(errors.ErrInvalidType, "Invalid rarity (%d) for card (%d)", card.Rarity, cardId)
		}
	}

	return [2][3]uint32{
		{commons, unCommons, rares},
		{commonsAll, unCommonsAll, raresAll},
	}, nil
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
