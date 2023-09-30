package keeper

import (
	"context"

	sdkerrors "cosmossdk.io/errors"
	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateSet(goCtx context.Context, msg *types.MsgCreateSet) (*types.MsgCreateSetResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	setId := k.Sets.GetNum(ctx)

	err := k.CollectSetCreationFee(ctx, msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(errors.ErrInsufficientFunds, err.Error())
	}

	set := types.Set{
		Name:         msg.Name,
		Cards:        []uint64{},
		Contributors: append([]string{msg.Creator}, msg.Contributors...),
		Artist:       msg.Artist,
		StoryWriter:  msg.StoryWriter,
		Status:       types.CStatus_design,
		TimeStamp:    0,
		ArtworkId:    k.Images.GetNum(ctx),
	}

	image := types.Image{}

	k.Images.Set(ctx, set.ArtworkId, &image)
	k.Sets.Set(ctx, setId, &set)

	return &types.MsgCreateSetResponse{}, nil
}
