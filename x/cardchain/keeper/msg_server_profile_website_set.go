package keeper

import (
	"context"

	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) ProfileWebsiteSet(goCtx context.Context, msg *types.MsgProfileWebsiteSet) (*types.MsgProfileWebsiteSetResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	user, err := k.GetUserFromString(ctx, msg.Creator)
	if err != nil {
		return nil, err
	}

	user.Website = msg.Website

	k.SetUserFromUser(ctx, user)

	return &types.MsgProfileWebsiteSetResponse{}, nil
}
