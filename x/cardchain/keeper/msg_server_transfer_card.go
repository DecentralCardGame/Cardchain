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
	creator, err := k.GetUserFromString(ctx, msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrUserDoesNotExist, err.Error())
	}
	receiver, err := k.GetUserFromString(ctx, msg.Receiver)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrUserDoesNotExist, err.Error())
	}

	if card.Owner == "" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Card has no owner")
	} else if msg.Creator != card.Owner { // Checks if the the msg sender is the same as the current owner
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner") // If not, throw an error
	}

	if card.Status == types.Status_scheme {
		creator.OwnedCardSchemes, err = UintPopItemFromArr(msg.CardId, creator.OwnedCardSchemes)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, err.Error())
		}
		receiver.OwnedCardSchemes = append(receiver.OwnedCardSchemes, msg.CardId)
	} else {
		creator.OwnedPrototypes, err = UintPopItemFromArr(msg.CardId, creator.OwnedPrototypes)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, err.Error())
		}
		receiver.OwnedPrototypes = append(receiver.OwnedPrototypes, msg.CardId)
	}

	card.Owner = msg.Receiver
	k.SetCard(ctx, msg.CardId, card)
	k.SetUserFromUser(ctx, creator)
	k.SetUserFromUser(ctx, receiver)

	return &types.MsgTransferCardResponse{}, nil
}
