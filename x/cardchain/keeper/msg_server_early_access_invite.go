package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) EarlyAccessInvite(goCtx context.Context, msg *types.MsgEarlyAccessInvite) (*types.MsgEarlyAccessInviteResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	inviter, err := k.GetUserFromString(ctx, msg.Creator)
	if err != nil {
		return nil, err
	}

	if inviter.EarlyAccess == nil || !inviter.EarlyAccess.Active || inviter.EarlyAccess.InvitedByUser != "" {
		return nil, errorsmod.Wrapf(sdkerrors.ErrUnauthorized, "inviter has to be enroled in early access to invite users")
	}

	if inviter.EarlyAccess.InvitedUser != "" {
		oldInvitee, _ := k.GetUserFromString(ctx, inviter.EarlyAccess.InvitedUser)
		oldInvitee.RemoveEarlyAccess()
		k.SetUserFromUser(ctx, oldInvitee)
	}

	invited, err := k.GetUserFromString(ctx, msg.User)
	if err != nil {
		return nil, err
	}

	if invited.EarlyAccess != nil && invited.EarlyAccess.Active {
		return nil, errorsmod.Wrapf(sdkerrors.ErrUnauthorized, "invited user is already enroled in early access")
	}

	invited.AddEarlyAccess(msg.Creator)
	inviter.EarlyAccess.InvitedUser = msg.User

	k.SetUserFromUser(ctx, inviter)
	k.SetUserFromUser(ctx, invited)
	return &types.MsgEarlyAccessInviteResponse{}, nil
}
