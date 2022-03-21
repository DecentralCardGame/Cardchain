package keeper_test

import (
	"testing"

	testkeeper "github.com/DecentralCardGame/Cardchain/testutil/keeper"
	"github.com/DecentralCardGame/Cardchain/x/cardchain/keeper"
	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestCouncil(t *testing.T) {
	k, ctx := testkeeper.CardchainKeeper(t)
	council := types.Council{
		CardId:   0,
		Treasury: sdk.NewInt64Coin("credits", 10),
		Status:   types.CouncelingStatus_councilCreated,
	}

	k.Councils.Set(ctx, 0, &council)

	require.EqualValues(t, council, *k.Councils.Get(ctx, 0))
	require.EqualValues(t, []*types.Council{&council}, k.Councils.GetAll(ctx))
	require.EqualValues(t, 1, k.Councils.GetNumber(ctx))
}

func TestResponseHash(t *testing.T) {
	hash := keeper.GetResponseHash(types.Response_Suggestion, "0")
	require.EqualValues(t, "20042f9326d18c469ae50aacb957247cca54101925022b8f463270bf266e756b", hash)
}

func TestGetCouncilPartiedVoters(t *testing.T) {
	responses := []*types.WrapClearResponse{
		&types.WrapClearResponse{"abc", types.Response_Yes, ""},
		&types.WrapClearResponse{"def", types.Response_Yes, ""},
		&types.WrapClearResponse{"ghi", types.Response_No, ""},
		&types.WrapClearResponse{"jkl", types.Response_No, ""},
		&types.WrapClearResponse{"mno", types.Response_Yes, ""},
		&types.WrapClearResponse{"pqr", types.Response_Suggestion, "bla"},
	}
	a, d, s := keeper.GetCouncilPartiedVoters(responses)
	require.EqualValues(t, []string{"abc", "def", "mno"}, a)
	require.EqualValues(t, []string{"ghi", "jkl"}, d)
	require.EqualValues(t, []string{"pqr"}, s)
}
