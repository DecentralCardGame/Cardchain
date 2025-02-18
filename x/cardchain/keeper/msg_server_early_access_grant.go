package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) EarlyAccessGrant(goCtx context.Context, msg *types.MsgEarlyAccessGrant) (*types.MsgEarlyAccessGrantResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if k.authority != msg.Authority {
		return nil, errorsmod.Wrapf(types.ErrInvalidSigner, "expected %s got %s", k.authority, msg.Authority)
	}

	for _, addr := range msg.Users {
		user, err := k.GetUserFromString(ctx, addr)
		if err != nil {
			return nil, err
		}

		user.AddEarlyAccess(msg.Authority)

		k.SetUserFromUser(ctx, user)
	}
	return &types.MsgEarlyAccessGrantResponse{}, nil
}
