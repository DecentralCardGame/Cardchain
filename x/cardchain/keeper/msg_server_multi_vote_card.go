package keeper

import (
	"context"

	sdkerrors "cosmossdk.io/errors"
	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) MultiVoteCard(goCtx context.Context, msg *types.MsgMultiVoteCard) (*types.MsgMultiVoteCardResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	var n = 0

	voter, err := k.GetUserFromString(ctx, msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidAccAddress, "Unable to convert to AccAddress")
	}

	for i, vote := range msg.Votes {
		err := k.voteCard(ctx, &voter, vote.CardId, vote.VoteType)
		if err != nil {
			return nil, err
		}
		n = i
	}

	k.incVotesAverageBy(ctx, int64(n+1))

	k.ClaimAirDrop(ctx, &voter, types.AirDrop_vote)
	k.SetUserFromUser(ctx, voter)

	return &types.MsgMultiVoteCardResponse{}, nil
}
