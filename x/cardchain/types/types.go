package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewMatchPlayer(addr string, cards []uint64, deck []uint64) *MatchPlayer {
	return &MatchPlayer{
		Addr:        addr,
		PlayedCards: cards,
		Deck:        deck,
		Confirmed:   false,
		Outcome:     Outcome_Aborted,
	}
}

func NewBoosterPack(ctx sdk.Context, setId uint64, raritiesPerPack []uint64, dropRatiosPerPack []uint64) BoosterPack {
	return BoosterPack{
		SetId:             setId,
		RaritiesPerPack:   raritiesPerPack,
		TimeStamp:         ctx.BlockHeight(),
		DropRatiosPerPack: dropRatiosPerPack,
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

func NewIgnoreCouncils() IgnoreCouncils {
	return IgnoreCouncils{
		false,
		false,
	}
}

// NewUser Constructor for User
func NewUser() User {
	return User{
		Alias:                "newPlayer",
		Cards:                []uint64{},
		OwnedCardSchemes:     []uint64{},
		OwnedPrototypes:      []uint64{},
		CouncilParticipation: NewCouncilParticipation(),
		ReportMatches:        false,
		AirDrops:             &AirDrops{},
	}
}

func NewCouncilParticipation() *CouncilParticipation {
	return &CouncilParticipation{
		Status:  CouncilStatus_unavailable,
		Council: 0,
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
