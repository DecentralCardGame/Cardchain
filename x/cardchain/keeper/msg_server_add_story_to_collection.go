package keeper

import (
	"context"

	sdkerrors "cosmossdk.io/errors"
	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) AddStoryToCollection(goCtx context.Context, msg *types.MsgAddStoryToCollection) (*types.MsgAddStoryToCollectionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	collection := k.Collections.Get(ctx, msg.CollectionId)
	if collection.StoryWriter != msg.Creator {
		return nil, sdkerrors.Wrap(errors.ErrUnauthorized, "Incorrect StoryWriter")
	}
	if collection.Status != types.CStatus_design {
		return nil, types.ErrCollectionNotInDesign
	}

	err := k.CollectCollectionConributionFee(ctx, msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(errors.ErrInsufficientFunds, err.Error())
	}

	collection.Story = msg.Story

	k.Collections.Set(ctx, msg.CollectionId, collection)

	return &types.MsgAddStoryToCollectionResponse{}, nil
}
