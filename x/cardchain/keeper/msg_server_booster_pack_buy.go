package keeper

import (
	"context"
	"fmt"
	"strconv"

	errorsmod "cosmossdk.io/errors"
	"github.com/DecentralCardGame/cardchain/util"
	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) BoosterPackBuy(goCtx context.Context, msg *types.MsgBoosterPackBuy) (*types.MsgBoosterPackBuyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	params := k.GetParams(ctx)

	creator, err := k.GetUserFromString(ctx, msg.Creator)
	if err != nil {
		return nil, errorsmod.Wrap(types.ErrUserDoesNotExist, err.Error())
	}

	set := k.SetK.Get(ctx, msg.SetId)

	if set.Status != types.SetStatus_active {
		return nil, errorsmod.Wrapf(sdkerrors.ErrUnauthorized, "invalid set status: %s", set.Status)
	}

	boosterPack := types.NewBoosterPack(
		ctx,
		msg.SetId,
		[]uint64{params.CommonsPerPack, params.UnCommonsPerPack, params.RaresPerPack},
		[]uint64{params.RareDropRatio, params.ExceptionalDropRatio, params.UniqueDropRatio},
	)

	// payment
	for _, contrib := range set.ContributorsDistribution {
		contribAddr, err := sdk.AccAddressFromBech32(contrib.Addr)
		if err != nil {
			return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "Unable to convert to AccAddress")
		}

		err = k.BankKeeper.SendCoins(ctx, creator.Addr, contribAddr, sdk.Coins{*contrib.Payment})
		if err != nil {
			return nil, errorsmod.Wrap(errors.ErrInsufficientFunds, err.Error())
		}
	}

	claimedAirdrop := k.ClaimAirDrop(ctx, &creator, types.AirDrop_buy)
	creator.BoosterPacks = append(creator.BoosterPacks, &boosterPack)

	k.SetUserFromUser(ctx, creator)

	inflationRate, err := strconv.ParseFloat(params.InflationRate, 8)
	pPool := k.Pools.Get(ctx, PublicPoolKey)
	newPool := util.MulCoinFloat(*pPool, inflationRate)
	k.Pools.Set(ctx, PublicPoolKey, &newPool)
	k.Logger().Info(fmt.Sprintf(":: PublicPool: %s", newPool))

	return &types.MsgBoosterPackBuyResponse{AirdropClaimed: claimedAirdrop}, nil
}
