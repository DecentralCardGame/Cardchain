package keeper

import (
	"context"
	"slices"

	errorsmod "cosmossdk.io/errors"
	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) SetContributorAdd(goCtx context.Context, msg *types.MsgSetContributorAdd) (*types.MsgSetContributorAddResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	set := k.Sets.Get(ctx, msg.SetId)
	if err := checkSetEditable(set, msg.Creator); err != nil {
		return nil, err
	}

	if slices.Contains(set.Contributors, msg.User) {
		return nil, errorsmod.Wrap(types.ErrInvalidData, "Contributor allready Contributor: "+msg.User)
	}

	if err := k.CollectSetConributionFee(ctx, msg.Creator); err != nil {
		return nil, errorsmod.Wrap(errors.ErrInsufficientFunds, err.Error())
	}

	set.Contributors = append(set.Contributors, msg.User)

	k.Sets.Set(ctx, msg.SetId, set)

	return &types.MsgSetContributorAddResponse{}, nil
}
