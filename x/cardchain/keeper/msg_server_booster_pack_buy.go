package keeper

import (
	"context"

	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) BoosterPackBuy(goCtx context.Context, msg *types.MsgBoosterPackBuy) (*types.MsgBoosterPackBuyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgBoosterPackBuyResponse{}, nil
}
