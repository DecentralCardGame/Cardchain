package keeper

import (
	"context"
	"encoding/json"
	"errors"

	"golang.org/x/exp/slices"
	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/keywords"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) SetCardRarity(goCtx context.Context, msg *types.MsgSetCardRarity) (*types.MsgSetCardRarityResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	card := k.Cards.Get(ctx, msg.CardId)
	collection := k.Collections.Get(ctx, msg.CollectionId)

	if collection.Contributors[0] != msg.Creator || !slices.Contains(collection.Cards, msg.CardId) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Creator")
	}

	cardobj, err := keywords.Unmarshal(card.Content)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrCardobject, err.Error())
	}

	rarity := cardobject.Rarity(msg.Rarity)
	if cardobj.Action != nil {
		cardobj.Action.Rarity = &rarity
	} else if cardobj.Place != nil {
		cardobj.Place.Rarity = &rarity
	} else if cardobj.Entity != nil {
		cardobj.Entity.Rarity = &rarity
	} else if cardobj.Headquarter != nil {
		cardobj.Headquarter.Rarity = &rarity
	} else {
		return nil, errors.New("No card-attributes")
	}

	cardbytes, err := json.Marshal(cardobj)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrCardobject, err.Error())
	}
	card.Content = cardbytes

	k.Cards.Set(ctx, msg.CardId, card)

	return &types.MsgSetCardRarityResponse{}, nil
}
