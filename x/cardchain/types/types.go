package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewMatchPlayer(addr string, cards []uint64) *MatchPlayer {
	return &MatchPlayer{
		Addr:        addr,
		PlayedCards: cards,
		Confirmed:   false,
		Outcome:     Outcome_Aborted,
	}
}

func NewIgnoreMatches() IgnoreMatches {
	return IgnoreMatches{
		false,
	}
}

func NewIgnoreSellOffers() IgnoreSellOffers {
	return IgnoreSellOffers{
		false,
		false,
	}
}

// NewUser Constructor for User
func NewUser() User {
	return User{
		Alias:            "newPlayer",
		Cards:            []uint64{},
		OwnedCardSchemes: []uint64{},
		OwnedPrototypes:  []uint64{},
		VoteRights:       []*VoteRight{},
		CouncilStatus:    CouncilStatus_unavailable,
		ReportMatches:    false,
	}
}

// NewVoteRight Constructor for VoteRights
func NewVoteRight(cardId uint64, expireBlock int64) VoteRight {
	return VoteRight{
		CardId:      cardId,
		ExpireBlock: expireBlock,
	}
}

// NewCard Constructor for Card
func NewCard(owner sdk.AccAddress) Card {
	return Card{
		Owner:    owner.String(),
		Notes:    "",
		FullArt:  true,
		Status:   Status_scheme,
		VotePool: sdk.NewInt64Coin("ucredits", 0),
	}
}

// NewVotingResults Constructor for VoteResults
func NewVotingResults() VotingResults {
	return VotingResults{
		TotalVotes:              0,
		TotalFairEnoughVotes:    0,
		TotalOverpoweredVotes:   0,
		TotalUnderpoweredVotes:  0,
		TotalInappropriateVotes: 0,
		CardResults:             []*VotingResult{},
		Notes:                   "none",
	}
}
