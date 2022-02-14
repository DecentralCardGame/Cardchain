package keeper

import (
	"context"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateSellOffer(goCtx context.Context, msg *types.MsgCreateSellOffer) (*types.MsgCreateSellOfferResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	creatorAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidAccAddress, "Unable to convert to AccAddress")
	}
	creator, err := k.GetUser(ctx, creatorAddr)
	if err != nil {
		return nil, err
	}

	newCards, err := uintPopElementFromArr(msg.Card, creator.Cards)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "User does not posses Card: %d", msg.Card)
	}
	creator.Cards = newCards

	sellOfferId := k.GetSellOffersNumber(ctx)
	sellOffer := types.SellOffer{
		Seller: msg.Creator,
		Buyer: "",
		Card: msg.Card,
		Price: msg.Price,
		Status: types.SellOfferStatus_open,
	}

	k.SetSellOffer(ctx, sellOfferId, sellOffer)
	k.SetUser(ctx, creatorAddr, creator)

	return &types.MsgCreateSellOfferResponse{}, nil
}
