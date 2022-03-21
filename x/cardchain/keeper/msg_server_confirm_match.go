package keeper

import (
	"context"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) ConfirmMatch(goCtx context.Context, msg *types.MsgConfirmMatch) (*types.MsgConfirmMatchResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	match := k.Matches.Get(ctx, msg.MatchId)
	if msg.Creator != match.PlayerA || msg.Creator != match.PlayerB {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Didn't participate in match")
	}

	if (match.OutcomeA != types.Outcome_Aborted && msg.Creator == match.PlayerA) || (match.OutcomeB != types.Outcome_Aborted && msg.Creator == match.PlayerB) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Already reported")
	}

	if match.Outcome == types.Outcome_Aborted {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Can't report, because match was aborted")
	}

	return &types.MsgConfirmMatchResponse{}, nil
}
