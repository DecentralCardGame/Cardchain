package keeper

import (
	"context"
	"sort"

	errorsmod "cosmossdk.io/errors"
	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

type sortStruct struct {
	Id  uint64
	Set *types.Set
}

func (k msgServer) SetActivate(goCtx context.Context, msg *types.MsgSetActivate) (*types.MsgSetActivateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if k.authority != msg.Authority {
		return nil, errorsmod.Wrapf(types.ErrInvalidSigner, "expected %s got %s", k.authority, msg.Authority)
	}

	set := k.SetK.Get(ctx, msg.SetId)

	if set.Status != types.SetStatus_finalized {
		return nil, errorsmod.Wrapf(errors.ErrUnauthorized, "Set status is %s but should be finalized", set.Status)
	}

	activeSetsIds := k.GetActiveSets(ctx)
	var activeSets []sortStruct
	if len(activeSets) >= int(k.GetParams(ctx).ActiveSetsAmount) {
		for _, id := range activeSetsIds {
			var set = k.SetK.Get(ctx, id)
			activeSets = append(activeSets, sortStruct{id, set})
		}
		sort.SliceStable(activeSets, func(i, j int) bool {
			return activeSets[i].Set.TimeStamp < activeSets[j].Set.TimeStamp
		},
		)
		yeetStruct := activeSets[0]
		yeetStruct.Set.Status = types.SetStatus_archived
		k.SetK.Set(ctx, yeetStruct.Id, yeetStruct.Set)
	}

	set.Status = types.SetStatus_active
	set.TimeStamp = ctx.BlockHeight()

	k.SetK.Set(ctx, msg.SetId, set)

	return &types.MsgSetActivateResponse{}, nil
}
