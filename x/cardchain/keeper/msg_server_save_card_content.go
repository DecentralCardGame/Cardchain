package keeper

import (
	"context"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) SaveCardContent(goCtx context.Context, msg *types.MsgSaveCardContent) (*types.MsgSaveCardContentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	card := k.Card.Get(ctx, msg.CardId)

	msgOwner, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidAccAddress, err.Error())
	}

	if card.Owner == "" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Card has no owner")
	} else if msg.Creator != card.Owner { // Checks if the the msg sender is the same as the current owner
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner")
	}

	card.Content = []byte(msg.Content)
	card.Status = types.Status_prototype
	card.Notes = msg.Notes
	card.Artist = msg.Artist
	k.Card.Set(ctx, msg.CardId, card)
	err = k.TransferSchemeToCard(ctx, msg.CardId, msgOwner)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "An error accured while converting a card to a scheme: "+err.Error())
	}

	return &types.MsgSaveCardContentResponse{}, nil
}
