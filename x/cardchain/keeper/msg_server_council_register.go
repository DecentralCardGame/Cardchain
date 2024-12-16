package keeper

import (
	"context"

	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CouncilRegister(goCtx context.Context, msg *types.MsgCouncilRegister) (*types.MsgCouncilRegisterResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCouncilRegisterResponse{}, nil
}
