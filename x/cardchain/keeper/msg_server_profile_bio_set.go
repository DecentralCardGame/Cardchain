package keeper

import (
	"context"

	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) ProfileBioSet(goCtx context.Context, msg *types.MsgProfileBioSet) (*types.MsgProfileBioSetResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgProfileBioSetResponse{}, nil
}
