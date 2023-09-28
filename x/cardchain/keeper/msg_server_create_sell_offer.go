package keeper

import (
	"context"
	"github.com/cosmos/cosmos-sdk/types/errors"

	sdkerrors "cosmossdk.io/errors"
	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateSellOffer(goCtx context.Context, msg *types.MsgCreateSellOffer) (*types.MsgCreateSellOfferResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	creator, err := k.GetUserFromString(ctx, msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrUserDoesNotExist, err.Error())
	}

	// Check if card is startercard
	card := k.Cards.Get(ctx, msg.Card)
	if card.StarterCard {
		return nil, sdkerrors.Wrap(errors.ErrUnauthorized, "Card is startercard and therefore not sellable")
	}

	newCards, err := PopItemFromArr(msg.Card, creator.Cards)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "User does not posses Card: %d", msg.Card)
	}
	creator.Cards = newCards

	sellOfferId := k.SellOffers.GetNum(ctx)
	sellOffer := types.SellOffer{
		Seller: msg.Creator,
		Buyer:  "",
		Card:   msg.Card,
		Price:  msg.Price,
		Status: types.SellOfferStatus_open,
	}

	k.SellOffers.Set(ctx, sellOfferId, &sellOffer)
	k.SetUserFromUser(ctx, creator)

	return &types.MsgCreateSellOfferResponse{}, nil
}
