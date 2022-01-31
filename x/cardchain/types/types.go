package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Here can constructors for types be placed, sice the files containing those
// types shall not be edited

// Module name for minting coins
const CoinsIssuerName = "cardchain"

func NewUser() User {
	return User{
		Alias:            "newPlayer",
		OwnedCardSchemes: []uint64{},
		OwnedCards:       []uint64{},
		VoteRights:       []*VoteRight{},
	}
}

func NewVoteRight(cardId uint64, expireBlock int64) VoteRight {
	return VoteRight{
		CardId:      cardId,
		ExpireBlock: expireBlock,
	}
}

func NewCard(owner sdk.AccAddress) Card {
	return Card{
		Owner:              owner.String(),
		Content:            []byte{},
		Image:              []byte{},
		Notes:              "",
		FullArt:            true,
		Status:             "scheme",
		VotePool:           sdk.NewInt64Coin("credits", 0),
		FairEnoughVotes:    0,
		OverpoweredVotes:   0,
		UnderpoweredVotes:  0,
		InappropriateVotes: 0,
		Nerflevel:          0,
	}
}

func CardNoB64FromCard(card Card) CardNoB64 {
	return CardNoB64{
		Owner:              card.Owner,
		Content:            string(card.Content),
		Image:              string(card.Image),
		Notes:              card.Notes,
		FullArt:            card.FullArt,
		Status:             card.Status,
		VotePool:           card.VotePool,
		FairEnoughVotes:    card.FairEnoughVotes,
		OverpoweredVotes:   card.OverpoweredVotes,
		UnderpoweredVotes:  card.UnderpoweredVotes,
		InappropriateVotes: card.InappropriateVotes,
		Nerflevel:          card.Nerflevel,
	}
}

func (m *Card) GetOwnerAddr() sdk.AccAddress {
	owner, err := sdk.AccAddressFromBech32(m.GetOwner())
	if err != nil {
		panic(err)
	}
	return owner
}

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
