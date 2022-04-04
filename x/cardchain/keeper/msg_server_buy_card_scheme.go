package keeper

import (
	"context"
	"math"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) BuyCardScheme(goCtx context.Context, msg *types.MsgBuyCardScheme) (*types.MsgBuyCardSchemeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	currId := k.Cards.GetNumber(ctx)
	price := k.GetCardAuctionPrice(ctx)

	bid, err := sdk.ParseCoinNormalized(msg.Bid)
	if err != nil {
		return nil, err
	}

	buyer, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidAccAddress, "Unable to convert to AccAddress")
	}

	if bid.IsLT(price) { // Checks if the bid is less than price
		return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "Bid not high enough")
	}

	err = k.BurnCoinsFromAddr(ctx, buyer, sdk.Coins{price}) // If so, deduct the Bid amount from the sender
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInsufficientFunds, "Buyer does not have enough coins %s", err)
	}

	k.AddPoolCredits(ctx, PublicPoolKey, price)
	if price.Amount.Mul(sdk.NewInt(2)).LT(sdk.NewInt(int64(math.Pow(10, 12)))) {
		k.SetCardAuctionPrice(ctx, price.Add(price))
	}

	newCard := types.NewCard(buyer)

	k.Cards.Set(ctx, currId, &newCard)
	k.AddOwnedCardScheme(ctx, currId, buyer)

	return &types.MsgBuyCardSchemeResponse{}, nil
}
