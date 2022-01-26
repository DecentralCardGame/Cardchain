package keeper

import (
	"context"
	"fmt"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	"github.com/cosmos/cosmos-sdk/codec"
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

	user := k.GetUser(ctx, address)

	res, err := codec.MarshalJSONIndent(codec.NewLegacyAmino(), user)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	fmt.Printf("User: %s\n", user)

	return &types.QueryQUserResponse{res}, nil
}
