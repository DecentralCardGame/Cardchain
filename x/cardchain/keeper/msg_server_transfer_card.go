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

	sender, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidAccAddress, "Unable to convert to AccAddress")
	}

	owner, err := sdk.AccAddressFromBech32(card.Owner)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidAccAddress, "Unable to convert to AccAddress")
	}

	if !sender.Equals(owner) { // Checks if the the msg sender is the same as the current owner
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner") // If not, throw an error
	}

	card.Owner = msg.Receiver
	k.SetCard(ctx, msg.CardId, card)

	return &types.MsgTransferCardResponse{}, nil
}
