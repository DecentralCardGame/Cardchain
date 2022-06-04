package keeper

import (
	"fmt"
	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"golang.org/x/exp/slices"
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

// GetMatchAddresses Get's and verifies the players of a match
func (k Keeper) GetMatchAddresses(ctx sdk.Context, match types.Match) (addresses []sdk.AccAddress, err error) {
	for _, player := range []string{match.PlayerA.Addr, match.PlayerB.Addr} {
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

// DistributeCoins to players of a match
func (k Keeper) DistributeCoins(ctx sdk.Context, match *types.Match, outcome types.Outcome) error {
	addresses, err := k.GetMatchAddresses(ctx, *match)
	if err != nil {
		return err
	}

	amountA, amountB := k.CalculateMatchReward(ctx, outcome)
	amounts := []sdk.Coin{amountA, amountB}
	for idx, address := range addresses {
		if !amounts[idx].IsZero() {
			err := k.MintCoinsToAddr(ctx, address, sdk.Coins{amounts[idx]})
			if err != nil {
				return sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, err.Error())
			}
			k.SubPoolCredits(ctx, WinnersPoolKey, amounts[idx])

			user, err := k.GetUser(ctx, address)
			if err != nil {
				return err
			}
			userObj := User{Addr: address, User: user}
			k.ClaimAirDrop(ctx, &userObj, types.AirDrop_play)
			k.SetUserFromUser(ctx, userObj)
		}
	}

	games := k.RunningAverages.Get(ctx, Games24ValueKey)
	games.Arr[len(games.Arr)-1]++
	k.RunningAverages.Set(ctx, Games24ValueKey, games)

	match.CoinsDistributed = true

	if outcome != types.Outcome_Aborted {
		for idx, address := range addresses {
			for _, cardId := range [][]uint64{match.PlayerA.PlayedCards, match.PlayerB.PlayedCards}[idx] {
				err = k.AddVoteRight(ctx, address, cardId)
				if err != nil {
					return sdkerrors.Wrap(types.ErrUserDoesNotExist, err.Error())
				}
			}
		}
	}

	return nil
}

// GetOutcome Get's the outcome of a match
func (k Keeper) GetOutcome(ctx sdk.Context, match types.Match) (outcome types.Outcome, err error) {
	// majority based vote
	outcomes := []types.Outcome{match.Outcome, match.PlayerA.Outcome, match.PlayerB.Outcome}
	slices.Sort(outcomes)
	switch outcomes[1] {
	case outcomes[0], outcomes[2]:
		outcome = outcomes[1]
	default:
		if match.PlayerA.Confirmed && match.PlayerB.Confirmed {
			outcome = types.Outcome_Aborted
		} else {
			err = fmt.Errorf("Not enought votes for decision")
		}
	}
	return
}
