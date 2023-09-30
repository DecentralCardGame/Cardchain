package keeper

import (
	"context"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SetSetArtist(goCtx context.Context, msg *types.MsgSetSetArtist) (*types.MsgSetSetArtistResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	set := k.Sets.Get(ctx, msg.SetId)
	err := checkSetEditable(set, msg.Creator)
	if err != nil {
		return nil, err
	}

	set.Artist = msg.Artist

	k.Sets.Set(ctx, msg.SetId, set)

	return &types.MsgSetSetArtistResponse{}, nil
}

