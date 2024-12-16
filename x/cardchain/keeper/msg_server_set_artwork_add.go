package keeper

import (
	"context"

	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SetArtworkAdd(goCtx context.Context, msg *types.MsgSetArtworkAdd) (*types.MsgSetArtworkAddResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSetArtworkAddResponse{}, nil
}
