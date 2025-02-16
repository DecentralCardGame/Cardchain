package keeper

import (
	"context"
	"math"

	errorsmod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CardSchemeBuy(goCtx context.Context, msg *types.MsgCardSchemeBuy) (*types.MsgCardSchemeBuyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	currId := k.cards.GetNum(ctx)
	price := *k.cardAuctionPrice.Get(ctx)
	bid := msg.Bid

	buyer, err := k.GetUserFromString(ctx, msg.Creator)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "Unable to convert to AccAddress")
	}

	if bid.IsLT(price) { // Checks if the bid is less than price
		return nil, errorsmod.Wrap(sdkerrors.ErrInsufficientFunds, "Bid not high enough")
	}

	err = k.BurnCoinsFromAddr(ctx, buyer.Addr, sdk.Coins{price}) // If so, deduct the Bid amount from the sender
	if err != nil {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInsufficientFunds, "Buyer does not have enough coins %s", err)
	}

	k.AddPoolCredits(ctx, PublicPoolKey, price)
	if price.Amount.Mul(sdkmath.NewInt(2)).LT(sdkmath.NewInt(int64(math.Pow(10, 12)))) {
		newPrice := price.Add(price)
		k.cardAuctionPrice.Set(ctx, &newPrice)
	}

	newCard := types.NewCard(buyer.Addr)
	newCard.ImageId = k.images.GetNum(ctx)
	image := types.Image{}

	buyer.OwnedCardSchemes = append(buyer.OwnedCardSchemes, currId)

	k.SetUserFromUser(ctx, buyer)
	k.cards.Set(ctx, currId, &newCard)
	k.images.Set(ctx, newCard.ImageId, &image)

	return &types.MsgCardSchemeBuyResponse{CardId: currId}, nil
}
