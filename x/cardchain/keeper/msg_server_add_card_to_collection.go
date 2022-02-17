package keeper

import (
	"context"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) AddCardToCollection(goCtx context.Context, msg *types.MsgAddCardToCollection) (*types.MsgAddCardToCollectionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	collectionSize := int(k.GetParams(ctx).CollectionSize)

	collection := k.GetCollection(ctx, msg.CollectionId)
	if !stringItemInList(msg.Creator, collection.Contributors) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Invalid contributor")
	}
	if collection.Status != types.CStatus_design {
		return nil, types.ErrCollectionNotInDesign
	}

	card := k.GetCard(ctx, msg.CardId)
	if card.Status != types.Status_permanent {
		return nil, sdkerrors.Wrap(types.ErrCardDoesNotExist, "Card is not permanent or does not exist")
	}

	if card.Owner != msg.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Invalid creator")
	}

	if len(collection.Cards) >= collectionSize {
		return nil, sdkerrors.Wrapf(types.ErrCollectionSize, "Max is %d", collectionSize)
	}

	if uintItemInList(msg.CardId, collection.Cards) {
		return nil, sdkerrors.Wrapf(types.ErrCardAlreadyInCollection, "Card: %d", msg.CardId)
	}

	err := k.CollectCollectionConributionFee(ctx, msg.Creator)
	if err != nil {
		return nil, err
	}

	collection.Cards = append(collection.Cards, msg.CardId)

	k.SetCollection(ctx, msg.CollectionId, collection)

	return &types.MsgAddCardToCollectionResponse{}, nil
}
