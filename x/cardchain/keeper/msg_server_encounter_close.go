package keeper

import (
	"context"

	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) EncounterClose(goCtx context.Context, msg *types.MsgEncounterClose) (*types.MsgEncounterCloseResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgEncounterCloseResponse{}, nil
}
