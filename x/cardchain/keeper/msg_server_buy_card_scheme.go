package keeper

import (
	"context"
	"math"

	sdkerrors "cosmossdk.io/errors"
	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) BuyCardScheme(goCtx context.Context, msg *types.MsgBuyCardScheme) (*types.MsgBuyCardSchemeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	currId := k.Cards.GetNum(ctx)
	price := k.GetCardAuctionPrice(ctx)
	bid := msg.Bid

	buyer, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidAccAddress, "Unable to convert to AccAddress")
	}

	if bid.IsLT(price) { // Checks if the bid is less than price
		return nil, sdkerrors.Wrap(errors.ErrInsufficientFunds, "Bid not high enough")
	}

	err = k.BurnCoinsFromAddr(ctx, buyer, sdk.Coins{price}) // If so, deduct the Bid amount from the sender
	if err != nil {
		return nil, sdkerrors.Wrapf(errors.ErrInsufficientFunds, "Buyer does not have enough coins %s", err)
	}

	k.AddPoolCredits(ctx, PublicPoolKey, price)
	if price.Amount.Mul(sdk.NewInt(2)).LT(sdk.NewInt(int64(math.Pow(10, 12)))) {
		k.SetCardAuctionPrice(ctx, price.Add(price))
	}

	newCard := types.NewCard(buyer)
	newCard.ImageId = k.Images.GetNum(ctx)
	image := types.Image{}

	k.Cards.Set(ctx, currId, &newCard)
	k.Images.Set(ctx, newCard.ImageId, &image)
	k.AddOwnedCardScheme(ctx, currId, buyer)

	return &types.MsgBuyCardSchemeResponse{
		CardId: currId,
	}, nil
}
