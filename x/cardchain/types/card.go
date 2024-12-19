package types

import (
	errorsmod "cosmossdk.io/errors"
	"github.com/DecentralCardGame/cardobject/keywords"
)

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
