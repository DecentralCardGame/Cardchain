package types

const (
	// ModuleName defines the module name
	ModuleName = "cardchain"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_cardchain"
)

var (
	ParamsKey = []byte("p_cardchain")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	ProductDetailsKey      = "ProductDetails/value/"
	ProductDetailsCountKey = "ProductDetails/count/"
)
