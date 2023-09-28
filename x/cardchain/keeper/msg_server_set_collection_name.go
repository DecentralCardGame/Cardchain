package keeper

import (
	"context"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SetCollectionName(goCtx context.Context, msg *types.MsgSetCollectionName) (*types.MsgSetCollectionNameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	collection := k.Collections.Get(ctx, msg.CollectionId)
	err := checkCollectionEditable(collection, msg.Creator)
	if err != nil {
		return nil, err
	}
	
	collection.Name = msg.Name

	k.Collections.Set(ctx, msg.CollectionId, collection)

	return &types.MsgSetCollectionNameResponse{}, nil
}
