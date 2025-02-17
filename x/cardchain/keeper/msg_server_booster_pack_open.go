package keeper

import (
	"context"
	"slices"

	errorsmod "cosmossdk.io/errors"
	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"golang.org/x/exp/rand"
)

func (k msgServer) BoosterPackOpen(goCtx context.Context, msg *types.MsgBoosterPackOpen) (*types.MsgBoosterPackOpenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	creator, err := k.GetUserFromString(ctx, msg.Creator)
	if err != nil {
		return nil, errorsmod.Wrap(types.ErrUserDoesNotExist, err.Error())
	}

	if uint64(len(creator.BoosterPacks)) <= msg.BoosterPackId {
		return nil, errorsmod.Wrapf(
			sdkerrors.ErrUnauthorized,
			"BoosterPackId %d is not in list length %d",
			msg.BoosterPackId,
			len(creator.BoosterPacks),
		)
	}
	rand.Seed(uint64(ctx.BlockTime().Unix()) + uint64(len(creator.BoosterPacks)))

	var (
		cardsList     []uint64
		cleanedRatios [3]uint64
	)
	pack := creator.BoosterPacks[msg.BoosterPackId]
	set := k.SetK.Get(ctx, pack.SetId)

	for idx, ratio := range pack.DropRatiosPerPack {
		if len(set.Rarities[idx+2].R) == 0 {
			cleanedRatios[idx] = 0
		} else {
			cleanedRatios[idx] = ratio
		}
	}

	for idx, num := range pack.RaritiesPerPack {
		for i := 0; i < int(num); i++ {
			if idx != 2 {
				cardsList = append(cardsList, set.Rarities[idx].R[rand.Intn(len(set.Rarities[idx].R))])
			} else {
				res := uint64(rand.Intn(int(cleanedRatios[0] + cleanedRatios[1] + cleanedRatios[2])))
				j := 4
				if res < cleanedRatios[0] {
					j = 2
				} else if res < cleanedRatios[0]+cleanedRatios[1] {
					j = 3
				}
				cardsList = append(cardsList, set.Rarities[j].R[rand.Intn(len(set.Rarities[j].R))])
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

	return &types.MsgBoosterPackOpenResponse{CardIds: cardsList}, nil
}
