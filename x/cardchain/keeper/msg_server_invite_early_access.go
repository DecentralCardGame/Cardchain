package keeper

import (
	"context"
	sdkerrors "cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) InviteEarlyAccess(goCtx context.Context, msg *types.MsgInviteEarlyAccess) (*types.MsgInviteEarlyAccessResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	inviter, err := k.GetUserFromString(ctx, msg.Creator)
	if err != nil {
		return nil, err
	}

	if inviter.EarlyAccess == nil || !inviter.EarlyAccess.Active || len(inviter.EarlyAccess.InvitedUser) != 0 {
		return nil, sdkerrors.Wrapf(errors.ErrUnauthorized, "inviter has to be enroled in early access to invite users")
	}

	invited, err := k.GetUserFromString(ctx, msg.User)
	if err != nil {
		return nil, err
	}

	if invited.EarlyAccess != nil && invited.EarlyAccess.Active {
		return nil, sdkerrors.Wrapf(errors.ErrUnauthorized, "invited user is already enroled in early access")
	}

	invited.EarlyAccess = &types.EarlyAccess{
		Active:        true,
		InvitedByUser: msg.Creator,
	}

	inviter.EarlyAccess.InvitedUser = msg.User

	k.SetUserFromUser(ctx, inviter)
	k.SetUserFromUser(ctx, invited)

	return &types.MsgInviteEarlyAccessResponse{}, nil
}
