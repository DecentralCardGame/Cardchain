package keeper

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/keywords"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) BuyCollection(goCtx context.Context, msg *types.MsgBuyCollection) (*types.MsgBuyCollectionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	params := k.GetParams(ctx)

	creator, err := k.GetUserFromString(ctx, msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrUserDoesNotExist, err.Error())
	}

	collection := k.Collections.Get(ctx, msg.CollectionId)
	if collection.Status != types.CStatus_active {
		return nil, types.ErrNoActiveCollection
	}

	cardsList := []uint64{}
	raritiesPerPack := []uint64{params.CommonsPerPack, params.UnCommonsPerPack, params.RaresPerPack}
	rarities := []string{"COMMON", "UNCOMMON", "RARES"}
	for idx, num := range raritiesPerPack {
		for i := 0; i < int(num); i++ {
			var rarityCards []uint64
			for _, cardId := range collection.Cards {
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

	// payment
	contribs := k.GetAllCollectionContributors(ctx, *collection)
	for _, contrib := range contribs {
		contribAddr, err := sdk.AccAddressFromBech32(contrib)
		if err != nil {
			return nil, sdkerrors.Wrap(types.ErrInvalidAccAddress, "Unable to convert to AccAddress")
		}

		err = k.BankKeeper.SendCoins(ctx, creator.Addr, contribAddr, sdk.Coins{QuoCoin(params.CollectionPrice, int64(len(contribs)))})
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, err.Error())
		}
	}

	creator.Cards = append(creator.Cards, cardsList...)
	claimedAirdrop := k.ClaimAirDrop(ctx, &creator, types.AirDrop_buy)

	k.SetUserFromUser(ctx, creator)

	inflationRate, err := strconv.ParseFloat(params.InflationRate, 8)
	pPool := k.Pools.Get(ctx, PublicPoolKey)
	newPool := MulCoinFloat(*pPool, inflationRate)
	k.Pools.Set(ctx, PublicPoolKey, &newPool)
	k.Logger(ctx).Info(fmt.Sprintf(":: PublicPool: %s", newPool))

	return &types.MsgBuyCollectionResponse{claimedAirdrop}, nil
}
