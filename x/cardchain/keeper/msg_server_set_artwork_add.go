package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) SetArtworkAdd(goCtx context.Context, msg *types.MsgSetArtworkAdd) (*types.MsgSetArtworkAddResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	set := k.SetK.Get(ctx, msg.SetId)
	image := k.Images.Get(ctx, set.ArtworkId)

	if set.Artist != msg.Creator {
		return nil, errorsmod.Wrap(errors.ErrUnauthorized, "Incorrect Artist")
	}
	if set.Status != types.SetStatus_design {
		return nil, errorsmod.Wrapf(errors.ErrUnauthorized, "Invalid set status is: %s", set.Status.String())
	}

	err := k.CollectSetConributionFee(ctx, msg.Creator)
	if err != nil {
		return nil, errorsmod.Wrap(errors.ErrInsufficientFunds, err.Error())
	}

	image.Image = msg.Image

	k.Images.Set(ctx, set.ArtworkId, image)

	return &types.MsgSetArtworkAddResponse{}, nil
}
