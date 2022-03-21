package keeper

import (
	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// SetMatchReporter Makes a user a match reporter
func (k Keeper) SetMatchReporter(ctx sdk.Context, address string) error {
	reporter, err := k.GetUserFromString(ctx, address)
	if err != nil {
		return err
	}
	reporter.ReportMatches = true

	k.SetUserFromUser(ctx, reporter)
	return nil
}

// CalculateMatchReward Calculates the match winning rewards
func (k Keeper) CalculateMatchReward(ctx sdk.Context, outcome types.Outcome) (amountA sdk.Coin, amountB sdk.Coin) {
	reward := k.GetMatchReward(ctx)
	amountA = sdk.NewInt64Coin("ucredits", 0)
	amountB = sdk.NewInt64Coin("ucredits", 0)

	if outcome == types.Outcome_AWon {
		amountA = reward
	} else if outcome == types.Outcome_BWon {
		amountB = reward
	} else if outcome == types.Outcome_Draw {
		amountA = QuoCoin(reward, 2)
		amountB = QuoCoin(reward, 2)
	}
	if outcome != types.Outcome_Aborted {
		k.SubPoolCredits(ctx, WinnersPoolKey, reward)
	}
	return
}

// GetMatchReward Calculates winner rewards
func (k Keeper) GetMatchReward(ctx sdk.Context) sdk.Coin {
	pool := k.Pools.Get(ctx, WinnersPoolKey)
	reward := QuoCoin(*pool, k.GetParams(ctx).WinnerReward)
	if reward.Amount.Int64() > 1000000 {
		return sdk.NewInt64Coin(reward.Denom, 1000000)
	}
	return reward
}

func (k Keeper) GetMatchAddresses(ctx sdk.Context, match types.Match) (addresses []sdk.AccAddress, err error) {
	for _, player := range []string{match.PlayerA, match.PlayerB} {
		var address sdk.AccAddress
		address, err = sdk.AccAddressFromBech32(player)
		if err != nil {
			err = sdkerrors.Wrap(types.ErrInvalidAccAddress, "Invalid player")
			return
		}
		addresses = append(addresses, address)
	}

	return
}
