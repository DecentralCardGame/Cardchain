package keeper

import (
	"context"
	"slices"

	errorsmod "cosmossdk.io/errors"
	"github.com/DecentralCardGame/cardchain/util"
	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) SetCardRemove(goCtx context.Context, msg *types.MsgSetCardRemove) (*types.MsgSetCardRemoveResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	set := k.sets.Get(ctx, msg.SetId)
	if !slices.Contains(set.Contributors, msg.Creator) {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "Invalid contributor")
	}
	if set.Status != types.SetStatus_design {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "set not in design")
	}

	newCards, err := util.PopItemFromArr(msg.CardId, set.Cards)
	if err != nil {
		return nil, errorsmod.Wrapf(err, "Card: %d", msg.CardId)
	}

	set.Cards = newCards

	k.sets.Set(ctx, msg.SetId, set)
	return &types.MsgSetCardRemoveResponse{}, nil
}
