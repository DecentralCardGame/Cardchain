package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) EncounterEdit(goCtx context.Context, msg *types.MsgEncounterEdit) (*types.MsgEncounterEditResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(err, "invalid authority address")
	}

	encounter := k.Encounterk.Get(ctx, msg.Id)

	if encounter.Owner != msg.Creator {
		return nil, errorsmod.Wrapf(sdkerrors.ErrUnauthorized, "Invalid Owner, owned by '%s'", encounter.Owner)
	}

	newEncounter := types.Encounter{
		Id:         msg.Id,
		Drawlist:   msg.Drawlist,
		Name:       msg.Name,
		Proven:     false,
		Owner:      msg.Creator,
		Parameters: msg.Parameters,
		ImageId:    encounter.ImageId,
	}

	err := k.validateEncounter(ctx, &newEncounter, msg.Creator)
	if err != nil {
		return nil, err
	}

	k.Images.Set(ctx, encounter.ImageId, &types.Image{Image: msg.Image})
	k.Encounterk.Set(ctx, msg.Id, &newEncounter)

	return &types.MsgEncounterEditResponse{}, nil
}
