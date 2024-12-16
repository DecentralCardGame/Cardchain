package keeper

import (
	"context"

	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CardArtistChange(goCtx context.Context, msg *types.MsgCardArtistChange) (*types.MsgCardArtistChangeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCardArtistChangeResponse{}, nil
}
