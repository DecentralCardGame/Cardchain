package types

import (
	//"fmt"
	//"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Here can constructors for types be placed, sice the files containing those
// types shall not be edited

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

// type Card struct {
// 	Owner              sdk.AccAddress
// 	Content            []byte
// 	Image							 []byte
// 	FullArt						 string
// 	Notes							 string
// 	Status             string
// 	VotePool           sdk.Coin
// 	FairEnoughVotes    uint64
// 	OverpoweredVotes   uint64
// 	UnderpoweredVotes  uint64
// 	InappropriateVotes uint64
// 	Nerflevel          int64
// }

func NewCard(owner sdk.AccAddress) Card {
	return Card{
		Owner:              owner,
		Content:            []byte{},
		Image:              []byte{},
		Notes:              "",
		FullArt:            "",
		Status:             "scheme",
		VotePool:           sdk.NewInt64Coin("credits", 0),
		FairEnoughVotes:    0,
		OverpoweredVotes:   0,
		UnderpoweredVotes:  0,
		InappropriateVotes: 0,
		Nerflevel:          0,
	}
}
