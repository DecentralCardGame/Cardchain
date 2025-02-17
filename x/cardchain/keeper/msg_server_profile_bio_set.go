package keeper

import (
	"context"

	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) ProfileBioSet(goCtx context.Context, msg *types.MsgProfileBioSet) (*types.MsgProfileBioSetResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	user, err := k.GetUserFromString(ctx, msg.Creator)
	if err != nil {
		return nil, err
	}

	user.Biography = msg.Bio

	k.SetUserFromUser(ctx, user)

	return &types.MsgProfileBioSetResponse{}, nil
}
