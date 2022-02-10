package keeper

import (
	"context"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) AddCardToCollection(goCtx context.Context, msg *types.MsgAddCardToCollection) (*types.MsgAddCardToCollectionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgAddCardToCollectionResponse{}, nil
}
