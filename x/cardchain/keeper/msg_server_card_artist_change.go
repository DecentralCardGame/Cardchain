package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CardArtistChange(goCtx context.Context, msg *types.MsgCardArtistChange) (*types.MsgCardArtistChangeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	card := k.Cards.Get(ctx, msg.CardId)

	if card.Status != types.Status_prototype {
		return nil, errorsmod.Wrap(types.ErrInvalidCardStatus, "Card has to be a prototype to be changeable")
	}

	if card.Owner != msg.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner")
	}

	newArtist, err := sdk.AccAddressFromBech32(msg.Artist)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "Invalid Artist address")
	}

	card.Artist = newArtist.String()

	k.Cards.Set(ctx, msg.CardId, card)

	return &types.MsgCardArtistChangeResponse{}, nil
}
