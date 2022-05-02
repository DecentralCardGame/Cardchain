package keeper

import (
	"context"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"golang.org/x/exp/slices"
)

func (k msgServer) AddContributorToCollection(goCtx context.Context, msg *types.MsgAddContributorToCollection) (*types.MsgAddContributorToCollectionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	collection := k.Collections.Get(ctx, msg.CollectionId)
	if msg.Creator != collection.Contributors[0] {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Invalid creator")
	}
	if collection.Status != types.CStatus_design {
		return nil, types.ErrCollectionNotInDesign
	}

	if slices.Contains(collection.Contributors, msg.User) {
		return nil, sdkerrors.Wrap(types.ErrContributor, "Contributor allready Contributor: "+msg.User)
	}

	err := k.CollectCollectionConributionFee(ctx, msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, err.Error())
	}

	collection.Contributors = append(collection.Contributors, msg.User)

	k.Collections.Set(ctx, msg.CollectionId, collection)

	return &types.MsgAddContributorToCollectionResponse{}, nil
}
