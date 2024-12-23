package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CouncilDeregister(goCtx context.Context, msg *types.MsgCouncilDeregister) (*types.MsgCouncilDeregisterResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	user, err := k.GetUserFromString(ctx, msg.Creator)
	if err != nil {
		return nil, errorsmod.Wrap(types.ErrUserDoesNotExist, err.Error())
	}

	if user.CouncilStatus != types.CouncilStatus_available {
		return nil, errorsmod.Wrapf(sdkerrors.ErrUnauthorized, "%s", user.CouncilStatus.String())
	}

	user.CouncilStatus = types.CouncilStatus_unavailable

	k.SetUserFromUser(ctx, user)

	return &types.MsgCouncilDeregisterResponse{}, nil
}
