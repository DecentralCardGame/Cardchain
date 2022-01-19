package types

// Here can constructors for types be placed, sice the files containing those
// types shall not be edited

func NewUser() User {
	return User{
		Alias:	      		"newPlayer",
		OwnedCardSchemes: []uint64{},
		OwnedCards: 			[]uint64{},
		VoteRights: 			[]*VoteRight{},
	}
}
