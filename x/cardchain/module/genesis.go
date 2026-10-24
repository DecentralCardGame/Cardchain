package cardchain

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/DecentralCardGame/cardchain/x/cardchain/keeper"
	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Validate before writing state. Never silently strip invalid card effects.
	if err := validateGenesisCards(genState.CardRecords); err != nil {
		panic(fmt.Errorf("invalid cardchain genesis: %w", err))
	}
	if err := genState.ValidateRunDeckReferences(); err != nil {
		panic(err)
	}
	// this line is used by starport scaffolding # genesis/module/init
	for id := range genState.Users {
		address, _ := sdk.AccAddressFromBech32(genState.Addresses[id])
		k.Users.Set(ctx, address, genState.Users[id])
	}
	for id, match := range genState.Matches {
		k.MatchK.Set(ctx, uint64(id), match)
	}
	for id, council := range genState.Councils {
		k.Councils.Set(ctx, uint64(id), council)
	}
	for id, sellOffer := range genState.SellOffers {
		k.SellOfferK.Set(ctx, uint64(id), sellOffer)
	}
	for id, image := range genState.Images {
		k.Images.Set(ctx, uint64(id), image)
	}
	for id, server := range genState.Servers {
		k.Servers.Set(ctx, uint64(id), server)
	}
	for id, pool := range genState.Pools {
		k.Pools.Set(ctx, k.Pools.KeyWords[id], pool)
	}
	for idx, average := range genState.RunningAverages {
		k.RunningAverages.Set(ctx, k.RunningAverages.KeyWords[idx], average)
	}
	for idx, encounter := range genState.Encounters {
		k.EncounterK.Set(ctx, uint64(idx), encounter)
	}
	if genState.CardAuctionPrice.Denom != "" {
		k.CardAuctionPrice.Set(ctx, &genState.CardAuctionPrice)
	}
	k.LastCardModified.Set(ctx, &genState.LastCardModified)
	for _, zealy := range genState.Zealys {
		k.Zealy.Set(ctx, zealy.ZealyId, zealy)
	}
	for _, upgradeFactor := range genState.UpgradeFactors {
		k.UpgradeFactorK.Set(ctx, upgradeFactor.Name, upgradeFactor)
	}
	k.Logger().Info("reading cards with id:")
	for currId, record := range genState.CardRecords {
		k.CardK.Set(ctx, uint64(currId), record)
	}
	if err := k.InitRunDeckGenesis(ctx, genState.RunDecks); err != nil {
		panic(err)
	}
	if genState.Params.AirDropValue.Denom == "" {
		defaultParams := types.DefaultParams()
		genState.Params.AirDropValue = defaultParams.AirDropValue
		genState.Params.AirDropMaxBlockHeight = defaultParams.AirDropMaxBlockHeight
	}
	if genState.Params.MatchWorkerDelay == 0 {
		genState.Params.MatchWorkerDelay = types.DefaultMatchWorkerDelay
	}

	if err := k.SetParams(ctx, genState.Params); err != nil {
		panic(err)
	}

	for id, set := range genState.Sets {
		if set.Status == types.SetStatus_active || set.Status == types.SetStatus_finalized {
			set.ContributorsDistribution = k.GetContributorDistribution(ctx, *set)
			set.Rarities = k.GetCardRaritiesInSet(ctx, set)
		}
		k.SetK.Set(ctx, uint64(id), set)
	}
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	genesis.RunDecks = k.ExportRunDeckGenesis(ctx)
	genesis.Params = k.GetParams(ctx)
	genesis.CardAuctionPrice = *k.CardAuctionPrice.Get(ctx)
	genesis.LastCardModified = *k.LastCardModified.Get(ctx)
	genesis.SellOffers = k.SellOfferK.GetAll(ctx)
	genesis.Pools = k.Pools.GetAll(ctx)
	genesis.CardRecords = k.CardK.GetAll(ctx)
	genesis.Matches = k.MatchK.GetAll(ctx)
	genesis.Councils = k.Councils.GetAll(ctx)
	genesis.RunningAverages = k.RunningAverages.GetAll(ctx)
	genesis.Sets = k.SetK.GetAll(ctx)
	genesis.Images = k.Images.GetAll(ctx)
	genesis.Servers = k.Servers.GetAll(ctx)
	users, accAddresses := k.GetAllUsers(ctx)
	genesis.Zealys = k.Zealy.GetAll(ctx)
	genesis.Encounters = k.EncounterK.GetAll(ctx)
	genesis.UpgradeFactors = k.UpgradeFactorK.GetAll(ctx)
	var addresses []string
	for _, addr := range accAddresses {
		addresses = append(addresses, addr.String())
	}
	genesis.Users = users
	genesis.Addresses = addresses
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
