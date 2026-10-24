package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const BIO_MAX_LENGTH = 400

func (k msgServer) ProfileBioSet(goCtx context.Context, msg *types.MsgProfileBioSet) (*types.MsgProfileBioSetResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if len(msg.Bio) > BIO_MAX_LENGTH {
		return nil, errorsmod.Wrapf(types.ErrInvalidData, "Website length exceded %d chars", BIO_MAX_LENGTH)
	}

	user, err := k.GetUserFromString(ctx, msg.Creator)
	if err != nil {
		return nil, err
	}

	user.Biography = msg.Bio

	k.SetUserFromUser(ctx, user)

	return &types.MsgProfileBioSetResponse{}, nil
}
