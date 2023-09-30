package keeper

import (
	"context"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) OpenMatch(goCtx context.Context, msg *types.MsgOpenMatch) (*types.MsgOpenMatchResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	creator, err := k.GetUserFromString(ctx, msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrUserDoesNotExist, err.Error())
	}
	if creator.ReportMatches == false {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Reporter")
	}

	matchId := k.Matches.GetNum(ctx)

	match := types.Match{
		Reporter:         msg.Creator,
		PlayerA:          types.NewMatchPlayer(msg.PlayerA, []uint64{}, msg.PlayerADeck),
		PlayerB:          types.NewMatchPlayer(msg.PlayerB, []uint64{}, msg.PlayerBDeck),
		CoinsDistributed: false,
	}

	k.Matches.Set(ctx, matchId, &match)

	return &types.MsgOpenMatchResponse{MatchId: matchId}, nil
}
