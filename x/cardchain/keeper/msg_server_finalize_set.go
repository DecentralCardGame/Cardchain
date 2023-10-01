package keeper

import (
	"context"
	"fmt"

	sdkerrors "cosmossdk.io/errors"
	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) FinalizeSet(goCtx context.Context, msg *types.MsgFinalizeSet) (*types.MsgFinalizeSetResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	setSize := int(k.GetParams(ctx).SetSize)

	set := k.Sets.Get(ctx, msg.SetId)
	err := checkSetEditable(set, msg.Creator)
	if err != nil  {
		return nil, err
	}

	err = k.CollectSetCreationFee(ctx, msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(errors.ErrInsufficientFunds, err.Error())
	}

	dist, err := k.GetRarityDistribution(ctx, *set, uint32(setSize))
	if err != nil {
		return nil, err
	}

	if dist[0] != dist[1] {
		return nil, fmt.Errorf("sets should contain [(common,unique,exceptional), uncommon, rare] %d, %d, %d but contains %d, %d, %d", dist[1][0], dist[1][1], dist[1][2], dist[0][0], dist[0][1], dist[0][2])
	}

	set.Status = types.CStatus_finalized

	k.Sets.Set(ctx, msg.SetId, set)

	return &types.MsgFinalizeSetResponse{}, nil
}
