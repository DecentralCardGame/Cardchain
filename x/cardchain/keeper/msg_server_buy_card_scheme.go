package keeper

import (
	"context"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) BuyCardScheme(goCtx context.Context, msg *types.MsgBuyCardScheme) (*types.MsgBuyCardSchemeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	lastId := k.GetLastCardSchemeId(ctx) // first get last card bought id
	currId := lastId + 1                      // iterate it by 1

	price := k.GetCardAuctionPrice(ctx)

	buyer, err := sdk.AccAddressFromBech32(msg.Buyer)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidAccAddress, "Unable to convert to AccAddress")
	}

	if msg.Bid.IsLT(price) { // Checks if the bid is less than price
		return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "Bid not high enough")
	}

	err = k.BurnCoinsFromAddr(ctx, buyer, sdk.Coins{price}) // If so, deduct the Bid amount from the sender
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "Buyer does not have enough coins")
	}

	k.AddPoolCredits(ctx, PublicPoolKey, price)
	k.SetCardAuctionPrice(ctx, price.Add(price))
	k.SetLastCardSchemeId(ctx, currId)

	newCard := types.NewCard(buyer)

	k.SetCard(ctx, currId, newCard)
	k.AddOwnedCardScheme(ctx, currId, buyer)

	return &types.MsgBuyCardSchemeResponse{}, nil
}
