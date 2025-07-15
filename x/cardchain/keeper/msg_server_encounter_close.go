package keeper

import (
	"context"
	"slices"

	errorsmod "cosmossdk.io/errors"
	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) EncounterClose(goCtx context.Context, msg *types.MsgEncounterClose) (*types.MsgEncounterCloseResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	reporter, err := k.GetUserFromString(ctx, msg.Creator)
	if err != nil {
		return nil, err
	}

	if !reporter.ReportMatches {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "unauthorized reporter")
	}

	user, err := k.GetUserFromString(ctx, msg.User)
	if err != nil {
		return nil, err
	}

	index := slices.Index(user.OpenEncounters, msg.EncounterId)

	if index == -1 {
		return nil, errorsmod.Wrapf(sdkerrors.ErrUnauthorized, "encounter %d isn't open for user", msg.EncounterId)
	}

	user.OpenEncounters = append(user.OpenEncounters[:index], user.OpenEncounters[index+1:]...)

	if msg.Won {
		user.WonEncounters = append(user.WonEncounters, msg.EncounterId)
		// TODO: Treasury reward here

		encounter := k.Encounterk.Get(ctx, msg.EncounterId)
		if !encounter.Proven {
			encounter.Proven = true
			k.Encounterk.Set(ctx, msg.EncounterId, encounter)
		}
	}

	k.SetUserFromUser(ctx, user)

	return &types.MsgEncounterCloseResponse{}, nil
}
