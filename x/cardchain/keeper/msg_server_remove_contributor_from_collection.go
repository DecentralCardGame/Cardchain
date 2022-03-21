package keeper

import (
	"context"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) RemoveContributorFromCollection(goCtx context.Context, msg *types.MsgRemoveContributorFromCollection) (*types.MsgRemoveContributorFromCollectionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	collection := k.Collections.Get(ctx, msg.CollectionId)
	if msg.Creator != collection.Contributors[0] {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Invalid creator")
	}
	if collection.Status != types.CStatus_design {
		return nil, types.ErrCollectionNotInDesign
	}

	newContributors, err := stringPopElementFromArr(msg.User, collection.Contributors)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "Contributor is not a contributor: "+msg.User)
	}

	collection.Contributors = newContributors

	k.Collections.Set(ctx, msg.CollectionId, collection)

	return &types.MsgRemoveContributorFromCollectionResponse{}, nil
}

func stringPopElementFromArr(element string, arr []string) ([]string, error) {
	for idx, val := range arr {
		if element == val {
			return append(arr[:idx], arr[idx+1:]...), nil
		}
	}
	return []string{}, types.ErrContributor
}
