package keeper

import (
	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetLastCardModifiedNow Sets the block height the last card was modified at, to now
func (k Keeper) SetLastCardModifiedNow(ctx sdk.Context) {
	timeStamp := types.TimeStamp{TimeStamp: uint64(ctx.BlockHeight())}
	k.lastCardModified.Set(ctx, &timeStamp)
}

// SetCardToTrial Sets a card to trial
func (k Keeper) SetCardToTrial(ctx sdk.Context, cardId uint64, votePool sdk.Coin) {
	card := k.cards.Get(ctx, cardId)
	card.ResetVotes()
	card.VotePool = card.VotePool.Add(votePool)
	card.Status = types.CardStatus_trial
	k.cards.Set(ctx, cardId, card)
	k.SetLastCardModifiedNow(ctx)
}
