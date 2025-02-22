package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/types/errors"

	sdkerrors "cosmossdk.io/errors"
	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) AddArtwork(goCtx context.Context, msg *types.MsgAddArtwork) (*types.MsgAddArtworkResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	card := k.Cards.Get(ctx, msg.CardId)
	image := k.Images.Get(ctx, card.ImageId)

	councilEnabled, err := k.FeatureFlagModuleInstance.Get(ctx, string(types.FeatureFlagName_Council))
	if err != nil {
		return nil, err
	}

	if councilEnabled && card.Status != types.Status_prototype && card.Status != types.Status_scheme {
		return nil, sdkerrors.Wrap(types.ErrInvalidCardStatus, "Card has to be a prototype to be changeable")
	}

	if card.Artist != msg.Creator {
		return nil, sdkerrors.Wrap(errors.ErrUnauthorized, "Incorrect Artist")
	}

	card.FullArt = msg.FullArt
	image.Image = msg.Image

	if card.Status == types.Status_suspended {
		card.Status = types.Status_permanent
	}

	k.Cards.Set(ctx, msg.CardId, card)
	k.Images.Set(ctx, card.ImageId, image)

	return &types.MsgAddArtworkResponse{}, nil
}
