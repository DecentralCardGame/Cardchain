package keeper

import (
	"context"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) AddArtwork(goCtx context.Context, msg *types.MsgAddArtwork) (*types.MsgAddArtworkResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgAddArtworkResponse{}, nil
}
