package types

const (
	// ModuleName defines the module name
	ModuleName = "cardchain"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_cardchain"

	// GovModuleName duplicates the gov module's name to avoid a dependency with x/gov.
	// It should be synced with the gov module's name if it is ever changed.
	// See: https://github.com/cosmos/cosmos-sdk/blob/v0.52.0-beta.2/x/gov/types/keys.go#L9
	GovModuleName = "gov"
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
