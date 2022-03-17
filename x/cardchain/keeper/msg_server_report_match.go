package keeper

import (
	"context"
	"time"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) ReportMatch(goCtx context.Context, msg *types.MsgReportMatch) (*types.MsgReportMatchResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	creator, err := k.GetUserFromString(ctx, msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrUserDoesNotExist, err.Error())
	}
	if creator.ReportMatches == false {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Reporter")
	}

	matchId := k.Matches.GetNumber(ctx)

	match := types.Match{
		uint64(time.Now().Unix()),
		msg.Creator,
		msg.PlayerA,
		msg.PlayerB,
		msg.CardsA,
		msg.CardsB,
		msg.Outcome,
	}

	addresses := []sdk.AccAddress{}

	for _, player := range []string{msg.PlayerA, msg.PlayerB} {
		var address sdk.AccAddress
		address, err = sdk.AccAddressFromBech32(player)
		if err != nil {
			return nil, sdkerrors.Wrap(types.ErrInvalidAccAddress, "Invalid player")
		}
		addresses = append(addresses, address)
	}

	cards := [][]uint64{msg.CardsB, msg.CardsA}

	if msg.Outcome != types.Outcome_Aborted {
		for idx := range addresses {
			for _, cardId := range cards[idx] {
				err = k.AddVoteRight(ctx, addresses[idx], cardId)
				if err != nil {
					return nil, sdkerrors.Wrap(types.ErrUserDoesNotExist, err.Error())
				}
			}
		}
	}

	k.Matches.Set(ctx, matchId, &match)

	amountA, amountB := k.CalculateMatchReward(ctx, msg.Outcome)
	amounts := []sdk.Coin{amountA, amountB}
	for idx := range addresses {
		err := k.MintCoinsToAddr(ctx, addresses[idx], sdk.Coins{amounts[idx]})
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, err.Error())
		}
		k.SubPoolCredits(ctx, WinnersPoolKey, amounts[idx])
	}

	games := k.GetRunningAverage(ctx, Games24ValueKey)
	games.Arr[len(games.Arr)-1]++
	k.SetRunningAverage(ctx, Games24ValueKey, games)

	return &types.MsgReportMatchResponse{}, nil
}
