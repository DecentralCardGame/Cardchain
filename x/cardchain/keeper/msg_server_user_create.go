package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) UserCreate(goCtx context.Context, msg *types.MsgUserCreate) (*types.MsgUserCreateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := types.CheckAliasLimit(msg.Alias); err != nil {
		return nil, err
	}

	user, err := k.GetUserFromString(ctx, msg.NewUser)
	if err != nil {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if user.Alias != "" {
		return nil, errorsmod.Wrapf(sdkerrors.ErrUnauthorized, "user already initialized")
	}

	err = k.InitUser(ctx, user.Addr, msg.Alias)

	return &types.MsgUserCreateResponse{}, err
}
