package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CardVoteMulti(goCtx context.Context, msg *types.MsgCardVoteMulti) (*types.MsgCardVoteMultiResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	voter, err := k.GetUserFromString(ctx, msg.Creator)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "Unable to convert to AccAddress")
	}

	err = k.multiVote(ctx, &voter, msg.Votes)
	if err != nil {
		return nil, err
	}

	k.ClaimAirDrop(ctx, &voter, types.AirDrop_vote)
	k.SetUserFromUser(ctx, voter)

	return &types.MsgCardVoteMultiResponse{}, nil
}
