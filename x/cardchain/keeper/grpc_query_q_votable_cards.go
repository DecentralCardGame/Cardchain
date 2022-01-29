package keeper

import (
	"context"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) QVotableCards(goCtx context.Context, req *types.QueryQVotableCardsRequest) (*types.QueryQVotableCardsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	address, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "could not parse user address")
	}

	unreg := true
	noRights := true
	voteRights := []*types.VoteRight{}

	user, err := k.GetUser(ctx, address)
	if err == nil {
		unreg = false
		noRights = len(user.VoteRights) == 0
		voteRights = user.VoteRights
	}

	return &types.QueryQVotableCardsResponse{unreg, noRights, voteRights}, nil
}
