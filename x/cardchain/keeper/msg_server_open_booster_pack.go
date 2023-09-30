package keeper

import (
	"context"
	"math/rand"

	"golang.org/x/exp/slices"

	sdkerrors "cosmossdk.io/errors"
	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/keywords"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) OpenBoosterPack(goCtx context.Context, msg *types.MsgOpenBoosterPack) (*types.MsgOpenBoosterPackResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	creator, err := k.GetUserFromString(ctx, msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrUserDoesNotExist, err.Error())
	}

	if uint64(len(creator.BoosterPacks)) <= msg.BoosterPackId {
		return nil, sdkerrors.Wrapf(
			types.ErrBoosterPack,
			"BoosterPackId %d is not in list length %d",
			msg.BoosterPackId,
			len(creator.BoosterPacks),
		)
	}

	var cardsList []uint64
	rarities := []string{"COMMON", "UNCOMMON", "RARE"}

	set := k.Sets.Get(ctx, creator.BoosterPacks[msg.BoosterPackId].SetId)

	for idx, num := range creator.BoosterPacks[msg.BoosterPackId].RaritiesPerPack {
		for i := 0; i < int(num); i++ {
			var rarityCards []uint64
			for _, cardId := range set.Cards {
				cardobj, err := keywords.Unmarshal(k.Cards.Get(ctx, cardId).Content)
				if err != nil {
					return nil, sdkerrors.Wrap(types.ErrCardobject, err.Error())
				}
				rarity, err := GetCardRarity(cardobj)
				if err != nil {
					return nil, sdkerrors.Wrap(types.ErrCardobject, err.Error())
				}
				if *rarity == cardobject.Rarity(rarities[idx]) {
					rarityCards = append(rarityCards, cardId)
				}
			}
			cardsList = append(cardsList, rarityCards[rand.Intn(len(rarityCards))])
		}
	}

	creator.Cards = append(creator.Cards, cardsList...)
	creator.BoosterPacks = slices.Delete(
		creator.BoosterPacks,
		int(msg.BoosterPackId),
		int(msg.BoosterPackId+1),
	)

	k.SetUserFromUser(ctx, creator)

	return &types.MsgOpenBoosterPackResponse{CardIds: cardsList}, nil
}
