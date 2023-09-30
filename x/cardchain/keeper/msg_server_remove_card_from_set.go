package keeper

import (
	"context"

	sdkerrors "cosmossdk.io/errors"
	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"golang.org/x/exp/slices"
)

func (k msgServer) RemoveCardFromSet(goCtx context.Context, msg *types.MsgRemoveCardFromSet) (*types.MsgRemoveCardFromSetResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	set := k.Sets.Get(ctx, msg.SetId)
	if !slices.Contains(set.Contributors, msg.Creator) {
		return nil, sdkerrors.Wrap(errors.ErrUnauthorized, "Invalid contributor")
	}
	if set.Status != types.CStatus_design {
		return nil, types.ErrSetNotInDesign
	}

	newCards, err := PopItemFromArr(msg.CardId, set.Cards)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "Card: %d", msg.CardId)
	}

	set.Cards = newCards

	k.Sets.Set(ctx, msg.SetId, set)

	return &types.MsgRemoveCardFromSetResponse{}, nil
}
