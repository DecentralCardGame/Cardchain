package keeper

import (
	"context"

	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) ProfileWebsiteSet(goCtx context.Context, msg *types.MsgProfileWebsiteSet) (*types.MsgProfileWebsiteSetResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgProfileWebsiteSetResponse{}, nil
}
