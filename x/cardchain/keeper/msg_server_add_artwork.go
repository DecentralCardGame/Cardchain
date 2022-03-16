package keeper

import (
	"fmt"
	"context"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) AddArtwork(goCtx context.Context, msg *types.MsgAddArtwork) (*types.MsgAddArtworkResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	card := k.GetCard(ctx, msg.CardId)

	if card.Artist != msg.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Artist")
	}

	if len(msg.Image) > 500000 {
		return nil, sdkerrors.Wrap(types.ErrImageSizeExceeded, fmt.Sprint(len(msg.Image)))
	}

	card.FullArt = msg.FullArt
	card.Image = msg.Image

	if card.Status == types.Status_suspended {
		card.Status = types.Status_permanent
	}

	k.SetCard(ctx, msg.CardId, card)

	return &types.MsgAddArtworkResponse{}, nil
}
