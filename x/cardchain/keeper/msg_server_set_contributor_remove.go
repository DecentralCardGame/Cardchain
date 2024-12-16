package keeper

import (
	"context"

	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SetContributorRemove(goCtx context.Context, msg *types.MsgSetContributorRemove) (*types.MsgSetContributorRemoveResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSetContributorRemoveResponse{}, nil
}
