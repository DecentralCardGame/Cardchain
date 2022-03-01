package types

// ResetVotes resets a cards votes
func (card *Card) ResetVotes() {
	card.Voters = []string{}
	card.FairEnoughVotes = 0
	card.OverpoweredVotes = 0
	card.UnderpoweredVotes = 0
	card.InappropriateVotes = 0
}
