package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const WEBSITE_MAX_LENGTH = 50

func (k msgServer) ProfileWebsiteSet(goCtx context.Context, msg *types.MsgProfileWebsiteSet) (*types.MsgProfileWebsiteSetResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if len(msg.Website) > WEBSITE_MAX_LENGTH {
		return nil, errorsmod.Wrapf(types.ErrInvalidData, "Website length exceded %d chars", WEBSITE_MAX_LENGTH)
	}

	user, err := k.GetUserFromString(ctx, msg.Creator)
	if err != nil {
		return nil, err
	}

	user.Website = msg.Website

	k.SetUserFromUser(ctx, user)

	return &types.MsgProfileWebsiteSetResponse{}, nil
}
