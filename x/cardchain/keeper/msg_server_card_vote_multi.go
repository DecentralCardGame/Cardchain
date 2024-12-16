package keeper

import (
	"context"

	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CardVoteMulti(goCtx context.Context, msg *types.MsgCardVoteMulti) (*types.MsgCardVoteMultiResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCardVoteMultiResponse{}, nil
}
