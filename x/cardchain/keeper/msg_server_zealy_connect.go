package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) ZealyConnect(goCtx context.Context, msg *types.MsgZealyConnect) (*types.MsgZealyConnectResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, err := k.GetUserFromString(ctx, msg.Creator)
	if err != nil {
		return nil, errorsmod.Wrap(types.ErrUserDoesNotExist, msg.Creator)
	}

	zealy := k.zealy.Get(ctx, msg.ZealyId)
	if zealy.ZealyId == msg.ZealyId {
		return nil, errorsmod.Wrapf(sdkerrors.ErrUnauthorized, "ZealyId `%s` is already registered", msg.ZealyId)
	}

	iterator := k.zealy.GetIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {

		var gotten types.Zealy
		k.cdc.MustUnmarshal(iterator.Value(), &gotten)

		if gotten.Address == msg.Creator {
			return nil, errorsmod.Wrapf(sdkerrors.ErrUnauthorized, "User `%s` has already registered a zealyId", msg.Creator)
		}
	}

	k.zealy.Set(ctx, msg.ZealyId, &types.Zealy{Address: msg.Creator, ZealyId: msg.ZealyId})

	return &types.MsgZealyConnectResponse{}, nil
}
