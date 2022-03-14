package keeper

import (
	"context"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) RewokeCouncilRegistration(goCtx context.Context, msg *types.MsgRewokeCouncilRegistration) (*types.MsgRewokeCouncilRegistrationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	user, err := k.GetUserFromString(ctx, msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrUserDoesNotExist, err.Error())
	}

	if user.CouncilStatus != types.CouncilStatus_available {
		return nil, sdkerrors.Wrapf(types.ErrInvalidUserStatus, "%s", user.CouncilStatus.String())
	}

	user.CouncilStatus = types.CouncilStatus_unavailable

	k.SetUserFromUser(ctx, user)

	return &types.MsgRewokeCouncilRegistrationResponse{}, nil
}
