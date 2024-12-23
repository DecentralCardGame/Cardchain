package types

func NewMatchPlayer(addr string, cards []uint64, deck []uint64) *MatchPlayer {
	return &MatchPlayer{
		Addr:        addr,
		PlayedCards: cards,
		Deck:        deck,
		Confirmed:   false,
		Outcome:     Outcome_Aborted,
	}
}
