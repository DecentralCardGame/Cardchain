package keeper

import (
	"context"

	sdkerrors "cosmossdk.io/errors"
	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/group/errors"
)

func (k msgServer) EncounterDo(goCtx context.Context, msg *types.MsgEncounterDo) (*types.MsgEncounterDoResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	reporter, err := k.GetMsgCreator(ctx, msg)
	if err != nil {
		return nil, err
	}

	if !reporter.ReportMatches {
		return nil, sdkerrors.Wrap(errors.ErrUnauthorized, "unauthorized reporter")
	}

	maxId := k.Encounters.GetNum(ctx)
	if msg.EncounterId >= maxId {
		return nil, sdkerrors.Wrap(types.ErrInvalidData, "encounter doesnt exist")
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
