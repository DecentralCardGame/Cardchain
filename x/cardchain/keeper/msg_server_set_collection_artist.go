package keeper

import (
	"context"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SetCollectionArtist(goCtx context.Context, msg *types.MsgSetCollectionArtist) (*types.MsgSetCollectionArtistResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	collection := k.Collections.Get(ctx, msg.CollectionId)
	err := checkCollectionEditable(collection, msg.Creator)
	if err != nil {
		return nil, err
	}

	collection.Artist = msg.Artist

	k.Collections.Set(ctx, msg.CollectionId, collection)

	return &types.MsgSetCollectionArtistResponse{}, nil
}
