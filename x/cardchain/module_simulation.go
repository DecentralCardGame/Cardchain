package cardchain

import (
	"math/rand"

	"github.com/DecentralCardGame/Cardchain/testutil/sample"
	cardchainsimulation "github.com/DecentralCardGame/Cardchain/x/cardchain/simulation"
	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = cardchainsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateuser = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateuser int = 100

	opWeightMsgBuyCardScheme = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgBuyCardScheme int = 100

	opWeightMsgVoteCard = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgVoteCard int = 100

	opWeightMsgSaveCardContent = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSaveCardContent int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	cardchainGenesis := types.GenesisState{
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&cardchainGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateuser int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateuser, &weightMsgCreateuser, nil,
		func(_ *rand.Rand) {
			weightMsgCreateuser = defaultWeightMsgCreateuser
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateuser,
		cardchainsimulation.SimulateMsgCreateuser(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgBuyCardScheme int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgBuyCardScheme, &weightMsgBuyCardScheme, nil,
		func(_ *rand.Rand) {
			weightMsgBuyCardScheme = defaultWeightMsgBuyCardScheme
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgBuyCardScheme,
		cardchainsimulation.SimulateMsgBuyCardScheme(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgVoteCard int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgVoteCard, &weightMsgVoteCard, nil,
		func(_ *rand.Rand) {
			weightMsgVoteCard = defaultWeightMsgVoteCard
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgVoteCard,
		cardchainsimulation.SimulateMsgVoteCard(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSaveCardContent int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSaveCardContent, &weightMsgSaveCardContent, nil,
		func(_ *rand.Rand) {
			weightMsgSaveCardContent = defaultWeightMsgSaveCardContent
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSaveCardContent,
		cardchainsimulation.SimulateMsgSaveCardContent(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}