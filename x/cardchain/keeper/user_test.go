package keeper_test

// Won't test until bankkeeper works
// import (
// 	"testing"
//
// 	testkeeper "github.com/DecentralCardGame/Cardchain/testutil/keeper"
// 	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
//   sdk "github.com/cosmos/cosmos-sdk/types"
// 	"github.com/stretchr/testify/require"
// )
//
// func TestSellOffer(t *testing.T) {
// 	k, ctx := testkeeper.CardchainKeeper(t)
//   setUpCard(ctx, k)
// 	sellOffer := types.SellOffer {
//     "cosmos15kq043zhu0wjyuw9av0auft06y3v2kxss862qf",
//     "",
//     0,
//     sdk.NewInt64Coin("ucredits", 3),
//     types.SellOfferStatus_open,
//   }
//   k.SetSellOffer(ctx, 1, sellOffer)
//
//   require.EqualValues(t, 1, k.GetSellOffersNumber(ctx))
// 	require.EqualValues(t, sellOffer, k.GetSellOffer(ctx, 0))
//   require.EqualValues(t, []*types.SellOffer{&sellOffer}, k.GetAllSellOffers(ctx))
// }
