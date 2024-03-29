package keeper_test

import (
	"testing"

	testkeeper "github.com/DecentralCardGame/Cardchain/testutil/keeper"
	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestSellOffer(t *testing.T) {
	k, ctx := testkeeper.CardchainKeeper(t)
	setUpCard(ctx, k)
	sellOffer := types.SellOffer{
		Seller: "cosmos15kq043zhu0wjyuw9av0auft06y3v2kxss862qf",
		Price:  sdk.NewInt64Coin("ucredits", 3),
		Status: types.SellOfferStatus_open,
	}
	k.SellOffers.Set(ctx, 0, &sellOffer)

	require.EqualValues(t, 1, k.SellOffers.GetNumber(ctx))
	require.EqualValues(t, sellOffer, *k.SellOffers.Get(ctx, 0))
	require.EqualValues(t, []*types.SellOffer{&sellOffer}, k.SellOffers.GetAll(ctx))
}
