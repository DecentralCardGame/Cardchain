package keeper

import (
	"context"

	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CardSchemeBuy(goCtx context.Context, msg *types.MsgCardSchemeBuy) (*types.MsgCardSchemeBuyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCardSchemeBuyResponse{}, nil
}
