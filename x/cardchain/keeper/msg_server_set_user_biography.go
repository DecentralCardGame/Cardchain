package keeper

import (
	"context"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) SetUserBiography(goCtx context.Context, msg *types.MsgSetUserBiography) (*types.MsgSetUserBiographyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	user, err := k.GetUserFromString(ctx, msg.Creator)
	if err != nil {
		return nil, err
	}

	if len(msg.Biography) > 400 {
		return nil, sdkerrors.Wrap(types.ErrStringLength, "Website length exceded 400 chars")
	}

	user.Biography = msg.Biography

	k.SetUserFromUser(ctx, user)

	return &types.MsgSetUserBiographyResponse{}, nil
}
