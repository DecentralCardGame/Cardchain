package keeper

import (
	"context"
	"fmt"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) FinalizeCollection(goCtx context.Context, msg *types.MsgFinalizeCollection) (*types.MsgFinalizeCollectionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	collectionSize := int(k.GetParams(ctx).CollectionSize)

	collection := k.Collections.Get(ctx, msg.CollectionId)
	if msg.Creator != collection.Contributors[0] {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Invalid creator")
	}

	if collection.Status != types.CStatus_design {
		return nil, types.ErrCollectionNotInDesign
	}

	if len(collection.Cards) != collectionSize {
		return nil, sdkerrors.Wrapf(types.ErrCollectionSize, "Has to be %d", collectionSize)
	}

	err := k.CollectCollectionCreationFee(ctx, msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, err.Error())
	}

	dist, err := k.GetRarityDistribution(ctx, *collection, uint64(collectionSize))
	if err != nil {
		return nil, err
	}

	if dist[0] != dist[1] {
		return nil, fmt.Errorf("Collections should contain (c,u,r) %d, %d, %d but contains %d, %d, %d", dist[1][0], dist[1][1], dist[1][2], dist[0][0], dist[0][1], dist[0][2])
	}

	collection.Status = types.CStatus_finalized

	k.Collections.Set(ctx, msg.CollectionId, collection)

	return &types.MsgFinalizeCollectionResponse{}, nil
}
