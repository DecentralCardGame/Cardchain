package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"github.com/DecentralCardGame/cardchain/util"
	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) SellOfferCreate(goCtx context.Context, msg *types.MsgSellOfferCreate) (*types.MsgSellOfferCreateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	creator, err := k.GetUserFromString(ctx, msg.Creator)
	if err != nil {
		return nil, errorsmod.Wrap(types.ErrUserDoesNotExist, err.Error())
	}

	// Check if card is startercard
	card := k.CardK.Get(ctx, msg.CardId)
	if card.StarterCard {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "Card is startercard and therefore not sellable")
	}

	newCards, err := util.PopItemFromArr(msg.CardId, creator.Cards)
	if err != nil {
		return nil, errorsmod.Wrapf(err, "User does not posses Card: %d", msg.CardId)
	}
	creator.Cards = newCards

	sellOfferId := k.SellOfferK.GetNum(ctx)
	sellOffer := types.SellOffer{
		Seller: msg.Creator,
		Buyer:  "",
		Card:   msg.CardId,
		Price:  msg.Price,
		Status: types.SellOfferStatus_open,
	}

	k.SellOfferK.Set(ctx, sellOfferId, &sellOffer)
	k.SetUserFromUser(ctx, creator)

	return &types.MsgSellOfferCreateResponse{}, nil
}
