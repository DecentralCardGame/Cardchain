package keeper

import (
	"context"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) MultiVoteCard(goCtx context.Context, msg *types.MsgMultiVoteCard) (*types.MsgMultiVoteCardResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	var n int = 0
	
	voter, err := k.GetUserFromString(ctx, msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidAccAddress, "Unable to convert to AccAddress")
	}
	
	defer k.incVotesAverageBy(ctx, int64(n + 1))
	defer k.SetUserFromUser(ctx, voter)
	
	voteRights := voter.VoteRights
	
	for i, vote := range msg.Votes {
		err := k.voteCard(ctx, &voter, &voteRights, vote.CardId, vote.VoteType)
		if err != nil {
			return nil, err	
		}
		n = i
	}
	
	k.incVotesAverageBy(ctx, 1)

	k.ClaimAirDrop(ctx, &voter, types.AirDrop_vote)

	return &types.MsgMultiVoteCardResponse{}, nil
}
