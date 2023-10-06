package keeper

import (
	"context"
	"math/rand"

	"golang.org/x/exp/slices"

	sdkerrors "cosmossdk.io/errors"
	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
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

	var (
		cardsList     []uint64
		cleanedRatios [3]uint64
	)
	set := k.Sets.Get(ctx, creator.BoosterPacks[msg.BoosterPackId].SetId)
	rarityNums := k.getCardRaritiesInSet(ctx, set)
	for idx, ratio := range creator.BoosterPacks[msg.BoosterPackId].DropRatiosPerPack {
		if len(rarityNums[idx+2]) == 0 {
			cleanedRatios[idx] = 0
		} else {
			cleanedRatios[idx] = ratio
		}
	}

	for idx, num := range creator.BoosterPacks[msg.BoosterPackId].RaritiesPerPack {
		for i := 0; i < int(num); i++ {
			if idx != 2 {
				cardsList = append(cardsList, rarityNums[idx][rand.Intn(len(rarityNums[idx]))])
			} else {
				res := uint64(rand.Intn(int(cleanedRatios[0] + cleanedRatios[1] + cleanedRatios[2])))
				j := 4
				if res < cleanedRatios[0] {
					j = 2
				} else if res < cleanedRatios[0]+cleanedRatios[1] {
					j = 3
				}
				cardsList = append(cardsList, rarityNums[j][rand.Intn(len(rarityNums[j]))])
			}
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

func (k Keeper) getCardRaritiesInSet(ctx sdk.Context, set *types.Set) (rarityNums [5][]uint64) {
	for _, cardId := range set.Cards {
		card := k.Cards.Get(ctx, cardId)
		rarityNums[int(card.Rarity)] = append(rarityNums[int(card.Rarity)], cardId)
	}
	return
}
