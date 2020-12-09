package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
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
