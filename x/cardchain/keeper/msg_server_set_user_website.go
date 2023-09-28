package keeper

import (
	"context"

	sdkerrors "cosmossdk.io/errors"
	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SetUserWebsite(goCtx context.Context, msg *types.MsgSetUserWebsite) (*types.MsgSetUserWebsiteResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	user, err := k.GetUserFromString(ctx, msg.Creator)
	if err != nil {
		return nil, err
	}

	if len(msg.Website) > 50 {
		return nil, sdkerrors.Wrap(types.ErrStringLength, "Website length exceded 50 chars")
	}

	user.Website = msg.Website

	k.SetUserFromUser(ctx, user)

	return &types.MsgSetUserWebsiteResponse{}, nil
}
