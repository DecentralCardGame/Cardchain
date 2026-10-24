package keeper_test

import (
	"bytes"
	"context"
	"testing"

	keepertest "github.com/DecentralCardGame/cardchain/testutil/keeper"
	"github.com/DecentralCardGame/cardchain/x/cardchain/keeper"
	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/gogoproto/proto"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type runFixture struct {
	k                      keeper.Keeper
	ctx                    sdk.Context
	msg                    types.RunDeckMsgServer
	owner, reporter, other string
}

func setupRunDeck(t *testing.T) runFixture {
	k, ctx := keepertest.CardchainKeeper(t)
	ctx = ctx.WithBlockHeight(10)
	owner := sdk.AccAddress(bytes.Repeat([]byte{1}, 20))
	reporter := sdk.AccAddress(bytes.Repeat([]byte{2}, 20))
	other := sdk.AccAddress(bytes.Repeat([]byte{3}, 20))
	k.Users.Set(ctx, owner, &types.User{Alias: "owner"})
	k.Users.Set(ctx, reporter, &types.User{Alias: "reporter", ReportMatches: true})
	k.Users.Set(ctx, other, &types.User{Alias: "other"})
	k.CardK.Set(ctx, 0, &types.Card{Status: types.CardStatus_permanent, Content: []byte(`{"balance":"old"}`)})
	k.CardK.Set(ctx, 1, &types.Card{})
	return runFixture{k, ctx, keeper.NewRunDeckMsgServerImpl(k), owner.String(), reporter.String(), other.String()}
}
func (f runFixture) submit(ref string) *types.MsgSubmitRunDeck {
	return &types.MsgSubmitRunDeck{Reporter: f.reporter, Owner: f.owner, RunRef: ref, RecipeFormatVersion: 1, Cards: []*types.RunDeckCard{
		{InstanceId: 0, BaseCardId: 0, ForgeRecipes: []string{`{ "FixedPackageId": "package.a" }`, `{"FixedPackageId":"package.a"}`}},
		{InstanceId: 42, BaseCardId: 0, ForgeRecipes: []string{`{"Generated":{"Effects":[{"Sacrifice":{"Target":"THIS"}}],"EffectParameters":[{"EffectIndex":0,"Amount":2}]}}`}},
	}}
}

func TestRunDeckSubmitDistillAndRoundTrip(t *testing.T) {
	f := setupRunDeck(t)
	req := f.submit("stable-run-1")
	receipt, err := f.msg.SubmitRunDeck(f.ctx, req)
	require.NoError(t, err)
	require.Zero(t, receipt.DeckId)
	req.Cards[0].ForgeRecipes[0] = "mutated caller object"
	before, err := f.k.RunDeck(f.ctx, &types.QueryRunDeckRequest{Id: 0})
	require.NoError(t, err)
	require.Equal(t, `{"FixedPackageId":"package.a"}`, before.Deck.Cards[0].ForgeRecipes[0])
	require.Len(t, before.Deck.Cards[0].ForgeRecipes, 2)
	_, err = f.msg.SubmitRunDeck(f.ctx, f.submit("stable-run-1"))
	require.Equal(t, codes.AlreadyExists, status.Code(err))
	// Base balance changes never rewrite durable recipes.
	f.k.CardK.Set(f.ctx, 0, &types.Card{Status: types.CardStatus_permanent, Content: []byte(`{"balance":"new"}`)})
	receipt2, err := f.msg.DistillRunDeck(f.ctx.WithBlockHeight(20), &types.MsgDistillRunDeck{Owner: f.owner, DeckId: 0, InstanceId: 42})
	require.NoError(t, err)
	require.Zero(t, receipt2.AllStarId)
	star, err := f.k.AllStar(f.ctx, &types.QueryAllStarRequest{Id: 0})
	require.NoError(t, err)
	require.True(t, proto.Equal(before.Deck.Cards[1], star.AllStar.Card))
	require.Equal(t, f.owner, star.AllStar.Owner)
	_, err = f.msg.DistillRunDeck(f.ctx, &types.MsgDistillRunDeck{Owner: f.owner, DeckId: 0, InstanceId: 0})
	require.Equal(t, codes.FailedPrecondition, status.Code(err))
	require.EqualValues(t, 1, f.k.AllStarCards.GetNum(f.ctx))
	snapshot := f.k.ExportRunDeckGenesis(f.ctx)
	require.NoError(t, types.ValidateRunDeckGenesis(snapshot))
	k2, ctx2 := keepertest.CardchainKeeper(t)
	require.NoError(t, k2.InitRunDeckGenesis(ctx2, snapshot))
	require.True(t, proto.Equal(snapshot, k2.ExportRunDeckGenesis(ctx2)))
	// Rebuilt uniqueness and owner indexes survive import.
	page, err := k2.RunDecksByOwner(ctx2, &types.QueryRunDecksRequest{Owner: f.owner})
	require.NoError(t, err)
	require.Len(t, page.Decks, 1)
	require.True(t, page.Decks[0].Consumed)
	stars, err := k2.AllStars(ctx2, &types.QueryAllStarsRequest{Owner: f.owner})
	require.NoError(t, err)
	require.Len(t, stars.AllStars, 1)
	snapshot.AllStars[0].Card.ForgeRecipes[0] = `{"FixedPackageId":"tampered"}`
	require.Error(t, types.ValidateRunDeckGenesis(snapshot))
}

func TestRunDeckRejectsUnauthorizedAndInvalidSubmissions(t *testing.T) {
	f := setupRunDeck(t)
	cases := []func(*types.MsgSubmitRunDeck){
		func(m *types.MsgSubmitRunDeck) { m.Reporter = f.other; f.k.GetServerId(f.ctx, f.other) },
		func(m *types.MsgSubmitRunDeck) { m.Owner = sdk.AccAddress(bytes.Repeat([]byte{9}, 20)).String() },
		func(m *types.MsgSubmitRunDeck) { m.Cards[0].BaseCardId = 99 },
		func(m *types.MsgSubmitRunDeck) { m.Cards[0].BaseCardId = 1 },
		func(m *types.MsgSubmitRunDeck) { m.Cards[1].InstanceId = 0 },
		func(m *types.MsgSubmitRunDeck) {
			m.Cards[0].ForgeRecipes = []string{`{"Generated":{"Effects":[{"Invalid":{}}]}}`}
		},
		func(m *types.MsgSubmitRunDeck) { m.RunRef = "" },
	}
	for _, mutate := range cases {
		req := f.submit("same-ref")
		mutate(req)
		_, err := f.msg.SubmitRunDeck(f.ctx, req)
		require.Error(t, err)
		require.Zero(t, f.k.RunDecks.GetNum(f.ctx))
	}
	_, err := f.msg.SubmitRunDeck(f.ctx, f.submit("same-ref"))
	require.NoError(t, err)
	for _, request := range []*types.MsgDistillRunDeck{
		{Owner: f.other, DeckId: 0, InstanceId: 0}, {Owner: f.owner, DeckId: 0, InstanceId: 123}, {Owner: f.owner, DeckId: 99, InstanceId: 0},
	} {
		_, err = f.msg.DistillRunDeck(f.ctx, request)
		require.Error(t, err)
		require.Zero(t, f.k.AllStarCards.GetNum(f.ctx))
		deck, _ := f.k.RunDeck(f.ctx, &types.QueryRunDeckRequest{Id: 0})
		require.False(t, deck.Deck.Consumed)
	}
	_, err = f.msg.SubmitRunDeck(context.Background(), nil)
	require.Error(t, err)
}

func TestRunDeckPaginationAndGenesisCorruption(t *testing.T) {
	f := setupRunDeck(t)
	for _, ref := range []string{"a", "b", "c"} {
		_, err := f.msg.SubmitRunDeck(f.ctx, f.submit(ref))
		require.NoError(t, err)
	}
	a, err := f.k.RunDecksByOwner(f.ctx, &types.QueryRunDecksRequest{Owner: f.owner, Limit: 2})
	require.NoError(t, err)
	require.Len(t, a.Decks, 2)
	require.True(t, a.HasMore)
	require.EqualValues(t, 2, a.NextCursor)
	b, err := f.k.RunDecksByOwner(f.ctx, &types.QueryRunDecksRequest{Owner: f.owner, Limit: 2, Cursor: a.NextCursor})
	require.NoError(t, err)
	require.Len(t, b.Decks, 1)
	require.False(t, b.HasMore)
	empty, err := f.k.RunDecksByOwner(f.ctx, &types.QueryRunDecksRequest{Owner: f.other})
	require.NoError(t, err)
	require.Empty(t, empty.Decks)
	_, err = f.k.RunDecksByOwner(f.ctx, &types.QueryRunDecksRequest{Owner: f.owner, Limit: 101})
	require.Error(t, err)
	_, err = f.k.AllStar(f.ctx, &types.QueryAllStarRequest{Id: 0})
	require.Equal(t, codes.NotFound, status.Code(err))
	state := f.k.ExportRunDeckGenesis(f.ctx)
	state.Decks[1].RunRef = "a"
	require.Error(t, types.ValidateRunDeckGenesis(state))
	state = f.k.ExportRunDeckGenesis(f.ctx)
	state.Decks[0].Consumed = true
	require.Error(t, types.ValidateRunDeckGenesis(state))
	state = f.k.ExportRunDeckGenesis(f.ctx)
	state.Decks[0].Id = 2
	require.Error(t, types.ValidateRunDeckGenesis(state))
}
