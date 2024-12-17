package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) SetStoryAdd(goCtx context.Context, msg *types.MsgSetStoryAdd) (*types.MsgSetStoryAddResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	set := k.Sets.Get(ctx, msg.SetId)
	if set.StoryWriter != msg.Creator {
		return nil, errorsmod.Wrap(errors.ErrUnauthorized, "Incorrect StoryWriter")
	}
	if set.Status != types.CStatus_design {
		return nil, errorsmod.Wrapf(errors.ErrUnauthorized, "Invalid set status is: %s", set.Status.String())
	}

	err := k.CollectSetConributionFee(ctx, msg.Creator)
	if err != nil {
		return nil, errorsmod.Wrap(errors.ErrInsufficientFunds, err.Error())
	}

	set.Story = msg.Story

	k.Sets.Set(ctx, msg.SetId, set)
	return &types.MsgSetStoryAddResponse{}, nil
}
