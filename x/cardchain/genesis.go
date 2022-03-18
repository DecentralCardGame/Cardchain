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
	for id, collection := range genState.Collections {
		k.Collections.Set(ctx, uint64(id), collection)
	}
	for id, sellOffer := range genState.SellOffers {
		k.SellOffers.Set(ctx, uint64(id), sellOffer)
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
	k.Logger(ctx).Info("reading cards with id:")
	for currId, record := range genState.CardRecords {
		_, err := keywords.Unmarshal(record.Content)
		if err != nil {
			k.Logger(ctx).Error(fmt.Sprintf("%d :\n\t%s\n\t%s\n-----", currId, err.Error(), record.Content))
		}

		k.Cards.Set(ctx, uint64(currId), record)
	}
	k.Logger(ctx).Info("Params", genState.Params)
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	// this line is used by starport scaffolding # genesis/module/export
	params := k.GetParams(ctx)
	cardAuctionPrice := k.GetCardAuctionPrice(ctx)
	sellOffers := k.SellOffers.GetAll(ctx)
	pools := k.Pools.GetAll(ctx)
	records := k.Cards.GetAll(ctx)
	matches := k.Matches.GetAll(ctx)
	councils := k.Councils.GetAll(ctx)
	runningAverages := k.RunningAverages.GetAll(ctx)
	collections := k.Collections.GetAll(ctx)
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
		Collections:      collections,
		SellOffers:       sellOffers,
		Pools:            pools,
		Councils:         councils,
		Addresses:        addresses,
		CardAuctionPrice: cardAuctionPrice,
		RunningAverages:  runningAverages,
	}
}
