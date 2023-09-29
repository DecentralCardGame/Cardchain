package keeper

import (
	"context"

	sdkerrors "cosmossdk.io/errors"
	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"golang.org/x/exp/slices"
)

func (k msgServer) AddContributorToCollection(goCtx context.Context, msg *types.MsgAddContributorToCollection) (*types.MsgAddContributorToCollectionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	collection := k.Collections.Get(ctx, msg.CollectionId)
	err := checkCollectionEditable(collection, msg.Creator)
	if err != nil {
		return nil, err
	}

	if slices.Contains(collection.Contributors, msg.User) {
		return nil, sdkerrors.Wrap(types.ErrContributor, "Contributor allready Contributor: "+msg.User)
	}

	err = k.CollectCollectionConributionFee(ctx, msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(errors.ErrInsufficientFunds, err.Error())
	}

	collection.Contributors = append(collection.Contributors, msg.User)

	k.Collections.Set(ctx, msg.CollectionId, collection)

	return &types.MsgAddContributorToCollectionResponse{}, nil
}
