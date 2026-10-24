package keeper

import (
	"context"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"math"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/gogoproto/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
)

var _ types.RunDeckMsgServer = runDeckMsgServer{}
var _ types.RunDeckQueryServer = Keeper{}

type runDeckMsgServer struct{ Keeper }

func NewRunDeckMsgServerImpl(k Keeper) types.RunDeckMsgServer { return runDeckMsgServer{k} }

func (k Keeper) runDeckIndex(ctx sdk.Context, kind, owner string) prefix.Store {
	return prefix.NewStore(runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx)), []byte("RunDeckIndexV1/"+kind+"/"+owner+"/"))
}
func runIDKey(id uint64) []byte { return binary.BigEndian.AppendUint64(nil, id) }
func completionKey(reporter, ref string) []byte {
	h := sha256.Sum256([]byte(reporter + "\x00" + ref))
	return h[:]
}
func (k Keeper) indexRunDeck(ctx sdk.Context, deck *types.SavedRunDeck) {
	k.runDeckIndex(ctx, "owner", deck.Owner).Set(runIDKey(deck.Id), []byte{1})
	k.runDeckIndex(ctx, "completed", "").Set(completionKey(deck.Reporter, deck.RunRef), runIDKey(deck.Id))
}
func (k Keeper) indexAllStar(ctx sdk.Context, card *types.AllStarCard) {
	k.runDeckIndex(ctx, "stars", card.Owner).Set(runIDKey(card.Id), []byte{1})
}

func (k runDeckMsgServer) SubmitRunDeck(goCtx context.Context, msg *types.MsgSubmitRunDeck) (*types.MsgSubmitRunDeckResponse, error) {
	if msg == nil {
		return nil, status.Error(codes.InvalidArgument, "nil request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := types.ValidateRunDeckIdentity(msg.Owner, msg.Reporter, msg.RunRef); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	reporter, err := k.GetUserFromString(ctx, msg.Reporter)
	if err != nil || !reporter.ReportMatches {
		return nil, status.Error(codes.PermissionDenied, "appointed reporter required")
	}
	ownerAddress, _ := sdk.AccAddressFromBech32(msg.Owner)
	if !k.Users.Has(ctx, ownerAddress) {
		return nil, status.Error(codes.NotFound, "owner not registered")
	}
	if k.runDeckIndex(ctx, "completed", "").Has(completionKey(msg.Reporter, msg.RunRef)) {
		return nil, status.Error(codes.AlreadyExists, "run already submitted by this reporter")
	}
	if err = types.ValidateRunDeckCards(msg.RecipeFormatVersion, msg.Cards); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	for _, card := range msg.Cards {
		if card.BaseCardId >= k.CardK.GetNum(ctx) {
			return nil, status.Error(codes.InvalidArgument, "base card does not exist")
		}
		base := k.CardK.Get(ctx, card.BaseCardId)
		if base.Status == types.CardStatus_none || len(base.Content) == 0 {
			return nil, status.Error(codes.InvalidArgument, "base card is empty or removed")
		}
	}
	id := k.RunDecks.GetNum(ctx)
	if id == math.MaxUint64 {
		return nil, status.Error(codes.ResourceExhausted, "deck IDs exhausted")
	}
	deck := &types.SavedRunDeck{Id: id, Owner: msg.Owner, Reporter: msg.Reporter, RunRef: msg.RunRef, RecipeFormatVersion: msg.RecipeFormatVersion, CreatedHeight: ctx.BlockHeight()}
	for _, c := range msg.Cards {
		cloned := proto.Clone(c).(*types.RunDeckCard)
		for i, raw := range cloned.ForgeRecipes {
			canonical, err := types.CanonicalForgeRecipe(raw)
			if err != nil {
				return nil, status.Error(codes.InvalidArgument, err.Error())
			}
			cloned.ForgeRecipes[i] = canonical
		}
		deck.Cards = append(deck.Cards, cloned)
	}
	// All validation precedes the first write. The SDK caches this entire message.
	k.RunDecks.Append(ctx, deck)
	k.indexRunDeck(ctx, deck)
	ctx.EventManager().EmitEvent(sdk.NewEvent("run_deck_submitted", sdk.NewAttribute("deck_id", fmt.Sprint(id)), sdk.NewAttribute("owner", deck.Owner), sdk.NewAttribute("reporter", deck.Reporter)))
	return &types.MsgSubmitRunDeckResponse{DeckId: id}, nil
}

func (k runDeckMsgServer) DistillRunDeck(goCtx context.Context, msg *types.MsgDistillRunDeck) (*types.MsgDistillRunDeckResponse, error) {
	if msg == nil {
		return nil, status.Error(codes.InvalidArgument, "nil request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := types.ValidateRunDeckAddress(msg.Owner); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	if msg.DeckId >= k.RunDecks.GetNum(ctx) {
		return nil, status.Error(codes.NotFound, "deck not found")
	}
	deck := k.RunDecks.Get(ctx, msg.DeckId)
	if deck.Owner != msg.Owner {
		return nil, status.Error(codes.PermissionDenied, "deck owner required")
	}
	if deck.Consumed {
		return nil, status.Error(codes.FailedPrecondition, "deck already consumed")
	}
	var selected *types.RunDeckCard
	for _, c := range deck.Cards {
		if c.InstanceId == msg.InstanceId {
			selected = c
			break
		}
	}
	if selected == nil {
		return nil, status.Error(codes.NotFound, "card occurrence not in deck")
	}
	id := k.AllStarCards.GetNum(ctx)
	if id == math.MaxUint64 {
		return nil, status.Error(codes.ResourceExhausted, "All-Star IDs exhausted")
	}
	star := &types.AllStarCard{Id: id, Owner: deck.Owner, SourceDeckId: deck.Id, RecipeFormatVersion: deck.RecipeFormatVersion, Card: proto.Clone(selected).(*types.RunDeckCard), CreatedHeight: ctx.BlockHeight()}
	deck.Consumed = true
	deck.ConsumedHeight = ctx.BlockHeight()
	deck.AllStarId = id
	k.AllStarCards.Append(ctx, star)
	k.indexAllStar(ctx, star)
	k.RunDecks.Set(ctx, deck.Id, deck)
	ctx.EventManager().EmitEvent(sdk.NewEvent("run_deck_distilled", sdk.NewAttribute("deck_id", fmt.Sprint(deck.Id)), sdk.NewAttribute("all_star_id", fmt.Sprint(id)), sdk.NewAttribute("owner", deck.Owner)))
	return &types.MsgDistillRunDeckResponse{AllStarId: id}, nil
}

func (k Keeper) RunDeck(goCtx context.Context, req *types.QueryRunDeckRequest) (*types.QueryRunDeckResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "nil request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	if req.Id >= k.RunDecks.GetNum(ctx) {
		return nil, status.Error(codes.NotFound, "deck not found")
	}
	return &types.QueryRunDeckResponse{Deck: k.RunDecks.Get(ctx, req.Id)}, nil
}
func (k Keeper) AllStar(goCtx context.Context, req *types.QueryAllStarRequest) (*types.QueryAllStarResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "nil request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	if req.Id >= k.AllStarCards.GetNum(ctx) {
		return nil, status.Error(codes.NotFound, "All-Star not found")
	}
	return &types.QueryAllStarResponse{AllStar: k.AllStarCards.Get(ctx, req.Id)}, nil
}
func (k Keeper) runDeckPage(ctx sdk.Context, kind, owner string, cursor uint64, limit uint32) (ids []uint64, next uint64, more bool, err error) {
	if e := types.ValidateRunDeckAddress(owner); e != nil {
		return nil, 0, false, status.Error(codes.InvalidArgument, e.Error())
	}
	if limit == 0 {
		limit = 20
	}
	if limit > 100 {
		return nil, 0, false, status.Error(codes.InvalidArgument, "limit must not exceed 100")
	}
	iter := k.runDeckIndex(ctx, kind, owner).Iterator(runIDKey(cursor), nil)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		id := binary.BigEndian.Uint64(iter.Key())
		if len(ids) == int(limit) {
			return ids, id, true, nil
		}
		ids = append(ids, id)
	}
	return ids, 0, false, nil
}
func (k Keeper) RunDecksByOwner(goCtx context.Context, req *types.QueryRunDecksRequest) (*types.QueryRunDecksResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "nil request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	ids, next, more, err := k.runDeckPage(ctx, "owner", req.Owner, req.Cursor, req.Limit)
	if err != nil {
		return nil, err
	}
	res := &types.QueryRunDecksResponse{NextCursor: next, HasMore: more}
	for _, id := range ids {
		res.Decks = append(res.Decks, k.RunDecks.Get(ctx, id))
	}
	return res, nil
}
func (k Keeper) AllStars(goCtx context.Context, req *types.QueryAllStarsRequest) (*types.QueryAllStarsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "nil request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	ids, next, more, err := k.runDeckPage(ctx, "stars", req.Owner, req.Cursor, req.Limit)
	if err != nil {
		return nil, err
	}
	res := &types.QueryAllStarsResponse{NextCursor: next, HasMore: more}
	for _, id := range ids {
		res.AllStars = append(res.AllStars, k.AllStarCards.Get(ctx, id))
	}
	return res, nil
}

func (k Keeper) InitRunDeckGenesis(ctx sdk.Context, state *types.RunDeckGenesis) error {
	if err := types.ValidateRunDeckGenesis(state); err != nil {
		return err
	}
	if state == nil {
		return nil
	}
	if k.RunDecks.GetNum(ctx) != 0 || k.AllStarCards.GetNum(ctx) != 0 {
		return fmt.Errorf("run deck genesis requires empty stores")
	}
	for _, deck := range state.Decks {
		k.RunDecks.Append(ctx, deck)
		k.indexRunDeck(ctx, deck)
	}
	for _, card := range state.AllStars {
		k.AllStarCards.Append(ctx, card)
		k.indexAllStar(ctx, card)
	}
	return nil
}
func (k Keeper) ExportRunDeckGenesis(ctx sdk.Context) *types.RunDeckGenesis {
	// Explicit ID iteration avoids the legacy generic GetAll iterator lifecycle.
	state := &types.RunDeckGenesis{}
	for i := uint64(0); i < k.RunDecks.GetNum(ctx); i++ {
		state.Decks = append(state.Decks, k.RunDecks.Get(ctx, i))
	}
	for i := uint64(0); i < k.AllStarCards.GetNum(ctx); i++ {
		state.AllStars = append(state.AllStars, k.AllStarCards.Get(ctx, i))
	}
	return state
}
