package keeper

import (
	"testing"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/keeper"
	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store"

	// storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	typesparams "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmdb "github.com/tendermint/tm-db"
)

func CardchainKeeper(t testing.TB) (*keeper.Keeper, sdk.Context) {
	storeKey := sdk.NewKVStoreKey(types.StoreKey)
	usersStoreKey := sdk.NewKVStoreKey(types.UsersStoreKey)
	cardsStoreKey := sdk.NewKVStoreKey(types.CardsStoreKey)
	matchesStoreKey := sdk.NewKVStoreKey(types.MatchesStoreKey)
	setsStoreKey := sdk.NewKVStoreKey(types.SetsStoreKey)
	internalStoreKey := sdk.NewKVStoreKey(types.InternalStoreKey)
	sellOffersStoreKey := sdk.NewKVStoreKey(types.SellOffersStoreKey)
	poolsStoreKey := sdk.NewKVStoreKey(types.PoolsStoreKey)
	serversStoreKey := sdk.NewKVStoreKey(types.ServersStoreKey)
	zealyStoreKey := sdk.NewKVStoreKey(types.ZealyStoreKey)
	encountersStoreKey := sdk.NewKVStoreKey(types.EncountersStoreKey)
	runningAveragesStoreKey := sdk.NewKVStoreKey(types.RunningAveragesStoreKey)
	councilsStoreKey := sdk.NewKVStoreKey(types.CouncilsStoreKey)

	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	stateStore.MountStoreWithDB(storeKey, sdk.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(usersStoreKey, sdk.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(cardsStoreKey, sdk.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(matchesStoreKey, sdk.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(setsStoreKey, sdk.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(internalStoreKey, sdk.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(sellOffersStoreKey, sdk.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(poolsStoreKey, sdk.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(serversStoreKey, sdk.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(zealyStoreKey, sdk.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(councilsStoreKey, sdk.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(runningAveragesStoreKey, sdk.StoreTypeIAVL, db)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)

	paramsSubspace := typesparams.NewSubspace(cdc,
		types.Amino,
		storeKey,
		internalStoreKey,
		"CardchainParams",
	)
	k := keeper.NewKeeper(
		cdc,
		usersStoreKey,
		cardsStoreKey,
		matchesStoreKey,
		setsStoreKey,
		sellOffersStoreKey,
		poolsStoreKey,
		councilsStoreKey,
		runningAveragesStoreKey,
		serversStoreKey,
		zealyStoreKey,
		internalStoreKey,
		paramsSubspace,
		nil, // That's why minting fails
	)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())

	// Initialize params
	k.SetParams(ctx, types.DefaultParams())

	return k, ctx
}
