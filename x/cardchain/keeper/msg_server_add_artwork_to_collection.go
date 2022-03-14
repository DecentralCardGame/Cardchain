package keeper

import (
	"context"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) AddArtworkToCollection(goCtx context.Context, msg *types.MsgAddArtworkToCollection) (*types.MsgAddArtworkToCollectionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	collection := k.GetCollection(ctx, msg.CollectionId)
	if collection.Artist != msg.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Artist")
	}
	if collection.Status != types.CStatus_design {
		return nil, types.ErrCollectionNotInDesign
	}

	err := k.CollectCollectionConributionFee(ctx, msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, err.Error())
	}

	collection.Artwork = msg.Image

	k.SetCollection(ctx, msg.CollectionId, collection)

	return &types.MsgAddArtworkToCollectionResponse{}, nil
}
