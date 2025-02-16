package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) SetFinalize(goCtx context.Context, msg *types.MsgSetFinalize) (*types.MsgSetFinalizeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	setSize := int(k.GetParams(ctx).SetSize)

	set := k.sets.Get(ctx, msg.SetId)
	err := checkSetEditable(set, msg.Creator)
	if err != nil {
		return nil, err
	}

	image := k.images.Get(ctx, set.ArtworkId)

	if len(image.Image) == 0 {
		return nil, errorsmod.Wrapf(sdkerrors.ErrUnauthorized, "Set artwork is needed")
	}

	err = k.CollectSetCreationFee(ctx, msg.Creator)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInsufficientFunds, err.Error())
	}

	dist, err := k.GetRarityDistribution(ctx, *set, uint32(setSize))
	if err != nil {
		return nil, err
	}

	if dist[0] != dist[1] {
		return nil, errorsmod.Wrapf(sdkerrors.ErrUnauthorized, "Sets should contain [(common,unique,exceptional), uncommon, rare] %d, %d, %d but contains %d, %d, %d", dist[1][0], dist[1][1], dist[1][2], dist[0][0], dist[0][1], dist[0][2])
	}

	set.Status = types.CStatus_finalized
	set.ContributorsDistribution = k.GetContributorDistribution(ctx, *set)
	set.Rarities = k.GetCardRaritiesInSet(ctx, set)

	k.sets.Set(ctx, msg.SetId, set)

	return &types.MsgSetFinalizeResponse{}, nil
}

func (k Keeper) GetCardRaritiesInSet(ctx sdk.Context, set *types.Set) []*types.InnerRarities {
	var rarityNums []*types.InnerRarities = make([]*types.InnerRarities, 5)
	for idx := range rarityNums {
		rarityNums[idx] = &types.InnerRarities{}
	}
	for _, cardId := range set.Cards {
		card := k.cards.Get(ctx, cardId)
		rarityNums[int(card.Rarity)].R = append(rarityNums[int(card.Rarity)].R, cardId)
	}
	return rarityNums[:]
}
