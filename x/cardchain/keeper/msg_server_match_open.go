package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) MatchOpen(goCtx context.Context, msg *types.MsgMatchOpen) (*types.MsgMatchOpenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	creator, err := k.GetUserFromString(ctx, msg.Creator)
	if err != nil {
		return nil, errorsmod.Wrap(types.ErrUserDoesNotExist, err.Error())
	}
	if creator.ReportMatches == false {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Reporter")
	}

	matchId := k.Matches.GetNum(ctx)

	match := types.Match{
		Reporter:         msg.Creator,
		PlayerA:          types.NewMatchPlayer(msg.PlayerA, []uint64{}, msg.PlayerADeck),
		PlayerB:          types.NewMatchPlayer(msg.PlayerB, []uint64{}, msg.PlayerBDeck),
		CoinsDistributed: false,
	}

	k.Matches.Set(ctx, matchId, &match)

	return &types.MsgMatchOpenResponse{MatchId: matchId}, nil
}
