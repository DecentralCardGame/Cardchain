package keeper

import (
	"context"
	"slices"

	sdkerrors "cosmossdk.io/errors"
	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) AddCardToSet(goCtx context.Context, msg *types.MsgAddCardToSet) (*types.MsgAddCardToSetResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	setSize := int(k.GetParams(ctx).SetSize)

	set := k.Sets.Get(ctx, msg.SetId)
	if !slices.Contains(set.Contributors, msg.Creator) {
		return nil, sdkerrors.Wrap(errors.ErrUnauthorized, "Invalid contributor")
	}
	if set.Status != types.CStatus_design {
		return nil, types.ErrSetNotInDesign
	}

	iter := k.Sets.GetItemIterator(ctx)
	for ; iter.Valid(); iter.Next() {
		idx, coll := iter.Value()
		if coll.Status != types.CStatus_archived && slices.Contains(coll.Cards, msg.CardId) {
			return nil, sdkerrors.Wrapf(types.ErrCardAlreadyInSet, "Set: %d", idx)
		}
	}

	card := k.Cards.Get(ctx, msg.CardId)
	if card.Status != types.Status_permanent {
		return nil, sdkerrors.Wrap(types.ErrCardDoesNotExist, "Card is not permanent or does not exist")
	}

	if card.Owner != msg.Creator {
		return nil, sdkerrors.Wrap(errors.ErrUnauthorized, "Invalid creator")
	}

	if len(set.Cards) >= setSize {
		return nil, sdkerrors.Wrapf(types.ErrSetSize, "Max is %d", setSize)
	}

	if slices.Contains(set.Cards, msg.CardId) {
		return nil, sdkerrors.Wrapf(types.ErrCardAlreadyInSet, "Card: %d", msg.CardId)
	}

	err := k.CollectSetConributionFee(ctx, msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(errors.ErrInsufficientFunds, err.Error())
	}

	set.Cards = append(set.Cards, msg.CardId)

	k.Sets.Set(ctx, msg.SetId, set)

	return &types.MsgAddCardToSetResponse{}, nil
}
