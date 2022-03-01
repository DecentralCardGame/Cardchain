package keeper

import (
	"context"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) RegisterForCouncil(goCtx context.Context, msg *types.MsgRegisterForCouncil) (*types.MsgRegisterForCouncilResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	user, err := k.GetUserFromString(ctx, msg.Creator)
	if err != nil {
		return nil, err
	}

	if user.CouncilStatus != types.CouncilStatus_unavailable {
		return nil, sdkerrors.Wrapf(types.ErrInvalidUserStatus, "%s", user.CouncilStatus.String())
	}

	allCouncils := k.GetAllCouncils(ctx)
	for i := len(allCouncils) - 1; i >= 0; i-- {
		var council = allCouncils[i]
		if council.Status == types.CouncelingStatus_councilOpen {
			council.Voters = append(council.Voters, msg.Creator)
			if len(council.Voters) == 5 {
				council.Status = types.CouncelingStatus_councilCreated
			}
			for _, addr := range council.Voters {
				usr, err := k.GetUserFromString(ctx, addr)
				if err != nil {
					return nil, err
				}
				if len(council.Voters) == 5 {
					usr.CouncilStatus = types.CouncilStatus_startedCouncil
				} else {
					usr.CouncilStatus = types.CouncilStatus_openCouncil
				}
				k.SetUserFromUser(ctx, usr)
			}
			k.SetCouncil(ctx, uint64(i), *council)
			break
		}
	}

	user, err = k.GetUserFromString(ctx, msg.Creator)
	if err != nil {
		return nil, err
	}
	if user.CouncilStatus == types.CouncilStatus_unavailable {
		user.CouncilStatus = types.CouncilStatus_available
	}
	k.SetUserFromUser(ctx, user)

	return &types.MsgRegisterForCouncilResponse{}, nil
}
