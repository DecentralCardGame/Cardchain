package keeper

import (
	"context"

	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) EarlyAccessInvite(goCtx context.Context, msg *types.MsgEarlyAccessInvite) (*types.MsgEarlyAccessInviteResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgEarlyAccessInviteResponse{}, nil
}
