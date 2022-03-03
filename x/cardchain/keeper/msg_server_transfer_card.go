package keeper

import (
	"context"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) TransferCard(goCtx context.Context, msg *types.MsgTransferCard) (*types.MsgTransferCardResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// if the vote right is valid, get the Card
	card := k.GetCard(ctx, msg.CardId)

	if card.Owner == "" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Card has no owner")
	} else if msg.Creator != card.Owner { // Checks if the the msg sender is the same as the current owner
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner") // If not, throw an error
	}

	card.Owner = msg.Receiver
	k.SetCard(ctx, msg.CardId, card)

	return &types.MsgTransferCardResponse{}, nil
}
