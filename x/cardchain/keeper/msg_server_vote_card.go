package keeper

import (
	"context"

	sdkerrors "cosmossdk.io/errors"
	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) VoteCard(goCtx context.Context, msg *types.MsgVoteCard) (*types.MsgVoteCardResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	voter, err := k.GetUserFromString(ctx, msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidAccAddress, "Unable to convert to AccAddress")
	}

	err = k.voteCard(ctx, &voter, msg.CardId, msg.VoteType, false)
	if err != nil {
		return nil, err
	}

	k.incVotesAverageBy(ctx, 1)

	claimedAirdrop := k.ClaimAirDrop(ctx, &voter, types.AirDrop_vote)
	k.SetUserFromUser(ctx, voter)

	return &types.MsgVoteCardResponse{AirdropClaimed: claimedAirdrop}, nil
}
