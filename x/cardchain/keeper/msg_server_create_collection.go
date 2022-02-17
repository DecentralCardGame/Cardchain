package keeper

import (
	"context"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateCollection(goCtx context.Context, msg *types.MsgCreateCollection) (*types.MsgCreateCollectionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	price := k.GetParams(ctx).CollectionCreationFee
	collectionId := k.GetCollectionsNumber(ctx)

	buyer, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidAccAddress, "Unable to convert to AccAddress")
	}
	err = k.BurnCoinsFromAddr(ctx, buyer, sdk.Coins{price})
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "Creator does not have enough coins")
	}
	k.AddPoolCredits(ctx, PublicPoolKey, price)

	collection := types.Collection{
		Name:         msg.Name,
		Cards:        []uint64{},
		Contributors: append([]string{msg.Creator}, msg.Contributors...),
		Artist:       msg.Artist,
		StoryWriter:  msg.StoryWriter,
		Status:       types.CStatus_design,
		TimeStamp:    0,
	}

	k.SetCollection(ctx, collectionId, collection)

	return &types.MsgCreateCollectionResponse{}, nil
}
