package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"github.com/DecentralCardGame/cardchain/x/featureflag/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Set(goCtx context.Context, msg *types.MsgSet) (*types.MsgSetResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if k.authority != msg.Authority {
		return nil, errorsmod.Wrapf(types.ErrInvalidSigner, "expected %s got %s", k.authority, msg.Authority)
	}

	err := k.FlagExists(msg.Module, msg.Name)
	if err != nil {
		return nil, err
	}

	key := GetKey(msg.Module, msg.Name)

	flag := k.GetFlag(ctx, key)
	flag.Set = msg.Value
	k.SetFlag(ctx, key, flag)

	return &types.MsgSetResponse{}, nil
}
