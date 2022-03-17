package keeper_test

import (
	"testing"

	testkeeper "github.com/DecentralCardGame/Cardchain/testutil/keeper"
	"github.com/DecentralCardGame/Cardchain/x/cardchain/keeper"
	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Sets up a Card
func setUpCard(ctx sdk.Context, k *keeper.Keeper) types.Card {
	addr, err := sdk.AccAddressFromBech32("cosmos15kq043zhu0wjyuw9av0auft06y3v2kxss862qf")
	if err != nil {
		panic(err)
	}
	card := types.NewCard(addr)
	card.Artist = addr.String()

	k.Cards.Set(ctx, 0, card)
	return card
}

func TestCard(t *testing.T) {
	k, ctx := testkeeper.CardchainKeeper(t)
	card := setUpCard(ctx, k)

	require.EqualValues(t, card, k.Cards.Get(ctx, 0))
	require.EqualValues(t, []*types.Card{&card}, k.Cards.GetAll(ctx))
	require.EqualValues(t, 1, k.Cards.GetNumber(ctx))
}

func TestCardAuctionPrice(t *testing.T) {
	k, ctx := testkeeper.CardchainKeeper(t)

	coin1 := sdk.NewInt64Coin("ucredits", 156)

	k.SetCardAuctionPrice(ctx, coin1)

	require.EqualValues(t, coin1, k.GetCardAuctionPrice(ctx))
}

func TestSetCardToTrial(t *testing.T) {
	k, ctx := testkeeper.CardchainKeeper(t)
	card := setUpCard(ctx, k)

	votePool := sdk.NewInt64Coin("ucredits", 10000)
	k.SetCardToTrial(ctx, 0, votePool)
	card = k.Cards.Get(ctx, 0)

	require.EqualValues(t, votePool, card.VotePool)
	require.EqualValues(t, types.Status_trial, card.Status)
	require.EqualValues(t, *new([]string), card.Voters)
}
