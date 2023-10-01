package keeper

import (
	"context"

	sdkerrors "cosmossdk.io/errors"
	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors2 "github.com/cosmos/cosmos-sdk/types/errors"
	"golang.org/x/exp/slices"
)

func (k msgServer) SetCardRarity(goCtx context.Context, msg *types.MsgSetCardRarity) (*types.MsgSetCardRarityResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	card := k.Cards.Get(ctx, msg.CardId)
	set := k.Sets.Get(ctx, msg.SetId)

	if set.Contributors[0] != msg.Creator || !slices.Contains(set.Cards, msg.CardId) {
		return nil, sdkerrors.Wrap(sdkerrors2.ErrUnauthorized, "Incorrect Creator")
	}

	card.Rarity = msg.Rarity
	k.Cards.Set(ctx, msg.CardId, card)

	return &types.MsgSetCardRarityResponse{}, nil
}
