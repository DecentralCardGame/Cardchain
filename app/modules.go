package app

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	"github.com/cosmos/cosmos-sdk/x/gov"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/cosmos/cosmos-sdk/x/mint"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/cosmos/cosmos-sdk/x/staking"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

// BankModule defines a custom wrapper around the x/bank module's AppModuleBasic
// implementation to provide custom default genesis state.
type BankModule struct {
	bank.AppModuleBasic
}

// DefaultGenesis returns custom x/bank module genesis state.
func (BankModule) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	Metadata := banktypes.Metadata{
		Name:        "Burning Pitchfork",
		Symbol:      "BPF",
		Description: "The native staking token of the CrowdControlNetwork.",
		DenomUnits: []*banktypes.DenomUnit{
			{"ubpf", uint32(0), []string{"microbpf"}},
			{"mbpf", uint32(3), []string{"millibpf"}},
			{"bpf", uint32(6), nil},
		},
		Base:    "ubpf",
		Display: "bpf",
	}

	genState := banktypes.DefaultGenesisState()
	genState.DenomMetadata = append(genState.DenomMetadata, Metadata)

	return cdc.MustMarshalJSON(genState)
}

// StakingModule defines a custom wrapper around the x/staking module's
// AppModuleBasic implementation to provide custom default genesis state.
type StakingModule struct {
	staking.AppModuleBasic
}

// DefaultGenesis returns custom x/staking module genesis state.
func (StakingModule) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	params := stakingtypes.DefaultParams()
	params.BondDenom = BondDenom

	return cdc.MustMarshalJSON(&stakingtypes.GenesisState{
		Params: params,
	})
}

// MintModule defines a custom wrapper around the x/mint module's
// AppModuleBasic implementation to provide custom default genesis state.
type MintModule struct {
	mint.AppModuleBasic
}

// DefaultGenesis returns custom x/mint module genesis state.
func (MintModule) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	genState := minttypes.DefaultGenesisState()
	genState.Params.MintDenom = BondDenom

	return cdc.MustMarshalJSON(genState)
}

// GovModule defines a custom wrapper around the x/gov module's
// AppModuleBasic implementation to provide custom default genesis state.
type GovModule struct {
	gov.AppModuleBasic
}

// DefaultGenesis returns custom x/gov module genesis state.
func (GovModule) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	minDeposit := sdk.NewCoins(sdk.NewCoin(BondDenom, govtypes.DefaultMinDepositTokens))
	genState := govtypes.DefaultGenesisState()
	genState.DepositParams.MinDeposit = minDeposit

	return cdc.MustMarshalJSON(genState)
}

// CrisisModule defines a custom wrapper around the x/crisis module's
// AppModuleBasic implementation to provide custom default genesis state.
type CrisisModule struct {
	crisis.AppModuleBasic
}

// DefaultGenesis returns custom x/crisis module genesis state.
func (CrisisModule) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	return cdc.MustMarshalJSON(&crisistypes.GenesisState{
		ConstantFee: sdk.NewCoin(BondDenom, sdk.NewInt(1000)),
	})
}
