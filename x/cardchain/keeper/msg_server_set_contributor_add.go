package keeper

import (
	"context"

	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SetContributorAdd(goCtx context.Context, msg *types.MsgSetContributorAdd) (*types.MsgSetContributorAddResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSetContributorAddResponse{}, nil
}
