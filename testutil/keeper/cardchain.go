package keeper

import (
	"testing"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/keeper"
	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	typesparams "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/stretchr/testify/require"
	"github.com/cometbft/cometbft/libs/log"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	dbm "github.com/cometbft/cometbft-db"
)

func CardchainKeeper(t testing.TB) (*keeper.Keeper, sdk.Context) {
	storeKey := sdk.NewKVStoreKey(types.StoreKey)
	usersStoreKey := sdk.NewKVStoreKey(types.UsersStoreKey)
	cardsStoreKey := sdk.NewKVStoreKey(types.CardsStoreKey)
	matchesStoreKey := sdk.NewKVStoreKey(types.MatchesStoreKey)
	collectionsStoreKey := sdk.NewKVStoreKey(types.CollectionsStoreKey)
	internalStoreKey := sdk.NewKVStoreKey(types.InternalStoreKey)
	sellOffersStoreKey := sdk.NewKVStoreKey(types.SellOffersStoreKey)
	poolsStoreKey := sdk.NewKVStoreKey(types.PoolsStoreKey)
	serversStoreKey := sdk.NewKVStoreKey(types.ServersStoreKey)
	runningAveragesStoreKey := sdk.NewKVStoreKey(types.RunningAveragesStoreKey)
	councilsStoreKey := sdk.NewKVStoreKey(types.CouncilsStoreKey)
	imagesStorekey := sdk.NewKVStoreKey(types.ImagesStoreKey)

	db := dbm.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(usersStoreKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(cardsStoreKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(matchesStoreKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(collectionsStoreKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(internalStoreKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(sellOffersStoreKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(poolsStoreKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(serversStoreKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(councilsStoreKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(runningAveragesStoreKey, storetypes.StoreTypeIAVL, db)
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
		collectionsStoreKey,
		sellOffersStoreKey,
		poolsStoreKey,
		councilsStoreKey,
		runningAveragesStoreKey,
		imagesStorekey,
		serversStoreKey,
		internalStoreKey,
		paramsSubspace,
		nil, // That's why minting fails
	)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())

	// Initialize params
	k.SetParams(ctx, types.DefaultParams())

	return k, ctx
}
