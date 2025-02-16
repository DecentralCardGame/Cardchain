package keeper

import (
	"fmt"

	"cosmossdk.io/core/store"
	"cosmossdk.io/log"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	gtk "github.com/DecentralCardGame/cardchain/types/generic_type_keeper"
	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	ffKeeper "github.com/DecentralCardGame/cardchain/x/featureflag/keeper"
)

type (
	Keeper struct {
		cdc          codec.BinaryCodec
		storeService store.KVStoreService
		logger       log.Logger

		cards            gtk.GenericUint64TypeKeeper[*types.Card]
		councils         gtk.GenericUint64TypeKeeper[*types.Council]
		sellOffers       gtk.GenericUint64TypeKeeper[*types.SellOffer]
		sets             gtk.GenericUint64TypeKeeper[*types.Set]
		matches          gtk.GenericUint64TypeKeeper[*types.Match]
		servers          gtk.GenericUint64TypeKeeper[*types.Server]
		RunningAverages  gtk.KeywordedGenericTypeKeeper[*types.RunningAverage]
		Pools            gtk.KeywordedGenericTypeKeeper[*sdk.Coin]
		images           gtk.GenericUint64TypeKeeper[*types.Image]
		encounters       gtk.GenericUint64TypeKeeper[*types.Encounter]
		users            gtk.GenericAddressTypeKeeper[*types.User]
		zealy            gtk.GenericStringTypeKeeper[*types.Zealy]
		lastCardModified gtk.SingleValueGenericTypeKeeper[*types.TimeStamp]
		cardAuctionPrice gtk.SingleValueGenericTypeKeeper[*sdk.Coin]

		// the address capable of executing a MsgUpdateParams message. Typically, this
		// should be the x/gov module account.
		authority string

		BankKeeper                types.BankKeeper
		FeatureFlagModuleInstance ffKeeper.ModuleInstance
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeService store.KVStoreService,
	logger log.Logger,
	authority string,
	bankKeeper types.BankKeeper,
	featureFlagKeeper types.FeatureFlagKeeper,

) Keeper {
	if _, err := sdk.AccAddressFromBech32(authority); err != nil {
		panic(fmt.Sprintf("invalid authority address: %s", authority))
	}

	return Keeper{
		cdc:          cdc,
		storeService: storeService,
		authority:    authority,
		logger:       logger,

		cards:            gtk.NewUintGTK[*types.Card]("Cards", storeService, cdc, gtk.GetEmpty[types.Card]),
		councils:         gtk.NewUintGTK[*types.Council]("Councils", storeService, cdc, gtk.GetEmpty[types.Council]),
		sellOffers:       gtk.NewUintGTK[*types.SellOffer]("SellOffers", storeService, cdc, gtk.GetEmpty[types.SellOffer]),
		sets:             gtk.NewUintGTK[*types.Set]("Sets", storeService, cdc, gtk.GetEmpty[types.Set]),
		matches:          gtk.NewUintGTK[*types.Match]("Matches", storeService, cdc, gtk.GetEmpty[types.Match]),
		RunningAverages:  gtk.NewKGTK[*types.RunningAverage]("RunningAverages", storeService, cdc, gtk.GetEmpty[types.RunningAverage], []string{Games24ValueKey, Votes24ValueKey}),
		Pools:            gtk.NewKGTK[*sdk.Coin]("Pools", storeService, cdc, gtk.GetEmpty[sdk.Coin], []string{PublicPoolKey, WinnersPoolKey, BalancersPoolKey}),
		images:           gtk.NewUintGTK[*types.Image]("Images", storeService, cdc, gtk.GetEmpty[types.Image]),
		servers:          gtk.NewUintGTK[*types.Server]("Servers", storeService, cdc, gtk.GetEmpty[types.Server]),
		encounters:       gtk.NewUintGTK[*types.Encounter]("Encounters", storeService, cdc, gtk.GetEmpty[types.Encounter]),
		users:            gtk.NewAddressGTK[*types.User]("Users", storeService, cdc, gtk.GetEmpty[types.User]),
		zealy:            gtk.NewStringGTK[*types.Zealy]("Zealy", storeService, cdc, gtk.GetEmpty[types.Zealy]),
		lastCardModified: gtk.NewSingleValueGenericTypeKeeper[*types.TimeStamp]("LastCardModified", storeService, cdc, gtk.GetEmpty[types.TimeStamp]),
		cardAuctionPrice: gtk.NewSingleValueGenericTypeKeeper[*sdk.Coin]("CardAuctionPrice", storeService, cdc, gtk.GetEmpty[sdk.Coin]),

		FeatureFlagModuleInstance: featureFlagKeeper.GetModuleInstance(types.ModuleName, []string{string(types.FeatureFlagName_Council), string(types.FeatureFlagName_Matches)}),
		BankKeeper:                bankKeeper,
	}
}

// GetAuthority returns the module's authority.
func (k Keeper) GetAuthority() string {
	return k.authority
}

// Logger returns a module-specific logger.
func (k Keeper) Logger() log.Logger {
	return k.logger.With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
