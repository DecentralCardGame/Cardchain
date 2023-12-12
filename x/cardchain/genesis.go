package cardchain

import (
	"fmt"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/keeper"
	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	"github.com/DecentralCardGame/cardobject/keywords"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	for id := range genState.Users {
		address, _ := sdk.AccAddressFromBech32(genState.Addresses[id])
		k.SetUser(ctx, address, *genState.Users[id])
	}
	for id, match := range genState.Matches {
		k.Matches.Set(ctx, uint64(id), match)
	}
	for id, council := range genState.Councils {
		k.Councils.Set(ctx, uint64(id), council)
	}
	for id, sellOffer := range genState.SellOffers {
		k.SellOffers.Set(ctx, uint64(id), sellOffer)
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
	if genState.CardAuctionPrice.Denom != "" {
		k.SetCardAuctionPrice(ctx, genState.CardAuctionPrice)
	}
	if genState.LastCardModified != nil {
		k.SetLastCardModified(ctx, *genState.LastCardModified)
	}
	k.Logger(ctx).Info("reading cards with id:")
	for currId, record := range genState.CardRecords {
		_, err := keywords.Unmarshal(record.Content)
		if err != nil {
			k.Logger(ctx).Error(fmt.Sprintf("%d :\n\t%s\n\t%s\n-----", currId, err.Error(), record.Content))
		}

		k.Cards.Set(ctx, uint64(currId), record)
	}
	k.Logger(ctx).Info("Params", genState.Params)
	if genState.Params.AirDropValue.Denom == "" {
		defaultParams := types.DefaultParams()
		genState.Params.AirDropValue = defaultParams.AirDropValue
		genState.Params.AirDropMaxBlockHeight = defaultParams.AirDropMaxBlockHeight
	}
	if genState.Params.MatchWorkerDelay == 0 {
		genState.Params.MatchWorkerDelay = types.DefaultMatchWorkerDelay
	}
	k.SetParams(ctx, genState.Params)
	for id, set := range genState.Sets {
		if set.Status == types.CStatus_active || set.Status == types.CStatus_finalized {
			set.ContributorsDistribution = k.GetContributorDistribution(ctx, set)
		}
		k.Sets.Set(ctx, uint64(id), set)
	}
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	// this line is used by starport scaffolding # genesis/module/export
	params := k.GetParams(ctx)
	// params := types.DefaultParams()
	cardAuctionPrice := k.GetCardAuctionPrice(ctx)
	lastCardModified := k.GetLastCardModified(ctx)
	sellOffers := k.SellOffers.GetAll(ctx)
	pools := k.Pools.GetAll(ctx)
	records := k.Cards.GetAll(ctx)
	matches := k.Matches.GetAll(ctx)
	councils := k.Councils.GetAll(ctx)
	runningAverages := k.RunningAverages.GetAll(ctx)
	sets := k.Sets.GetAll(ctx)
	images := k.Images.GetAll(ctx)
	servers := k.Servers.GetAll(ctx)
	users, accAddresses := k.GetAllUsers(ctx)
	var addresses []string
	for _, addr := range accAddresses {
		addresses = append(addresses, addr.String())
	}
	return &types.GenesisState{
		Params:           params,
		CardRecords:      records,
		Users:            users,
		Matches:          matches,
		Sets:             sets,
		SellOffers:       sellOffers,
		Pools:            pools,
		Councils:         councils,
		Addresses:        addresses,
		CardAuctionPrice: cardAuctionPrice,
		Images:           images,
		RunningAverages:  runningAverages,
		Servers:          servers,
		LastCardModified: &lastCardModified,
	}
}
