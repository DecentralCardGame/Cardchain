package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CouncilRestart(goCtx context.Context, msg *types.MsgCouncilRestart) (*types.MsgCouncilRestartResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	council := k.Councils.Get(ctx, msg.CouncilId)
	card := k.CardK.Get(ctx, council.CardId)
	if card.Owner != msg.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "Invalid Owner")
	}

	if council.Status != types.CouncelingStatus_suggestionsMade {
		return nil, errorsmod.Wrapf(sdkerrors.ErrUnauthorized, "invalid status, have '%s', want '%s'", council.Status.String(), types.CouncelingStatus_suggestionsMade.String())
	}

	council.ClearResponses = []*types.WrapClearResponse{}
	council.HashResponses = []*types.WrapHashResponse{}
	council.Status = types.CouncelingStatus_councilCreated

	k.Councils.Set(ctx, msg.CouncilId, council)
	return &types.MsgCouncilRestartResponse{}, nil
}
