package keeper

import (
	"context"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) QUser(goCtx context.Context, req *types.QueryQUserRequest) (*types.QueryQUserResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	address, error := sdk.AccAddressFromBech32(req.Address)
	if error != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "could not parse user address")
	}

	user, err := k.GetUser(ctx, address)
	if err != nil {
		return nil, err
	}

	return &types.QueryQUserResponse{
		user.Alias,
		user.OwnedCardSchemes,
		user.OwnedCards,
		user.VoteRights,
	}, nil
}
