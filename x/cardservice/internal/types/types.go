package types

import (
	//"fmt"
	//"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type User struct {
	Alias      			 string
	OwnedCardSchemes []uint64
	OwnedCards			 []uint64
	VoteRights 			 []VoteRight
}

func NewUser() User {
	return User{
		Alias:	      		"newPlayer",
		OwnedCardSchemes: []uint64{},
		OwnedCards: 			[]uint64{},
		VoteRights: 			[]VoteRight{},
	}
}

type Card struct {
	Owner              sdk.AccAddress
	Content            []byte
	Image							 []byte
	FullArt						 string
	Notes							 string
	Status             string
	VotePool           sdk.Coin
	FairEnoughVotes    uint64
	OverpoweredVotes   uint64
	UnderpoweredVotes  uint64
	InappropriateVotes uint64
	Nerflevel          int64
}

func NewCard(owner sdk.AccAddress) Card {
	return Card{
		Owner:              owner,
		Content:            []byte{},
		Image:							[]byte{},
		Notes:							"",
		FullArt:						"",
		Status:             "scheme",
		VotePool:           sdk.NewInt64Coin("credits", 0),
		FairEnoughVotes:    0,
		OverpoweredVotes:   0,
		UnderpoweredVotes:  0,
		InappropriateVotes: 0,
		Nerflevel:          0,
	}
}

type CardNoB64 struct {
	Owner              sdk.AccAddress
	Content            string
	Image							 string
	Notes							 string
	FullArt						 string
	Status             string
	VotePool           sdk.Coin
	FairEnoughVotes    uint64
	OverpoweredVotes   uint64
	UnderpoweredVotes  uint64
	InappropriateVotes uint64
	Nerflevel          int64
}

func CardNoB64FromCard(card Card) CardNoB64 {
	return CardNoB64{
		Owner:              card.Owner,
		Content:            string(card.Content),
		Image:							string(card.Image),
		Notes:							card.Notes,
		FullArt:						card.FullArt,
		Status:             card.Status,
		VotePool:           card.VotePool,
		FairEnoughVotes:    card.FairEnoughVotes,
		OverpoweredVotes:   card.OverpoweredVotes,
		UnderpoweredVotes:  card.UnderpoweredVotes,
		InappropriateVotes: card.InappropriateVotes,
		Nerflevel:          card.Nerflevel,
	}
}

type VoteRight struct {
	CardId      uint64
	ExpireBlock int64
}

func NewVoteRight(cardId uint64, expireBlock int64) VoteRight {
	return VoteRight{
		CardId:      cardId,
		ExpireBlock: expireBlock,
	}
}

type VotingResult struct {
	CardId      				uint64
	FairEnoughVotes   	uint64
	OverpoweredVotes  	uint64
	UnderpoweredVotes 	uint64
	InappropriateVotes	uint64
	Result							string
}

type VotingResults struct {
	TotalVotes							uint64
	TotalFairEnoughVotes		uint64
	TotalOverpoweredVotes		uint64
	TotalUnderpoweredVotes	uint64
	TotalInappropriateVotes	uint64
	CardResults	 						[]VotingResult
	Notes										string
}

func NewVotingResults() VotingResults {
	return VotingResults{
		TotalVotes:							0,
		TotalFairEnoughVotes:		0,
		TotalOverpoweredVotes:	0,
		TotalUnderpoweredVotes:	0,
		TotalInappropriateVotes:0,
		CardResults: []VotingResult{},
		Notes: "none",
	}
}
