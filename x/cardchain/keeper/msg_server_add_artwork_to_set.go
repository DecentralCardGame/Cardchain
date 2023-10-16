package keeper

import (
	"context"
	"fmt"

	sdkerrors "cosmossdk.io/errors"
	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) AddArtworkToSet(goCtx context.Context, msg *types.MsgAddArtworkToSet) (*types.MsgAddArtworkToSetResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	set := k.Sets.Get(ctx, msg.SetId)
	image := k.Images.Get(ctx, set.ArtworkId)

	if set.Artist != msg.Creator {
		return nil, sdkerrors.Wrap(errors.ErrUnauthorized, "Incorrect Artist")
	}
	if set.Status != types.CStatus_design {
		return nil, types.ErrSetNotInDesign
	}

	if len(msg.Image) > 500000 {
		return nil, sdkerrors.Wrap(types.ErrImageSizeExceeded, fmt.Sprint(len(msg.Image)))
	}

	err := k.CollectSetConributionFee(ctx, msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(errors.ErrInsufficientFunds, err.Error())
	}

	image.Image = msg.Image

	k.Images.Set(ctx, set.ArtworkId, image)

	return &types.MsgAddArtworkToSetResponse{}, nil
}
