package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CardDonate(goCtx context.Context, msg *types.MsgCardDonate) (*types.MsgCardDonateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.BurnCoinsFromString(ctx, msg.Creator, sdk.Coins{msg.Amount}) // If so, deduct the Bid amount from the sender
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInsufficientFunds, "Donator does not have enough coins")
	}

	card := k.Cards.Get(ctx, msg.CardId)
	card.VotePool = card.VotePool.Add(msg.Amount)
	k.Cards.Set(ctx, msg.CardId, card)

	return &types.MsgCardDonateResponse{}, nil
}
