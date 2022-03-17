package keeper

import (
	"context"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) RestartCouncil(goCtx context.Context, msg *types.MsgRestartCouncil) (*types.MsgRestartCouncilResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	council := k.GetCouncil(ctx, msg.CouncilId)
	card := k.Card.Get(ctx, council.CardId)
	if card.Owner != msg.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Invalid Owner")
	}

	if council.Status != types.CouncelingStatus_suggestionsMade {
		return nil, sdkerrors.Wrapf(types.ErrCouncilStatus, "Have '%s', want '%s'", council.Status.String(), types.CouncelingStatus_suggestionsMade.String())
	}

	council.ClearResponses = []*types.WrapClearResponse{}
	council.HashResponses = []*types.WrapHashResponse{}
	council.Status = types.CouncelingStatus_councilCreated

	k.SetCouncil(ctx, msg.CouncilId, council)

	return &types.MsgRestartCouncilResponse{}, nil
}
