package keeper

import (
	"context"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) ChangeArtist(goCtx context.Context, msg *types.MsgChangeArtist) (*types.MsgChangeArtistResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	card := k.Cards.Get(ctx, msg.CardID)

	if card.Owner != msg.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner")
	}

	newArtist, err := sdk.AccAddressFromBech32(msg.Artist)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidAccAddress, "Invalid Artist address")
	}

	card.Artist = newArtist.String()

	k.Cards.Set(ctx, msg.CardID, card)

	return &types.MsgChangeArtistResponse{}, nil
}
