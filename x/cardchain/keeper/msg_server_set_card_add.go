package keeper

import (
	"context"
	"slices"

	errorsmod "cosmossdk.io/errors"
	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) SetCardAdd(goCtx context.Context, msg *types.MsgSetCardAdd) (*types.MsgSetCardAddResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	setSize := int(k.GetParams(ctx).SetSize)

	set := k.sets.Get(ctx, msg.SetId)
	if !slices.Contains(set.Contributors, msg.Creator) {
		return nil, errorsmod.Wrap(errors.ErrUnauthorized, "Invalid contributor")
	}
	if set.Status != types.CStatus_design {
		return nil, errorsmod.Wrapf(errors.ErrUnauthorized, "Invalid set status is: %s", set.Status.String())
	}

	iter := k.sets.GetItemIterator(ctx)
	for ; iter.Valid(); iter.Next() {
		idx, coll := iter.Value()
		if coll.Status != types.CStatus_archived && slices.Contains(coll.Cards, msg.CardId) {
			return nil, errorsmod.Wrapf(types.ErrCardAlreadyInSet, "Set: %d", idx)
		}
	}

	card := k.cards.Get(ctx, msg.CardId)
	if card.Status != types.CardStatus_permanent {
		return nil, errorsmod.Wrap(types.ErrCardDoesNotExist, "Card is not permanent or does not exist")
	}

	if card.Owner != msg.Creator {
		return nil, errorsmod.Wrap(errors.ErrUnauthorized, "Invalid creator")
	}

	if len(set.Cards) >= setSize {
		return nil, errorsmod.Wrapf(types.ErrInvalidData, "Max set size is %d", setSize)
	}

	if slices.Contains(set.Cards, msg.CardId) {
		return nil, errorsmod.Wrapf(types.ErrCardAlreadyInSet, "Card: %d", msg.CardId)
	}

	err := k.CollectSetConributionFee(ctx, msg.Creator)
	if err != nil {
		return nil, errorsmod.Wrap(errors.ErrInsufficientFunds, err.Error())
	}

	set.Cards = append(set.Cards, msg.CardId)

	k.sets.Set(ctx, msg.SetId, set)

	return &types.MsgSetCardAddResponse{}, nil
}
