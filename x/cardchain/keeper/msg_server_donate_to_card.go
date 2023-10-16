package keeper

import (
	"context"

	sdkerrors "cosmossdk.io/errors"
	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) DonateToCard(goCtx context.Context, msg *types.MsgDonateToCard) (*types.MsgDonateToCardResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.BurnCoinsFromString(ctx, msg.Creator, sdk.Coins{msg.Amount}) // If so, deduct the Bid amount from the sender
	if err != nil {
		return nil, sdkerrors.Wrap(errors.ErrInsufficientFunds, "Donator does not have enough coins")
	}

	card := k.Cards.Get(ctx, msg.CardId)
	card.VotePool = card.VotePool.Add(msg.Amount)
	k.Cards.Set(ctx, msg.CardId, card)

	return &types.MsgDonateToCardResponse{}, nil
}
