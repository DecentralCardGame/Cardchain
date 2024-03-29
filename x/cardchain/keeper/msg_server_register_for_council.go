package keeper

import (
	"context"

	sdkerrors "cosmossdk.io/errors"
	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) RegisterForCouncil(goCtx context.Context, msg *types.MsgRegisterForCouncil) (*types.MsgRegisterForCouncilResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	user, err := k.GetUserFromString(ctx, msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrUserDoesNotExist, err.Error())
	}

	if user.CouncilStatus != types.CouncilStatus_unavailable {
		return nil, sdkerrors.Wrapf(types.ErrInvalidUserStatus, "%s", user.CouncilStatus.String())
	}

	iter := k.Councils.GetItemIterator(ctx)
	for ; iter.Valid(); iter.Next() {
		idx, council := iter.Value()
		if council.Status == types.CouncelingStatus_councilOpen {
			council.Voters = append(council.Voters, msg.Creator)
			if len(council.Voters) == 5 {
				council.Status = types.CouncelingStatus_councilCreated
			}
			for _, addr := range council.Voters {
				usr, err := k.GetUserFromString(ctx, addr)
				if err != nil {
					return nil, sdkerrors.Wrap(types.ErrUserDoesNotExist, err.Error())
				}
				if len(council.Voters) == 5 {
					usr.CouncilStatus = types.CouncilStatus_startedCouncil
				} else {
					usr.CouncilStatus = types.CouncilStatus_openCouncil
				}
				k.SetUserFromUser(ctx, usr)
			}
			k.Councils.Set(ctx, idx, council)
			break
		}
	}

	user, err = k.GetUserFromString(ctx, msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrUserDoesNotExist, err.Error())
	}
	if user.CouncilStatus == types.CouncilStatus_unavailable {
		user.CouncilStatus = types.CouncilStatus_available
	}
	k.SetUserFromUser(ctx, user)

	return &types.MsgRegisterForCouncilResponse{}, nil
}
