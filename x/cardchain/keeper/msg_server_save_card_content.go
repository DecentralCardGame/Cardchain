package keeper

import (
	"context"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SaveCardContent(goCtx context.Context, msg *types.MsgSaveCardContent) (*types.MsgSaveCardContentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSaveCardContentResponse{}, nil
}
