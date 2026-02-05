package keeper

import (
	"context"

	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SetArtistSet(goCtx context.Context, msg *types.MsgSetArtistSet) (*types.MsgSetArtistSetResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	set := k.SetK.Get(ctx, msg.SetId)
	err := checkSetEditable(set, msg.Creator)
	if err != nil {
		return nil, err
	}

	set.Artist = msg.Artist

	k.SetK.Set(ctx, msg.SetId, set)

	return &types.MsgSetArtistSetResponse{}, nil
}
