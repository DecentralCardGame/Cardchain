package keeper

import (
	"context"
	"strconv"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) RemoveCardFromCollection(goCtx context.Context, msg *types.MsgRemoveCardFromCollection) (*types.MsgRemoveCardFromCollectionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	collection := k.GetCollection(ctx, msg.CollectionId)
	if !stringItemInList(msg.Creator, collection.Contributors) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Invalid contributor")
	}
	if collection.Status != types.CStatus_design {
		return nil, types.ErrCollectionNotInDesign
	}

	newCards, err := uintPopElementFromArr(msg.CardId, collection.Cards)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "Card: "+strconv.Itoa(int(msg.CardId)))
	}

	collection.Cards = newCards

	k.SetCollection(ctx, msg.CollectionId, collection)

	return &types.MsgRemoveCardFromCollectionResponse{}, nil
}

func uintPopElementFromArr(element uint64, arr []uint64) ([]uint64, error) {
	for idx, val := range arr {
		if element == val {
			return append(arr[:idx], arr[idx+1:]...), nil
		}
	}
	return []uint64{}, types.ErrCardNotInCollection
}
