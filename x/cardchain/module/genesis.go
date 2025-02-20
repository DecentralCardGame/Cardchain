package cardchain

import (
	"encoding/json"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/DecentralCardGame/cardchain/x/cardchain/keeper"
	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/keywords"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
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
		k.Encounterk.Set(ctx, uint64(idx), encounter)
	}
	if genState.CardAuctionPrice.Denom != "" {
		k.CardAuctionPrice.Set(ctx, &genState.CardAuctionPrice)
	}
	k.LastCardModified.Set(ctx, &genState.LastCardModified)
	for _, zealy := range genState.Zealys {
		k.Zealy.Set(ctx, zealy.ZealyId, zealy)
	}
	k.Logger().Info("reading cards with id:")
	for currId, record := range genState.CardRecords {
		if len(record.Content) != 0 {
			_, err := keywords.Unmarshal(record.Content)

			if err != nil {
				k.Logger().Info(fmt.Sprintf("Failed to read %d :\n\t%s\n\t%s\n-----", currId, err.Error(), record.Content))
				var card keywords.Card
				json.Unmarshal(record.Content, &card)

				switch card.GetType() {
				case cardobject.ENTITYTYPE:
					card.Entity.Abilities.Clear()
				case cardobject.PLACETYPE:
					card.Place.Abilities.Clear()
				case cardobject.HEADQUARTERTYPE:
					card.Headquarter.Abilities.Clear()
				case cardobject.ACTIONTYPE:
					card.Action.Effects.Clear()
				}
				jsonBytes, _ := json.Marshal(card)
				record.Content = jsonBytes
			}
		}

		k.CardK.Set(ctx, uint64(currId), record)
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
	genesis.Encounters = k.Encounterk.GetAll(ctx)
	var addresses []string
	for _, addr := range accAddresses {
		addresses = append(addresses, addr.String())
	}
	genesis.Users = users
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
