package keeper

import (
	"context"

	sdkerrors "cosmossdk.io/errors"
	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) AddStoryToSet(goCtx context.Context, msg *types.MsgAddStoryToSet) (*types.MsgAddStoryToSetResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	set := k.Sets.Get(ctx, msg.SetId)
	if set.StoryWriter != msg.Creator {
		return nil, sdkerrors.Wrap(errors.ErrUnauthorized, "Incorrect StoryWriter")
	}
	if set.Status != types.CStatus_design {
		return nil, types.ErrSetNotInDesign
	}

	err := k.CollectSetConributionFee(ctx, msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(errors.ErrInsufficientFunds, err.Error())
	}

	set.Story = msg.Story

	k.Sets.Set(ctx, msg.SetId, set)

	return &types.MsgAddStoryToSetResponse{}, nil
}
