package keeper

import (
	"fmt"

	"cosmossdk.io/core/store"
	"cosmossdk.io/log"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	gtk "github.com/DecentralCardGame/cardchain/types/generic_type_keeper"
	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
)

type (
	Keeper struct {
		cdc          codec.BinaryCodec
		storeService store.KVStoreService
		logger       log.Logger

		Cards           gtk.GenericTypeKeeper[*types.Card]
		Councils        gtk.GenericTypeKeeper[*types.Council]
		SellOffers      gtk.GenericTypeKeeper[*types.SellOffer]
		Sets            gtk.GenericTypeKeeper[*types.Set]
		Matches         gtk.GenericTypeKeeper[*types.Match]
		Servers         gtk.GenericTypeKeeper[*types.Server]
		RunningAverages gtk.KeywordedGenericTypeKeeper[*types.RunningAverage]
		Pools           gtk.KeywordedGenericTypeKeeper[*sdk.Coin]
		Images          gtk.GenericTypeKeeper[*types.Image]
		Encounters      gtk.GenericTypeKeeper[*types.Encounter]

		// the address capable of executing a MsgUpdateParams message. Typically, this
		// should be the x/gov module account.
		authority string

		BankKeeper types.BankKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeService store.KVStoreService,
	logger log.Logger,
	authority string,
	bankKeeper types.BankKeeper,

) Keeper {
	if _, err := sdk.AccAddressFromBech32(authority); err != nil {
		panic(fmt.Sprintf("invalid authority address: %s", authority))
	}

	return Keeper{
		cdc:          cdc,
		storeService: storeService,
		authority:    authority,
		logger:       logger,

		Cards:           gtk.NewGTK[*types.Card]("Cards", storeService, cdc, gtk.GetEmpty[types.Card]),
		Councils:        gtk.NewGTK[*types.Council]("Councils", storeService, cdc, gtk.GetEmpty[types.Council]),
		SellOffers:      gtk.NewGTK[*types.SellOffer]("SellOffers", storeService, cdc, gtk.GetEmpty[types.SellOffer]),
		Sets:            gtk.NewGTK[*types.Set]("Sets", storeService, cdc, gtk.GetEmpty[types.Set]),
		Matches:         gtk.NewGTK[*types.Match]("Matches", storeService, cdc, gtk.GetEmpty[types.Match]),
		RunningAverages: gtk.NewKGTK[*types.RunningAverage]("RunningAverages", storeService, cdc, gtk.GetEmpty[types.RunningAverage], []string{Games24ValueKey, Votes24ValueKey}),
		Pools:           gtk.NewKGTK[*sdk.Coin]("Pools", storeService, cdc, gtk.GetEmpty[sdk.Coin], []string{PublicPoolKey, WinnersPoolKey, BalancersPoolKey}),
		Images:          gtk.NewGTK[*types.Image]("Images", storeService, cdc, gtk.GetEmpty[types.Image]),
		Servers:         gtk.NewGTK[*types.Server]("Servers", storeService, cdc, gtk.GetEmpty[types.Server]),
		Encounters:      gtk.NewGTK[*types.Encounter]("Encounters", storeService, cdc, gtk.GetEmpty[types.Encounter]),

		BankKeeper: bankKeeper,
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
