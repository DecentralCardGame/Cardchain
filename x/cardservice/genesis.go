package cardservice

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/DecentralCardGame/cardobject/keywords"
)

type GenesisState struct {
	CardRecords 				[]Card 						`json:"card_records"`
	Users 							[]User 						`json:"users"`
	SdkAddresses 				[]sdk.AccAddress	`json:"addresses"`
	LastCardSchemeId 		uint64						`json:"last_card_scheme_id"`
}

func NewGenesisState(cardRecords []Card, users []User, addresses []sdk.AccAddress, lastCardSchemeId uint64) GenesisState {
	return GenesisState{
		CardRecords: cardRecords,
		Users: users,
		SdkAddresses: addresses,
		LastCardSchemeId: lastCardSchemeId,
	}
}

func ValidateGenesis(data GenesisState) error {
	for _, record := range data.CardRecords {
		if record.Owner == nil {
			return fmt.Errorf("invalid CardRecord: Owner: %s. Error: Missing Owner", record.Owner)
		}
		if record.Content == nil {
			return fmt.Errorf("invalid CardRecord: Content: %s. Error: Missing Content", record.Content)
		}
		if record.Status == "" {
			return fmt.Errorf("invalid CardRecord: Status: %s. Error: Missing Status", record.Status)
		}
	}
	return nil
}

func DefaultGenesisState() GenesisState {
	return GenesisState{
		CardRecords: 					[]Card{},
		Users: 								[]User{},
		SdkAddresses: 				[]sdk.AccAddress{},
		LastCardSchemeId: 		0,
	}
}

func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) {
	for id := range data.Users {
		keeper.SetUser(ctx, data.SdkAddresses[id], data.Users[id])
	}
	fmt.Println("reading cards with id:")
	for id, record := range data.CardRecords {
		fmt.Println(id)
		lastId := keeper.GetLastCardSchemeId(ctx)
		currId := lastId + 1

		_, err := keywords.Unmarshal(record.Content)
		if err != nil {
			fmt.Println(err.Error())
		}

		keeper.SetLastCardSchemeId(ctx, currId)
		keeper.SetCard(ctx, currId, record)

		//if (record.Status == "scheme") {
			//keeper.AddOwnedCardScheme(ctx, currId, record.Owner)
		//}

	}
}

func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	records := k.GetAllCards(ctx)
	users, addresses := k.GetAllUsers(ctx)
	return GenesisState{
		CardRecords: 			records,
		Users: 						users,
		SdkAddresses: 		addresses,
		LastCardSchemeId:	k.GetLastCardSchemeId(ctx),
	}
}
