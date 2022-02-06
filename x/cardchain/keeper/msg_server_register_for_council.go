package keeper

import (
	"context"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) RegisterForCouncil(goCtx context.Context, msg *types.MsgRegisterForCouncil) (*types.MsgRegisterForCouncilResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgRegisterForCouncilResponse{}, nil
}
