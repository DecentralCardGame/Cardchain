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
		types.Outcome_Aborted,
		types.Outcome_Aborted,
		false,
	}

	addresses, err := k.GetMatchAddresses(ctx, match)
	if err != nil {
		return nil, err
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

	return &types.MsgReportMatchResponse{}, nil
}
