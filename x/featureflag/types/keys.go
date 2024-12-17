package types

const (
	// ModuleName defines the module name
	ModuleName = "featureflag"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_featureflag"
)

var (
	ParamsKey = []byte("p_featureflag")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
