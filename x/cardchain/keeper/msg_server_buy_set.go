package keeper

import (
	"context"
	"fmt"
	"strconv"

	sdkerrors "cosmossdk.io/errors"
	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) BuySet(goCtx context.Context, msg *types.MsgBuySet) (*types.MsgBuySetResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	params := k.GetParams(ctx)

	creator, err := k.GetUserFromString(ctx, msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrUserDoesNotExist, err.Error())
	}

	set := k.Sets.Get(ctx, msg.SetId)

	if set.Status != types.CStatus_active {
		return nil, types.ErrNoActiveSet
	}

	boosterPack := types.NewBoosterPack(
		ctx,
		msg.SetId,
		[]uint64{params.CommonsPerPack, params.UnCommonsPerPack, params.RaresPerPack},
	)

	// payment
	contribs := k.GetAllSetContributors(ctx, *set)
	for _, contrib := range contribs {
		contribAddr, err := sdk.AccAddressFromBech32(contrib)
		if err != nil {
			return nil, sdkerrors.Wrap(types.ErrInvalidAccAddress, "Unable to convert to AccAddress")
		}

		err = k.BankKeeper.SendCoins(ctx, creator.Addr, contribAddr, sdk.Coins{QuoCoin(params.SetPrice, int64(len(contribs)))})
		if err != nil {
			return nil, sdkerrors.Wrap(errors.ErrInsufficientFunds, err.Error())
		}
	}

	claimedAirdrop := k.ClaimAirDrop(ctx, &creator, types.AirDrop_buy)
	creator.BoosterPacks = append(creator.BoosterPacks, &boosterPack)

	k.SetUserFromUser(ctx, creator)

	inflationRate, err := strconv.ParseFloat(params.InflationRate, 8)
	pPool := k.Pools.Get(ctx, PublicPoolKey)
	newPool := MulCoinFloat(*pPool, inflationRate)
	k.Pools.Set(ctx, PublicPoolKey, &newPool)
	k.Logger(ctx).Info(fmt.Sprintf(":: PublicPool: %s", newPool))

	return &types.MsgBuySetResponse{AirdropClaimed: claimedAirdrop}, nil
}
