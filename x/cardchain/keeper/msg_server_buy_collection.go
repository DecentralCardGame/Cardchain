package keeper

import (
	"context"
	"math/rand"
	"fmt"
	"strconv"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) BuyCollection(goCtx context.Context, msg *types.MsgBuyCollection) (*types.MsgBuyCollectionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	creator, err := k.GetUserFromString(ctx, msg.Creator)
	if err != nil {
		return nil, err
	}

	collection := k.GetCollection(ctx, msg.CollectionId)

	if collection.Status != types.CStatus_active {
		return nil, types.ErrNoActiveCollection
	}

	cardsList := []uint64{}
	for i := 0; i < 13; i++ {
		cardsList = append(cardsList, collection.Cards[rand.Intn(len(collection.Cards))])
	}

	contribs := k.GetAllCollectionContributors(ctx, collection)

	for _, contrib := range contribs {
		contribAddr, err := sdk.AccAddressFromBech32(contrib)
		if err != nil {
			return nil, sdkerrors.Wrap(types.ErrInvalidAccAddress, "Unable to convert to AccAddress")
		}

		err = k.BankKeeper.SendCoins(ctx, creator.Addr, contribAddr, sdk.Coins{sdk.NewCoin("ucredits", k.GetParams(ctx).CollectionPrice.Amount.Quo(sdk.NewInt(int64(len(contribs)))))})
		if err != nil {
			return nil, err
		}
	}

	creator.Cards = append(creator.Cards, cardsList...)

	k.SetUserFromUser(ctx, creator)

	inflationRate, err := strconv.ParseFloat(k.GetParams(ctx).InflationRate, 8)
	pPool := k.GetPool(ctx, PublicPoolKey)
	pPool = sdk.NewInt64Coin("ucredits", int64(float64(pPool.Amount.Int64())*inflationRate))
	k.SetPool(ctx, PublicPoolKey, pPool)
	k.Logger(ctx).Info(fmt.Sprintf(":: PublicPool: %s", pPool))

	return &types.MsgBuyCollectionResponse{}, nil
}
