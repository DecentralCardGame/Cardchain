package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	"github.com/DecentralCardGame/cardobject/cardobject"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) EncounterCreate(goCtx context.Context, msg *types.MsgEncounterCreate) (*types.MsgEncounterCreateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(err, "invalid authority address")
	}

	err := types.ValidateImage(msg)
	if err != nil {
		return nil, err
	}

	id := k.EncounterK.GetNum(ctx)
	imageId := k.Images.GetNum(ctx)

	encounter := types.Encounter{
		Id:         id,
		Drawlist:   msg.Drawlist,
		Name:       msg.Name,
		Proven:     false,
		Owner:      msg.Creator,
		Parameters: msg.Parameters,
		ImageId:    imageId,
	}

	err = k.validateEncounter(ctx, &encounter, msg.Creator)
	if err != nil {
		return nil, err
	}

	k.Images.Set(ctx, imageId, &types.Image{Image: msg.Image})
	k.EncounterK.Set(ctx, id, &encounter)
	return &types.MsgEncounterCreateResponse{}, nil
}

func (k Keeper) validateEncounter(ctx sdk.Context, encounter *types.Encounter, creator string) error {
	if len(encounter.Drawlist) > 40 || len(encounter.Drawlist) < 1 {
		return errorsmod.Wrapf(
			types.ErrInvalidData, "invalid drawlist length, max 40 is '%d'", len(encounter.Drawlist),
		)
	}

	if encounter.Name == "" {
		return errorsmod.Wrap(types.ErrInvalidData, "encounter needs a name")
	}

	iter := k.EncounterK.GetItemIterator(ctx)
	for ; iter.Valid(); iter.Next() {
		_, e := iter.Value()

		if e.Name == encounter.Name && e.Id != encounter.Id {
			return errorsmod.Wrapf(
				errors.ErrUnauthorized,
				"encounter with same name already exists and is owned by '%s'",
				e.Owner,
			)
		}
	}

	err := k.validateDrawlist(ctx, encounter, creator)
	if err != nil {
		return err
	}

	return nil
}

func (k Keeper) validateDrawlist(ctx sdk.Context, encounter *types.Encounter, creator string) error {
	for idx, cardId := range encounter.Drawlist {
		card := k.CardK.Get(ctx, cardId)

		if card.Owner != creator {
			return errorsmod.Wrapf(
				sdkerrors.ErrUnauthorized,
				"creator has to own all cards, doesnt own '%d'", cardId,
			)
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
