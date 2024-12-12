package keeper

import (
	"context"

	sdkerrors "cosmossdk.io/errors"
	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/keywords"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/group/errors"
)

func (k msgServer) EncounterCreate(goCtx context.Context, msg *types.MsgEncounterCreate) (*types.MsgEncounterCreateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.validateDrawlist(ctx, msg)
	if err != nil {
		return nil, err
	}

	id := k.Encounters.GetNum(ctx)
	imageId := k.Images.GetNum(ctx)

	encounter := types.Encounter{
		Id:         id,
		Drawlist:   msg.Drawlist,
		Proven:     false,
		Owner:      msg.Creator,
		Parameters: msg.Parameters,
		ImageId:    imageId,
	}

	k.Images.Set(ctx, imageId, &types.Image{Image: msg.Image})
	k.Encounters.Set(ctx, id, &encounter)

	return &types.MsgEncounterCreateResponse{}, nil
}

func (k Keeper) validateDrawlist(ctx sdk.Context, msg *types.MsgEncounterCreate) error {
	for idx, cardId := range msg.Drawlist {
		card := k.Cards.Get(ctx, cardId)

		if card.Owner != msg.Creator {
			return sdkerrors.Wrapf(
				errors.ErrUnauthorized,
				"creator has to own all cards, doesnt own '%d'", cardId,
			)
		}

		cardObj, err := keywords.Unmarshal(card.Content)
		if err != nil {
			return err
		}

		if idx == 0 && cardObj.GetType() != cardobject.HEADQUARTERTYPE {
			return sdkerrors.Wrapf(
				types.ErrInvalidData,
				"first card has to be Headquarter but is: %s", cardObj.GetType(),
			)
		} else if cardObj.GetType() == cardobject.HEADQUARTERTYPE {
			return sdkerrors.Wrapf(
				types.ErrInvalidData,
				"only first card can be headquartter but card-%d is ", idx,
			)
		}
	}
	return nil
}
