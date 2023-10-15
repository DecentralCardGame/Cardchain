package keeper

import (
	"context"
	"slices"

	sdkerrors "cosmossdk.io/errors"
	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) AddContributorToSet(goCtx context.Context, msg *types.MsgAddContributorToSet) (*types.MsgAddContributorToSetResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	set := k.Sets.Get(ctx, msg.SetId)
	err := checkSetEditable(set, msg.Creator)
	if err != nil {
		return nil, err
	}

	if slices.Contains(set.Contributors, msg.User) {
		return nil, sdkerrors.Wrap(types.ErrContributor, "Contributor allready Contributor: "+msg.User)
	}

	err = k.CollectSetConributionFee(ctx, msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(errors.ErrInsufficientFunds, err.Error())
	}

	set.Contributors = append(set.Contributors, msg.User)

	k.Sets.Set(ctx, msg.SetId, set)

	return &types.MsgAddContributorToSetResponse{}, nil
}
