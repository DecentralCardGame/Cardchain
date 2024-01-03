package keeper

import (
	"context"

	sdkerrors "cosmossdk.io/errors"
	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) ReportMatch(goCtx context.Context, msg *types.MsgReportMatch) (*types.MsgReportMatchResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	match := k.Matches.Get(ctx, msg.MatchId)

	creator, err := k.GetUserFromString(ctx, msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrUserDoesNotExist, err.Error())
	}
	if creator.ReportMatches == false {
		return nil, sdkerrors.Wrap(errors.ErrUnauthorized, "Incorrect Reporter")
	}
	if msg.Creator != match.Reporter {
		return nil, sdkerrors.Wrapf(errors.ErrUnauthorized, "Wrong Reporter, reporter is %s", match.Reporter)
	}
	if match.CoinsDistributed {
		return nil, sdkerrors.Wrap(errors.ErrUnauthorized, "Match already reported")
	}

	match.PlayerA.PlayedCards = msg.PlayedCardsA
	match.PlayerB.PlayedCards = msg.PlayedCardsB
	match.Outcome = msg.Outcome
	match.ServerConfirmed = true
	match.Timestamp = uint64(ctx.BlockHeight())

	err = k.TryHandleMatchOutcome(ctx, match)
	if err != nil {
		return nil, err
	}

	k.Matches.Set(ctx, msg.MatchId, match)

	return &types.MsgReportMatchResponse{}, nil
}
