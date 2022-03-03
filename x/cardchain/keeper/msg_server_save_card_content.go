package keeper

import (
	"context"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) SaveCardContent(goCtx context.Context, msg *types.MsgSaveCardContent) (*types.MsgSaveCardContentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	card := k.GetCard(ctx, msg.CardId)

	msgOwner, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidAccAddress, err.Error())
	}

	// TODO Add error when writing to unowned card

	cardOwner, err := sdk.AccAddressFromBech32(card.Owner)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidAccAddress, err.Error())
	}

	if !msgOwner.Equals(cardOwner) { // Checks if the the msg sender is the same as the current owner
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner")
	}

	// TODO cards get a starting pool currently, this should be removed later and the starting pool should come after council decision
	card.VotePool.Add(sdk.NewInt64Coin("ucredits", 10000000))

	card.Content = []byte(msg.Content)
	// card.Image = []byte(msg.Image)
	card.Status = types.Status_prototype
	card.Notes = msg.Notes
	card.Artist = msg.Artist
	// card.FullArt = msg.FullArt
	k.SetCard(ctx, msg.CardId, card)
	k.TransferSchemeToCard(ctx, msg.CardId, msgOwner)

	return &types.MsgSaveCardContentResponse{}, nil
}
