package keeper

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/keywords"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) SetCardRarity(goCtx context.Context, msg *types.MsgSetCardRarity) (*types.MsgSetCardRarityResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	card := k.GetCard(ctx, msg.CardId)
	collection := k.GetCollection(ctx, msg.CollectionId)

	if collection.Contributors[0] != msg.Creator || !UintItemInList(msg.CardId, collection.Cards) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Creator")
	}

	cardobj, err := keywords.Unmarshal(card.Content)
	if err != nil {
		return nil, err
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
		return nil, err
	}
	card.Content = cardbytes

	k.SetCard(ctx, msg.CardId, card)

	return &types.MsgSetCardRarityResponse{}, nil
}
