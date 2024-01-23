package keeper

import (
	"context"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) ChangeAlias(goCtx context.Context, msg *types.MsgChangeAlias) (*types.MsgChangeAliasResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	user, err := k.GetUserFromString(ctx, msg.Creator)
	if err != nil {
		return nil, err
	}
	
	user.Alias = msg.Alias
	
	k.SetUserFromUser(ctx, user)

	return &types.MsgChangeAliasResponse{}, nil
}
