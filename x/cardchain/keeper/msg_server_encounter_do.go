package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) EncounterDo(goCtx context.Context, msg *types.MsgEncounterDo) (*types.MsgEncounterDoResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	reporter, err := k.GetUserFromString(ctx, msg.Creator)
	if err != nil {
		return nil, err
	}

	if !reporter.ReportMatches {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "unauthorized reporter")
	}

	maxId := k.Encounters.GetNum(ctx)
	if msg.EncounterId >= maxId {
		return nil, errorsmod.Wrap(types.ErrInvalidData, "encounter doesnt exist")
	}

	// TODO: Treasury fee here

	user, err := k.GetUserFromString(ctx, msg.User)
	if err != nil {
		return nil, err
	}

	user.OpenEncounters = append(user.OpenEncounters, msg.EncounterId)

	k.SetUserFromUser(ctx, user)

	return &types.MsgEncounterDoResponse{}, nil
}
