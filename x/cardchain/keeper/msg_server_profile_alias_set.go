package keeper

import (
	"context"

	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) ProfileAliasSet(goCtx context.Context, msg *types.MsgProfileAliasSet) (*types.MsgProfileAliasSetResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgProfileAliasSetResponse{}, nil
}
