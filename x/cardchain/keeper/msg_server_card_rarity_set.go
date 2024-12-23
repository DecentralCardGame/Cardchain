package keeper

import (
	"context"
	"slices"

	errorsmod "cosmossdk.io/errors"
	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CardRaritySet(goCtx context.Context, msg *types.MsgCardRaritySet) (*types.MsgCardRaritySetResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	card := k.Cards.Get(ctx, msg.CardId)
	set := k.Sets.Get(ctx, msg.SetId)

	if set.Contributors[0] != msg.Creator || !slices.Contains(set.Cards, msg.CardId) {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Creator")
	}

	card.Rarity = msg.Rarity
	k.Cards.Set(ctx, msg.CardId, card)

	return &types.MsgCardRaritySetResponse{}, nil
}
