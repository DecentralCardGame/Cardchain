package keeper

import (
	"context"

	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) ProfileAliasSet(goCtx context.Context, msg *types.MsgProfileAliasSet) (*types.MsgProfileAliasSetResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	user, err := k.GetUserFromString(ctx, msg.Creator)
	if err != nil {
		return nil, err
	}

	user.Alias = msg.Alias

	k.SetUserFromUser(ctx, user)

	return &types.MsgProfileAliasSetResponse{}, nil
}
