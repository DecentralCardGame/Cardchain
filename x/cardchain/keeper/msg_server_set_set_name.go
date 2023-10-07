package keeper

import (
	"context"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SetSetName(goCtx context.Context, msg *types.MsgSetSetName) (*types.MsgSetSetNameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	set := k.Sets.Get(ctx, msg.SetId)
	err := checkSetEditable(set, msg.Creator)
	if err != nil {
		return nil, err
	}

	set.Name = msg.Name

	k.Sets.Set(ctx, msg.SetId, set)

	return &types.MsgSetSetNameResponse{}, nil
}
