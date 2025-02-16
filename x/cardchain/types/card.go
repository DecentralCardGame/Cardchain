package types

import (
	errorsmod "cosmossdk.io/errors"
	"github.com/DecentralCardGame/cardobject/keywords"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewCard Constructor for Card
func NewCard(owner sdk.AccAddress) Card {
	return Card{
		Owner:    owner.String(),
		Notes:    "",
		FullArt:  true,
		Status:   CardStatus_scheme,
		VotePool: sdk.NewInt64Coin("ucredits", 0),
	}
}

// ResetVotes resets a cards votes
func (card *Card) ResetVotes() {
	card.Voters = []string{}
	card.FairEnoughVotes = 0
	card.OverpoweredVotes = 0
	card.UnderpoweredVotes = 0
	card.InappropriateVotes = 0
}

func (card *Card) GetCardObj() (*keywords.Card, error) {
	cardobj, err := keywords.Unmarshal(card.Content)
	if err != nil {
		errorsmod.Wrapf(ErrCardobject, "Invalid cardobject: %s", err)
	}

	return cardobj, nil
}
