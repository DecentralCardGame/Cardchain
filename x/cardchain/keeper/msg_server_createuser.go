package keeper

import (
	"context"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Createuser(goCtx context.Context, msg *types.MsgCreateuser) (*types.MsgCreateuserResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.CreateUser(ctx, msg.GetNewUserAddr(), msg.Alias)

	return &types.MsgCreateuserResponse{}, err
}
