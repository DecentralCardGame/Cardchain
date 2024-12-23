package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"github.com/DecentralCardGame/cardchain/util"
	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SetContributorRemove(goCtx context.Context, msg *types.MsgSetContributorRemove) (*types.MsgSetContributorRemoveResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	set := k.Sets.Get(ctx, msg.SetId)
	if err := checkSetEditable(set, msg.Creator); err != nil {
		return nil, err
	}

	newContributors, err := util.PopItemFromArr(msg.User, set.Contributors)
	if err != nil {
		return nil, errorsmod.Wrap(err, "Contributor is not a contributor: "+msg.User)
	}

	set.Contributors = newContributors

	k.Sets.Set(ctx, msg.SetId, set)

	return &types.MsgSetContributorRemoveResponse{}, nil
}
