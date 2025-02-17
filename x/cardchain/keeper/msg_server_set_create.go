package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) SetCreate(goCtx context.Context, msg *types.MsgSetCreate) (*types.MsgSetCreateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.CollectSetCreationFee(ctx, msg.Creator)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInsufficientFunds, err.Error())
	}

	set := types.Set{
		Name:         msg.Name,
		Cards:        []uint64{},
		Contributors: append([]string{msg.Creator}, msg.Contributors...),
		Artist:       msg.Artist,
		StoryWriter:  msg.StoryWriter,
		Status:       types.SetStatus_design,
		TimeStamp:    0,
		ArtworkId:    k.Images.GetNum(ctx),
	}

	image := types.Image{}

	k.Images.Set(ctx, set.ArtworkId, &image)
	k.Setk.Append(ctx, &set)

	return &types.MsgSetCreateResponse{}, nil
}
