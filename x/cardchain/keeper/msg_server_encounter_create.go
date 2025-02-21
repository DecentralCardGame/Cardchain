package keeper

import (
	"context"
	"slices"

	errorsmod "cosmossdk.io/errors"
	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	"github.com/DecentralCardGame/cardobject/cardobject"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) EncounterCreate(goCtx context.Context, msg *types.MsgEncounterCreate) (*types.MsgEncounterCreateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	creator, err := k.GetUserFromString(ctx, msg.Creator)
	if err != nil {
		return nil, err
	}

	var (
		id, imageId uint64
		override    bool = false
	)

	iter := k.Encounterk.GetItemIterator(ctx)
	for ; iter.Valid(); iter.Next() {
		encounterId, encounter := iter.Value()

		if encounter.Name == msg.Name {
			if encounter.Owner != msg.Creator {
				return nil, errorsmod.Wrapf(errors.ErrUnauthorized, "encounter with same name already exists and is owned by '%s'", encounter.Owner)
			}
			id = encounterId
			imageId = encounter.ImageId
			override = true
		}
	}

	if !override {
		id = k.Encounterk.GetNum(ctx)
		imageId = k.Images.GetNum(ctx)
	}

	err = k.validateDrawlist(ctx, msg, &creator)
	if err != nil {
		return nil, err
	}

	encounter := types.Encounter{
		Id:         id,
		Drawlist:   msg.Drawlist,
		Name:       msg.Name,
		Proven:     false,
		Owner:      msg.Creator,
		Parameters: msg.Parameters,
		ImageId:    imageId,
	}

	k.Images.Set(ctx, imageId, &types.Image{Image: msg.Image})
	k.Encounterk.Set(ctx, id, &encounter)
	k.SetUserFromUser(ctx, creator)
	return &types.MsgEncounterCreateResponse{}, nil
}

func (k Keeper) validateDrawlist(ctx sdk.Context, msg *types.MsgEncounterCreate, creator *User) error {
	for idx, cardId := range msg.Drawlist {
		card := k.CardK.Get(ctx, cardId)

		if card.Owner != msg.Creator {
			index := slices.Index(creator.Cards, cardId)
			if index != -1 {
				creator.Cards = append(creator.Cards[:index], creator.Cards[index+1:]...)
			} else {
				return errorsmod.Wrapf(
					sdkerrors.ErrUnauthorized,
					"creator has to own all cards, doesnt own '%d'", cardId,
				)
			}
		}

		cardObj, err := card.GetCardObj()
		if err != nil {
			return err
		}

		if idx == 0 {
			if cardObj.GetType() != cardobject.HEADQUARTERTYPE {
				return errorsmod.Wrapf(
					types.ErrInvalidData,
					"first card has to be Headquarter but is: %s", cardObj.GetType(),
				)
			}
		} else if cardObj.GetType() == cardobject.HEADQUARTERTYPE {
			return errorsmod.Wrapf(
				types.ErrInvalidData,
				"only first card can be headquartter but card-%d is ", idx,
			)
		}
	}
	return nil
}
