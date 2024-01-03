package types

import (
	"fmt"
	"math"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var _ paramtypes.ParamSet = (*Params)(nil)

const DefaultMatchWorkerDelay = 170

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams() Params {
	return Params{
		VotingRightsExpirationTime:      86000,
		SetSize:                         24,
		SetPrice:                        sdk.NewInt64Coin("ucredits", 10000000),
		ActiveSetsAmount:                3,
		SetCreationFee:                  sdk.NewInt64Coin("ucredits", int64(5000*math.Pow(10, 6))),
		CollateralDeposit:               sdk.NewInt64Coin("ucredits", int64(50*math.Pow(10, 6))),
		TrialVoteReward:                 sdk.NewInt64Coin("ucredits", int64(math.Pow(10, 6))),
		WinnerReward:                    int64(math.Pow(10, 6)),
		VotePoolFraction:                int64(math.Pow(10, 6)),
		VotingRewardCap:                 int64(math.Pow(10, 6)),
		HourlyFaucet:                    sdk.NewInt64Coin("ucredits", int64(50*math.Pow(10, 6))),
		InflationRate:                   "1.1", // TODO: Also make this a fixed point number
		RaresPerPack:                    1,
		CommonsPerPack:                  9,
		UnCommonsPerPack:                3,
		TrialPeriod:                     14 * 24 * 500,
		GameVoteRatio:                   20, // This is a fixed point number and will be devided by 100 when used
		CardAuctionPriceReductionPeriod: 20,
		AirDropValue:                    sdk.NewInt64Coin("ubpf", int64(5*math.Pow(10, 6))),
		AirDropMaxBlockHeight:           5000000,
		MatchWorkerDelay:                DefaultMatchWorkerDelay,
		RareDropRatio:                   150,
		ExceptionalDropRatio:            50,
		UniqueDropRatio:                 1,
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
		paramtypes.NewParamSetPair([]byte("SetSize"), &p.SetSize, validateSetSize),
		paramtypes.NewParamSetPair([]byte("SetPrice"), &p.SetPrice, validateSetPrice),
		paramtypes.NewParamSetPair([]byte("TrialVoteReward"), &p.TrialVoteReward, validateSetPrice),
		paramtypes.NewParamSetPair([]byte("ActiveSetsAmount"), &p.ActiveSetsAmount, validateActiveSetsAmount),
		paramtypes.NewParamSetPair([]byte("SetCreationFee"), &p.SetCreationFee, validateSetCreationFee),
		paramtypes.NewParamSetPair([]byte("CollateralDeposit"), &p.CollateralDeposit, validateCollateralDeposit),
		paramtypes.NewParamSetPair([]byte("WinnerReward"), &p.WinnerReward, validateWinnerReward),
		paramtypes.NewParamSetPair([]byte("VotePoolFraction"), &p.VotePoolFraction, validateVoterReward),
		paramtypes.NewParamSetPair([]byte("VotingRewardCap"), &p.VotingRewardCap, validateVoterReward),
		paramtypes.NewParamSetPair([]byte("HourlyFaucet"), &p.HourlyFaucet, validateHourlyFaucet),
		paramtypes.NewParamSetPair([]byte("InflationRate"), &p.InflationRate, validateInflationRate),
		paramtypes.NewParamSetPair([]byte("RaresPerPack"), &p.RaresPerPack, validatePerPack),
		paramtypes.NewParamSetPair([]byte("CommonsPerPack"), &p.CommonsPerPack, validatePerPack),
		paramtypes.NewParamSetPair([]byte("UnCommonsPerPack"), &p.UnCommonsPerPack, validatePerPack),
		paramtypes.NewParamSetPair([]byte("TrialPeriod"), &p.TrialPeriod, validateTrialPeriod),
		paramtypes.NewParamSetPair([]byte("GameVoteRatio"), &p.GameVoteRatio, validateGameVoteRatio),
		paramtypes.NewParamSetPair([]byte("CardAuctionPriceReductionPeriod"), &p.CardAuctionPriceReductionPeriod, validateCardAuctionPriceReductionPeriod),
		paramtypes.NewParamSetPair([]byte("AirDropValue"), &p.AirDropValue, validateAirDropValue),
		paramtypes.NewParamSetPair([]byte("AirDropMaxBlockHeight"), &p.AirDropMaxBlockHeight, validateAirDropMaxBlockHeight),
		paramtypes.NewParamSetPair([]byte("MatchWorkerDelay"), &p.MatchWorkerDelay, validateMatchWorkerDelay),
		paramtypes.NewParamSetPair([]byte("RareDropRatio"), &p.RareDropRatio, validateRareDropRatio),
		paramtypes.NewParamSetPair([]byte("ExceptionalDropRatio"), &p.ExceptionalDropRatio, validateExceptionalDropRatio),
		paramtypes.NewParamSetPair([]byte("UniqueDropRatio"), &p.UniqueDropRatio, validateUniqueDropRatio),
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

func validateSetSize(i interface{}) error {
	v, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("invalid SetSize: %d", v)
	}

	return nil
}

func validateSetPrice(i interface{}) error {
	v, ok := i.(sdk.Coin)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == sdk.NewInt64Coin("ucredits", 0) {
		return fmt.Errorf("invalid SetPrice: %v", v)
	}

	return nil
}

func validateActiveSetsAmount(i interface{}) error {
	v, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("invalid ActiveSetsAmount: %d", v)
	}

	return nil
}

func validateTrialPeriod(i interface{}) error {
	v, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("invalid TrialPeriod: %d", v)
	}

	return nil
}

func validatePerPack(i interface{}) error {
	v, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("invalid number per pack: %d", v)
	}

	return nil
}

func validateGameVoteRatio(i interface{}) error {
	v, ok := i.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("invalid GameVoteRatio: %d", v)
	}

	return nil
}

func validateCardAuctionPriceReductionPeriod(i interface{}) error {
	v, ok := i.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("invalid CardAuctionPriceReductionPeriod: %d", v)
	}

	return nil
}

func validateSetCreationFee(i interface{}) error {
	v, ok := i.(sdk.Coin)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == sdk.NewInt64Coin("ucredits", 0) {
		return fmt.Errorf("invalid SetCreationFee: %v", v)
	}

	return nil
}

func validateAirDropValue(i interface{}) error {
	v, ok := i.(sdk.Coin)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == sdk.NewInt64Coin("ucredits", 0) {
		return fmt.Errorf("invalid AirDropValue: %v", v)
	}

	return nil
}

func validateCollateralDeposit(i interface{}) error {
	v, ok := i.(sdk.Coin)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == sdk.NewInt64Coin("ucredits", 0) {
		return fmt.Errorf("invalid CollateralDeposit: %v", v)
	}

	return nil
}

func validateWinnerReward(i interface{}) error {
	v, ok := i.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("invalid WinnerReward: %d", v)
	}

	return nil
}

func validateVoterReward(i interface{}) error {
	v, ok := i.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("invalid VoterReward: %d", v)
	}

	return nil
}

func validateAirDropMaxBlockHeight(i interface{}) error {
	v, ok := i.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v < 0 {
		return fmt.Errorf("invalid AirDropMaxBlockHeight: %d", v)
	}

	return nil
}

func validateHourlyFaucet(i interface{}) error {
	v, ok := i.(sdk.Coin)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == sdk.NewInt64Coin("ucredits", 0) {
		return fmt.Errorf("invalid HourlyFaucet: %v", v)
	}

	return nil
}

func validateInflationRate(i interface{}) error {
	v, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	_, err := strconv.ParseFloat(v, 8)
	if err != nil {
		return fmt.Errorf("invalid parameter: %d", err)
	}
	return nil
}

func validateMatchWorkerDelay(i interface{}) error {
	_, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	return nil
}

func validateRareDropRatio(i interface{}) error {
	_, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	return nil
}

func validateExceptionalDropRatio(i interface{}) error {
	_, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	return nil
}

func validateUniqueDropRatio(i interface{}) error {
	_, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	return nil
}
