package keeper

import (
	"context"

	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SetNameSet(goCtx context.Context, msg *types.MsgSetNameSet) (*types.MsgSetNameSetResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	set := k.SetK.Get(ctx, msg.SetId)
	err := checkSetEditable(set, msg.Creator)
	if err != nil {
		return nil, err
	}

	set.Name = msg.Name

	k.SetK.Set(ctx, msg.SetId, set)

	return &types.MsgSetNameSetResponse{}, nil
}
