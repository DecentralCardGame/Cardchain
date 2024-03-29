package keeper

import (
	"context"

	sdkerrors "cosmossdk.io/errors"
	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) RestartCouncil(goCtx context.Context, msg *types.MsgRestartCouncil) (*types.MsgRestartCouncilResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	council := k.Councils.Get(ctx, msg.CouncilId)
	card := k.Cards.Get(ctx, council.CardId)
	if card.Owner != msg.Creator {
		return nil, sdkerrors.Wrap(errors.ErrUnauthorized, "Invalid Owner")
	}

	if council.Status != types.CouncelingStatus_suggestionsMade {
		return nil, sdkerrors.Wrapf(types.ErrCouncilStatus, "Have '%s', want '%s'", council.Status.String(), types.CouncelingStatus_suggestionsMade.String())
	}

	council.ClearResponses = []*types.WrapClearResponse{}
	council.HashResponses = []*types.WrapHashResponse{}
	council.Status = types.CouncelingStatus_councilCreated

	k.Councils.Set(ctx, msg.CouncilId, council)

	return &types.MsgRestartCouncilResponse{}, nil
}
