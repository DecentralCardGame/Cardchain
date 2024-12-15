package keeper

import (
	"context"
	"slices"

	sdkerrors "cosmossdk.io/errors"
	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) EncounterClose(goCtx context.Context, msg *types.MsgEncounterClose) (*types.MsgEncounterCloseResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	reporter, err := k.GetMsgCreator(ctx, msg)
	if err != nil {
		return nil, err
	}

	if !reporter.ReportMatches {
		return nil, sdkerrors.Wrap(errors.ErrUnauthorized, "unauthorized reporter")
	}

	user, err := k.GetUserFromString(ctx, msg.User)
	if err != nil {
		return nil, err
	}

	index := slices.Index(user.OpenEncounters, msg.EncounterId)

	if index == -1 {
		return nil, sdkerrors.Wrapf(errors.ErrUnauthorized, "encounter %d isn't open for user", msg.EncounterId)
	}

	user.OpenEncounters = append(user.OpenEncounters[:index], user.OpenEncounters[index+1:]...)

	if msg.Won {
		user.WonEncounters = append(user.WonEncounters, msg.EncounterId)
		// TODO: Treasury reward here

		encounter := k.Encounters.Get(ctx, msg.EncounterId)
		if !encounter.Proven {
			encounter.Proven = true
			k.Encounters.Set(ctx, msg.EncounterId, encounter)
		}
	}

	k.SetUserFromUser(ctx, user)

	return &types.MsgEncounterCloseResponse{}, nil
}
