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
		k.SetMatch(ctx, uint64(id), *match)
	}
	for id, council := range genState.Councils {
		k.SetCouncil(ctx, uint64(id), *council)
	}
	for id, collection := range genState.Collections {
		k.SetCollection(ctx, uint64(id), *collection)
	}
	for id, sellOffer := range genState.SellOffers {
		k.SetSellOffer(ctx, uint64(id), *sellOffer)
	}
	for id, pool := range genState.Pools {
		k.SetPool(ctx, k.PoolKeys[id], pool)
	}
	for key, value := range genState.GeneralValues {
		k.SetGeneralValue(ctx, key, value)
	}
	if genState.CardAuctionPrice.Denom != "" {
		k.SetCardAuctionPrice(ctx, genState.CardAuctionPrice)
	}
	fmt.Println("reading cards with id:")
	for currId, record := range genState.CardRecords {
		_, err := keywords.Unmarshal(record.Content)
		if err != nil {
			fmt.Println(currId, ":")
			fmt.Println(err.Error())
			fmt.Println(string(record.Content))
			fmt.Println("-----")
		}

		k.SetCard(ctx, uint64(currId), *record)
	}
	fmt.Println("Params", genState.Params)
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	// this line is used by starport scaffolding # genesis/module/export
	params := k.GetParams(ctx)
	cardAuctionPrice := k.GetCardAuctionPrice(ctx)
	generalValues := k.GetAllGeneralValues(ctx)
	sellOffers := k.GetAllSellOffers(ctx)
	pools := k.GetAllPools(ctx)
	records := k.GetAllCards(ctx)
	matches := k.GetAllMatches(ctx)
	councils := k.GetAllCouncils(ctx)
	collections := k.GetAllCollections(ctx)
	users, accAddresses := k.GetAllUsers(ctx)
	var addresses []string
	for _, addr := range accAddresses {
		addresses = append(addresses, addr.String())
	}
	return &types.GenesisState{
		Params:           params,
		CardRecords:      records,
		GeneralValues:    generalValues,
		Users:            users,
		Matches:          matches,
		Collections:      collections,
		SellOffers:       sellOffers,
		Pools:            pools,
		Councils:         councils,
		Addresses:        addresses,
		CardAuctionPrice: cardAuctionPrice,
	}
}
