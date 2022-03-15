package types

const (
	// ModuleName defines the module name
	ModuleName = "cardchain"

	// StoreKey defines the primary module store key
	StoreKey                = ModuleName
	CardsStoreKey           = "Cards"
	UsersStoreKey           = "Users"
	MatchesStoreKey         = "Matches"
	CollectionsStoreKey     = "Collections"
	SellOffersStoreKey      = "SellOffers"
	PoolsStoreKey           = "Pools"
	CouncilsStoreKey        = "Councils"
	InternalStoreKey        = "Internal"
	RunningAveragesStoreKey = "RunningAverages"

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	// MemStoreKey = "mem_cardchain"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
