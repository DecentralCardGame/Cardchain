package keeper

import (
	"context"

	sdkerrors "cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) ConnectZealyAccount(goCtx context.Context, msg *types.MsgConnectZealyAccount) (*types.MsgConnectZealyAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, err := k.GetUserFromString(ctx, msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrUserDoesNotExist, msg.Creator)
	}

	_, err = k.GetZealy(ctx, msg.ZealyId)
	if err == nil {
		return nil, sdkerrors.Wrapf(errors.ErrUnauthorized, "ZealyId `%s` is already registered", msg.ZealyId)
	}

	iterator := k.GetZealyIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {

		var gotten types.Zealy
		k.cdc.MustUnmarshal(iterator.Value(), &gotten)

		if gotten.Address == msg.Creator {
			return nil, sdkerrors.Wrapf(errors.ErrUnauthorized, "User `%s` has already registered a zealyId", msg.Creator)
		}
	}

	k.SetZealy(ctx, msg.ZealyId, types.Zealy{Address: msg.Creator, ZealyId: msg.ZealyId})

	return &types.MsgConnectZealyAccountResponse{}, nil
}
