package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CardVote(goCtx context.Context, msg *types.MsgCardVote) (*types.MsgCardVoteResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	voter, err := k.GetUserFromString(ctx, msg.Creator)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "Unable to convert to AccAddress")
	}

	err = k.voteCard(ctx, &voter, msg.Vote.CardId, msg.Vote.VoteType)
	if err != nil {
		return nil, err
	}

	k.incVotesAverageBy(ctx, 1)

	claimedAirdrop := k.ClaimAirDrop(ctx, &voter, types.AirDrop_vote)
	k.SetUserFromUser(ctx, voter)

	return &types.MsgCardVoteResponse{AirdropClaimed: claimedAirdrop}, nil
}
