package types

const (
	// ModuleName is the name of the module
	ModuleName = "cardservice"

	// StoreKey to be used when creating the KVStore
	StoreKey = ModuleName
	CardsStoreKey = "Cards"
	UsersStoreKey = "Users"
	InternalStoreKey = "Internal"

	// RouterKey is the module name router key
	RouterKey = ModuleName

	// QuerierRoute to be used for querierer msgs
	QuerierRoute = ModuleName
)
