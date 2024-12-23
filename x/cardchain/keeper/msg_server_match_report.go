package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) MatchReport(goCtx context.Context, msg *types.MsgMatchReport) (*types.MsgMatchReportResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	match := k.Matches.Get(ctx, msg.MatchId)

	creator, err := k.GetUserFromString(ctx, msg.Creator)
	if err != nil {
		return nil, errorsmod.Wrap(types.ErrUserDoesNotExist, err.Error())
	}
	if creator.ReportMatches == false {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Reporter")
	}
	if msg.Creator != match.Reporter {
		return nil, errorsmod.Wrapf(sdkerrors.ErrUnauthorized, "Wrong Reporter, reporter is %s", match.Reporter)
	}
	if match.CoinsDistributed {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "Match already reported")
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

	return &types.MsgMatchReportResponse{}, nil
}
