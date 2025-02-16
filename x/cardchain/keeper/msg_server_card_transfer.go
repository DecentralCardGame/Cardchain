package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"github.com/DecentralCardGame/cardchain/util"
	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CardTransfer(goCtx context.Context, msg *types.MsgCardTransfer) (*types.MsgCardTransferResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// if the vote right is valid, get the Card
	card := k.cards.Get(ctx, msg.CardId)
	creator, err := k.GetUserFromString(ctx, msg.Creator)
	if err != nil {
		return nil, errorsmod.Wrap(types.ErrUserDoesNotExist, err.Error())
	}
	receiver, err := k.GetUserFromString(ctx, msg.Receiver)
	if err != nil {
		return nil, errorsmod.Wrap(types.ErrUserDoesNotExist, err.Error())
	}

	if card.Owner == "" {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "Card has no owner")
	} else if msg.Creator != card.Owner { // Checks if the the msg sender is the same as the current owner
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner") // If not, throw an error
	}

	if card.Status == types.CardStatus_scheme {
		creator.OwnedCardSchemes, err = util.PopItemFromArr(msg.CardId, creator.OwnedCardSchemes)
		if err != nil {
			return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, err.Error())
		}
		receiver.OwnedCardSchemes = append(receiver.OwnedCardSchemes, msg.CardId)
	} else {
		creator.OwnedPrototypes, err = util.PopItemFromArr(msg.CardId, creator.OwnedPrototypes)
		if err != nil {
			return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, err.Error())
		}
		receiver.OwnedPrototypes = append(receiver.OwnedPrototypes, msg.CardId)
	}

	card.Owner = msg.Receiver
	k.cards.Set(ctx, msg.CardId, card)
	k.SetUserFromUser(ctx, creator)
	k.SetUserFromUser(ctx, receiver)

	return &types.MsgCardTransferResponse{}, nil
}
