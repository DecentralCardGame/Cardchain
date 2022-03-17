package keeper

import (
	"context"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateCollection(goCtx context.Context, msg *types.MsgCreateCollection) (*types.MsgCreateCollectionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	collectionId := k.Collections.GetNumber(ctx)

	err := k.CollectCollectionCreationFee(ctx, msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, err.Error())
	}

	collection := types.Collection{
		Name:         msg.Name,
		Cards:        []uint64{},
		Contributors: append([]string{msg.Creator}, msg.Contributors...),
		Artist:       msg.Artist,
		StoryWriter:  msg.StoryWriter,
		Status:       types.CStatus_design,
		TimeStamp:    0,
	}

	k.Collections.Set(ctx, collectionId, &collection)

	return &types.MsgCreateCollectionResponse{}, nil
}
