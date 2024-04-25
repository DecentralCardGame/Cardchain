package keeper

import (
	"context"
	sdkerrors "cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) DisinviteEarlyAccess(goCtx context.Context, msg *types.MsgDisinviteEarlyAccess) (*types.MsgDisinviteEarlyAccessResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	inviter, err := k.GetUserFromString(ctx, msg.Creator)
	if err != nil {
		return nil, err
	}

	if inviter.EarlyAccess == nil || !inviter.EarlyAccess.Active || inviter.EarlyAccess.InvitedByUser != "" {
		return nil, sdkerrors.Wrapf(errors.ErrUnauthorized, "inviter has to be enroled in early access to invite users")
	}

	invited, err := k.GetUserFromString(ctx, msg.User)
	if err != nil {
		return nil, err
	}

	if invited.EarlyAccess == nil || !invited.EarlyAccess.Active || invited.EarlyAccess.InvitedByUser != msg.Creator || inviter.EarlyAccess.InvitedUser != msg.User {
		return nil, sdkerrors.Wrapf(errors.ErrUnauthorized, "invited user has not been invited to early access by the creator")
	}

	inviter.EarlyAccess.InvitedUser = ""
	invited.EarlyAccess.Active = false
	invited.EarlyAccess.InvitedByUser = ""

	k.SetUserFromUser(ctx, inviter)
	k.SetUserFromUser(ctx, invited)

	return &types.MsgDisinviteEarlyAccessResponse{}, nil
}
