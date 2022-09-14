package keeper

import (
	"context"
	"fmt"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) AddArtworkToCollection(goCtx context.Context, msg *types.MsgAddArtworkToCollection) (*types.MsgAddArtworkToCollectionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	collection := k.Collections.Get(ctx, msg.CollectionId)
	image := k.Images.Get(ctx, collection.ArtworkId)

	if collection.Artist != msg.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Artist")
	}
	if collection.Status != types.CStatus_design {
		return nil, types.ErrCollectionNotInDesign
	}

	if len(msg.Image) > 500000 {
		return nil, sdkerrors.Wrap(types.ErrImageSizeExceeded, fmt.Sprint(len(msg.Image)))
	}

	err := k.CollectCollectionConributionFee(ctx, msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, err.Error())
	}

	image.Image = msg.Image

	k.Images.Set(ctx, collection.ArtworkId, image)

	return &types.MsgAddArtworkToCollectionResponse{}, nil
}
