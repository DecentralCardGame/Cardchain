package keeper

import (
	"context"
	"fmt"

	errorsmod "cosmossdk.io/errors"
	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateProductDetails(goCtx context.Context, msg *types.MsgCreateProductDetails) (*types.MsgCreateProductDetailsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var productDetails = types.ProductDetails{
		Creator: msg.Creator,
		Name:    msg.Name,
		Desc:    msg.Desc,
	}

	id := k.AppendProductDetails(
		ctx,
		productDetails,
	)

	return &types.MsgCreateProductDetailsResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateProductDetails(goCtx context.Context, msg *types.MsgUpdateProductDetails) (*types.MsgUpdateProductDetailsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var productDetails = types.ProductDetails{
		Creator: msg.Creator,
		Id:      msg.Id,
		Name:    msg.Name,
		Desc:    msg.Desc,
	}

	// Checks that the element exists
	val, found := k.GetProductDetails(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetProductDetails(ctx, productDetails)

	return &types.MsgUpdateProductDetailsResponse{}, nil
}

func (k msgServer) DeleteProductDetails(goCtx context.Context, msg *types.MsgDeleteProductDetails) (*types.MsgDeleteProductDetailsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetProductDetails(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveProductDetails(ctx, msg.Id)

	return &types.MsgDeleteProductDetailsResponse{}, nil
}
