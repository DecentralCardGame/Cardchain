package keeper

import (
	"context"

	sdkerrors "cosmossdk.io/errors"
	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) RemoveContributorFromSet(goCtx context.Context, msg *types.MsgRemoveContributorFromSet) (*types.MsgRemoveContributorFromSetResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	set := k.Sets.Get(ctx, msg.SetId)
	if msg.Creator != set.Contributors[0] {
		return nil, sdkerrors.Wrap(errors.ErrUnauthorized, "Invalid creator")
	}
	if set.Status != types.CStatus_design {
		return nil, types.ErrSetNotInDesign
	}

	newContributors, err := stringPopElementFromArr(msg.User, set.Contributors)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "Contributor is not a contributor: "+msg.User)
	}

	set.Contributors = newContributors

	k.Sets.Set(ctx, msg.SetId, set)

	return &types.MsgRemoveContributorFromSetResponse{}, nil
}

func stringPopElementFromArr(element string, arr []string) ([]string, error) {
	for idx, val := range arr {
		if element == val {
			return append(arr[:idx], arr[idx+1:]...), nil
		}
	}
	return []string{}, types.ErrContributor
}
