package types

import (
	"fmt"
	"math"

	sdk "github.com/cosmos/cosmos-sdk/types"
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
		CollectionBaseUnitPrice:    sdk.NewInt64Coin("ucredits", 1000),
		ActiveCollectionsAmount:    3,
		CollectionCreationFee:      sdk.NewInt64Coin("ucredits", int64(5000*math.Pow(10, 6))),
		CollateralDeposit:					sdk.NewInt64Coin("ucredits", int64(50*math.Pow(10, 6))),
		WinnerReward:               sdk.NewInt64Coin("ucredits", 1),
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
		paramtypes.NewParamSetPair([]byte("CollectionBaseUnitPrice"), &p.CollectionBaseUnitPrice, validateCollectionBaseUnitPrice),
		paramtypes.NewParamSetPair([]byte("ActiveCollectionsAmount"), &p.ActiveCollectionsAmount, validateActiveCollectionsAmount),
		paramtypes.NewParamSetPair([]byte("CollectionCreationFee"), &p.CollectionCreationFee, validateCollectionCreationFee),
		paramtypes.NewParamSetPair([]byte("CollateralDeposit"), &p.CollateralDeposit, validateCollateralDeposit),
		paramtypes.NewParamSetPair([]byte("WinnerReward"), &p.WinnerReward, validateWinnerReward),
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
	v, ok := i.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("invalid VotingRightsExpirationTime: %d", v)
	}

	return nil
}

func validateCollectionSize(i interface{}) error {
	v, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("invalid CollectionSize: %d", v)
	}

	return nil
}

func validateCollectionBaseUnitPrice(i interface{}) error {
	v, ok := i.(sdk.Coin)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == sdk.NewInt64Coin("ucredits", 0) {
		return fmt.Errorf("invalid CollectionBaseUnitPrice: %d", v)
	}

	return nil
}

func validateActiveCollectionsAmount(i interface{}) error {
	v, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("invalid ActiveCollectionsAmount: %d", v)
	}

	return nil
}

func validateCollectionCreationFee(i interface{}) error {
	v, ok := i.(sdk.Coin)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == sdk.NewInt64Coin("ucredits", 0) {
		return fmt.Errorf("invalid CollectionCreationFee: %d", v)
	}

	return nil
}

func validateCollateralDeposit(i interface{}) error {
	v, ok := i.(sdk.Coin)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == sdk.NewInt64Coin("ucredits", 0) {
		return fmt.Errorf("invalid CollateralDeposit: %d", v)
	}

	return nil
}

func validateWinnerReward(i interface{}) error {
	v, ok := i.(sdk.Coin)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == sdk.NewInt64Coin("ucredits", 0) {
		return fmt.Errorf("invalid WinnerReward: %d", v)
	}

	return nil
}
