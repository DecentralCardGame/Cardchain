package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) EncounterDelete(goCtx context.Context, msg *types.MsgEncounterDelete) (*types.MsgEncounterDeleteResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(err, "invalid authority address")
	}

	encounter := k.Encounterk.Get(ctx, msg.Id)

	if encounter.Owner != msg.Creator {
		return nil, errorsmod.Wrap(errors.ErrUnauthorized, "incorrect owner")

	}

	k.Images.Set(ctx, encounter.ImageId, nil)
	k.Encounterk.Set(ctx, msg.Id, nil)

	return &types.MsgEncounterDeleteResponse{}, nil
}
