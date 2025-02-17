package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CouncilRegister(goCtx context.Context, msg *types.MsgCouncilRegister) (*types.MsgCouncilRegisterResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	user, err := k.GetUserFromString(ctx, msg.Creator)
	if err != nil {
		return nil, errorsmod.Wrap(types.ErrUserDoesNotExist, err.Error())
	}

	if user.CouncilStatus != types.CouncilStatus_unavailable {
		return nil, errorsmod.Wrapf(sdkerrors.ErrUnauthorized, "invalid council status %s", user.CouncilStatus.String())
	}

	iter := k.councils.GetItemIterator(ctx)
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
					return nil, errorsmod.Wrap(types.ErrUserDoesNotExist, err.Error())
				}
				if len(council.Voters) == 5 {
					usr.CouncilStatus = types.CouncilStatus_startedCouncil
				} else {
					usr.CouncilStatus = types.CouncilStatus_openCouncil
				}
				k.SetUserFromUser(ctx, usr)
			}
			k.councils.Set(ctx, idx, council)
			break
		}
	}

	user, err = k.GetUserFromString(ctx, msg.Creator)
	if err != nil {
		return nil, errorsmod.Wrap(types.ErrUserDoesNotExist, err.Error())
	}
	if user.CouncilStatus == types.CouncilStatus_unavailable {
		user.CouncilStatus = types.CouncilStatus_available
	}
	k.SetUserFromUser(ctx, user)

	return &types.MsgCouncilRegisterResponse{}, nil
}
