package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		ProductDetailsList: []ProductDetails{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in productDetails
	productDetailsIdMap := make(map[uint64]bool)
	productDetailsCount := gs.GetProductDetailsCount()
	for _, elem := range gs.ProductDetailsList {
		if _, ok := productDetailsIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for productDetails")
		}
		if elem.Id >= productDetailsCount {
			return fmt.Errorf("productDetails id should be lower or equal than the last id")
		}
		productDetailsIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
