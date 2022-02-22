package keeper

import (
	"context"
	"errors"
	"fmt"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/keywords"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) FinalizeCollection(goCtx context.Context, msg *types.MsgFinalizeCollection) (*types.MsgFinalizeCollectionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	collectionSize := int(k.GetParams(ctx).CollectionSize)

	collection := k.GetCollection(ctx, msg.CollectionId)
	if msg.Creator != collection.Contributors[0] {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Invalid creator")
	}

	if collection.Status != types.CStatus_design {
		return nil, types.ErrCollectionNotInDesign
	}

	if len(collection.Cards) != collectionSize {
		return nil, sdkerrors.Wrapf(types.ErrCollectionSize, "Has to be %d", collectionSize)
	}

	err := k.CollectCollectionCreationFee(ctx, msg.Creator)
	if err != nil {
		return nil, err
	}

	unCommonsAll := int(collectionSize / 3)
	raresAll := int(collectionSize / 3)
	commonsAll := collectionSize - raresAll - unCommonsAll

	var (
		unCommons, rares, commons int
	)

	for _, cardId := range collection.Cards {
		cardobj, err := keywords.Unmarshal(k.GetCard(ctx, cardId).Content)
		if err != nil {
			return nil, err
		}
		rarity, err := GetCardRarity(cardobj)
		if err != nil {
			return nil, err
		}
		switch *rarity {
		case cardobject.Rarity("COMMON"):
			commons++
		case cardobject.Rarity("UNCOMMON"):
			unCommons++
		case cardobject.Rarity("RARE"):
			rares++
		default:
			return nil, errors.New(fmt.Sprintf("Card '%d' has no type", cardId))
		}
	}

	if unCommons != unCommonsAll || commons != commonsAll || rares != raresAll {
		return nil, errors.New(fmt.Sprintf("Collections should contain (c,u,r) %d, %d, %d but contains %d, %d, %d", commonsAll, unCommonsAll, raresAll, commons, unCommons, rares))
	}

	collection.Status = types.CStatus_finalized

	k.SetCollection(ctx, msg.CollectionId, collection)

	return &types.MsgFinalizeCollectionResponse{}, nil
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
