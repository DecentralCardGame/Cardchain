package types

import (
	math "math"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
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
	return paramtypes.ParamSetPairs{}
}

// Validate validates the set of params
func (p Params) Validate() error {
	return nil
}
