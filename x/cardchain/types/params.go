package types

import (
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var _ paramtypes.ParamSet = (*Params)(nil)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams() Params {
	return Params{
		VotingRightsExpirationTime: 86000,
		CollectionSize:             5,
		CollectionPrice:            10,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams()
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair([]byte("VotingRightsExpirationTime"), &p.VotingRightsExpirationTime, validateVotingRightsExpirationTime),
		paramtypes.NewParamSetPair([]byte("CollectionSize"), &p.CollectionSize, validateCollectionSize),
		paramtypes.NewParamSetPair([]byte("CollectionPrice"), &p.CollectionPrice, validateCollectionPrice),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

func validateVotingRightsExpirationTime(i interface{}) error {
	return nil
}

func validateCollectionSize(i interface{}) error {
	return nil
}

func validateCollectionPrice(i interface{}) error {
	return nil
}
