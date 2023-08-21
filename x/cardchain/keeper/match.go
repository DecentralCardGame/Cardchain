package keeper

import (
	"fmt"
	"time"

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

// calculateMatchReward Calculates the match winning rewards
func (k Keeper) calculateMatchReward(ctx sdk.Context, outcome types.Outcome) (amountA sdk.Coin, amountB sdk.Coin) {
	reward := k.getMatchReward(ctx)
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

// getMatchReward Calculates winner rewards
func (k Keeper) getMatchReward(ctx sdk.Context) sdk.Coin {
	pool := k.Pools.Get(ctx, WinnersPoolKey)
	reward := QuoCoin(*pool, k.GetParams(ctx).WinnerReward)
	if reward.Amount.Int64() > 1000000 {
		return sdk.NewInt64Coin(reward.Denom, 1000000)
	}
	return reward
}

// getMatchAddresses Get's and verifies the players of a match
func (k Keeper) getMatchAddresses(ctx sdk.Context, match types.Match) (addresses []sdk.AccAddress, err error) {
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

func (k Keeper) getMatchUsers(ctx sdk.Context, match types.Match) (users []*User, err error) {
	for _, address := range []string{match.PlayerA.Addr, match.PlayerB.Addr} {
		user, err := k.GetUserFromString(ctx, address)
		if err != nil {
			return []*User{}, err
		}
		users = append(users, &user)
	}

	return
}

// distributeCoins to players of a match
func (k Keeper) distributeCoins(ctx sdk.Context, match *types.Match, outcome types.Outcome) error {
	addresses, err := k.getMatchAddresses(ctx, *match)
	if err != nil {
		return err
	}

	amountA, amountB := k.calculateMatchReward(ctx, outcome)
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
			err = fmt.Errorf("not enought votes for decision")
		}
	}
	return
}

func (k Keeper) TryHandleMatchOutcome(ctx sdk.Context, match *types.Match) error {
	if match.PlayerA.Confirmed && match.PlayerB.Confirmed && match.ServerConfirmed {
		return k.HandleMatchOutcome(ctx, match)
	}
	return nil
}

func (k Keeper) HandleMatchOutcome(ctx sdk.Context, match *types.Match) error {
	// Evaluate Outcome
	outcomes := []types.Outcome{match.Outcome, match.PlayerA.Outcome, match.PlayerB.Outcome}
	slices.Sort(outcomes)
	outcomes = slices.Compact(outcomes)
	switch i := uint64(len(outcomes)); i {
	case 1:
		k.ReportServerMatch(ctx, match.Reporter, 1, true)
	default:
		k.ReportServerMatch(ctx, match.Reporter, i-1, false)
	}

	outcome, err := k.GetOutcome(ctx, *match)
	match.Outcome = outcome

	err = k.distributeCoins(ctx, match, outcome)
	if err != nil {
		return err
	}

	err = k.voteMatchCards(ctx, match)
	if err != nil {
		return err
	}

	// TODO: Votes
	return nil
}

func (k Keeper) voteMatchCards(ctx sdk.Context, match *types.Match) error {
	users, err := k.getMatchUsers(ctx, *match)
	if err != nil {
		return err
	}
	players := []*types.MatchPlayer{match.PlayerA, match.PlayerB}
	for idx, player := range players {
		// filter voted cards cards
		otherPlayer := players[(idx+1)%2]
		var otherPlayerCards []uint64
		var cleanedVotes []*types.SingleVote
		if match.ServerConfirmed {
			otherPlayerCards = otherPlayer.PlayedCards
		} else {
			otherPlayerCards = otherPlayer.Deck
		}
		for _, vote := range player.VotedCards {
			if slices.Contains(otherPlayerCards, vote.CardId) {
				cleanedVotes = append(cleanedVotes, vote)
			}
		}
		err = k.multiVote(ctx, users[idx], cleanedVotes, true)
		if err != nil {
			return err
		}
		k.SetUserFromUser(ctx, *users[idx])
	}
	return nil
}

func (k Keeper) MatchWorker(ctx sdk.Context) {
	now := uint64(time.Now().Unix())
	if ctx.BlockHeight()%20 == 0 {
		matchIter := k.Matches.GetItemIterator(ctx)
		for ; matchIter.Valid(); matchIter.Next() {
			id, match := matchIter.Value()
			if !match.CoinsDistributed && match.Timestamp != 0 && match.Timestamp+k.GetParams(ctx).MatchWorkerDelay < now {
				err := k.HandleMatchOutcome(ctx, match)
				if err != nil {
					k.Logger(ctx).Error(fmt.Sprintf(":: Error with matchWorker: %s", err))
				}
				k.Matches.Set(ctx, id, match)
			}
		}
	}
}
