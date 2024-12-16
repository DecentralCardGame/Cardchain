package keeper

import (
	"context"

	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SellOfferRemove(goCtx context.Context, msg *types.MsgSellOfferRemove) (*types.MsgSellOfferRemoveResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSellOfferRemoveResponse{}, nil
}
