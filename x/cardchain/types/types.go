package types

func NewUser() User {
	return User{
		Alias:            "newPlayer",
		Cards:            []uint64{},
		OwnedCardSchemes: []uint64{},
		OwnedPrototypes:  []uint64{},
		CouncilStatus:    CouncilStatus_unavailable,
		ReportMatches:    false,
		AirDrops:         &AirDrops{},
		EarlyAccess:      &EarlyAccess{},
	}
}
