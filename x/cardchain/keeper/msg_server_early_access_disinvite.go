package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) EarlyAccessDisinvite(goCtx context.Context, msg *types.MsgEarlyAccessDisinvite) (*types.MsgEarlyAccessDisinviteResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	inviter, err := k.GetUserFromString(ctx, msg.Creator)
	if err != nil {
		return nil, err
	}

	if inviter.EarlyAccess == nil || !inviter.EarlyAccess.Active || inviter.EarlyAccess.InvitedByUser != "" {
		return nil, errorsmod.Wrapf(sdkerrors.ErrUnauthorized, "inviter has to be enroled in early access to invite users")
	}

	invited, err := k.GetUserFromString(ctx, msg.User)
	if err != nil {
		return nil, err
	}

	if invited.EarlyAccess == nil || !invited.EarlyAccess.Active || invited.EarlyAccess.InvitedByUser != msg.Creator || inviter.EarlyAccess.InvitedUser != msg.User {
		return nil, errorsmod.Wrapf(sdkerrors.ErrUnauthorized, "invited user has not been invited to early access by the creator")
	}

	inviter.EarlyAccess.InvitedUser = ""
	invited.RemoveEarlyAccess()

	k.SetUserFromUser(ctx, inviter)
	k.SetUserFromUser(ctx, invited)
	return &types.MsgEarlyAccessDisinviteResponse{}, nil
}
