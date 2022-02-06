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

	opWeightMsgTransferCard = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgTransferCard int = 100

	opWeightMsgDonateToCard = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDonateToCard int = 100

	opWeightMsgAddArtwork = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAddArtwork int = 100

	opWeightMsgSubmitCopyrightProposal = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSubmitCopyrightProposal int = 100

	opWeightMsgChangeArtist = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgChangeArtist int = 100

	opWeightMsgRegisterForCouncil = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRegisterForCouncil int = 100

	opWeightMsgReportMatch = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgReportMatch int = 100

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

	var weightMsgTransferCard int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgTransferCard, &weightMsgTransferCard, nil,
		func(_ *rand.Rand) {
			weightMsgTransferCard = defaultWeightMsgTransferCard
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgTransferCard,
		cardchainsimulation.SimulateMsgTransferCard(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDonateToCard int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDonateToCard, &weightMsgDonateToCard, nil,
		func(_ *rand.Rand) {
			weightMsgDonateToCard = defaultWeightMsgDonateToCard
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDonateToCard,
		cardchainsimulation.SimulateMsgDonateToCard(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgAddArtwork int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgAddArtwork, &weightMsgAddArtwork, nil,
		func(_ *rand.Rand) {
			weightMsgAddArtwork = defaultWeightMsgAddArtwork
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAddArtwork,
		cardchainsimulation.SimulateMsgAddArtwork(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSubmitCopyrightProposal int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSubmitCopyrightProposal, &weightMsgSubmitCopyrightProposal, nil,
		func(_ *rand.Rand) {
			weightMsgSubmitCopyrightProposal = defaultWeightMsgSubmitCopyrightProposal
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSubmitCopyrightProposal,
		cardchainsimulation.SimulateMsgSubmitCopyrightProposal(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgChangeArtist int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgChangeArtist, &weightMsgChangeArtist, nil,
		func(_ *rand.Rand) {
			weightMsgChangeArtist = defaultWeightMsgChangeArtist
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgChangeArtist,
		cardchainsimulation.SimulateMsgChangeArtist(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgRegisterForCouncil int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRegisterForCouncil, &weightMsgRegisterForCouncil, nil,
		func(_ *rand.Rand) {
			weightMsgRegisterForCouncil = defaultWeightMsgRegisterForCouncil
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRegisterForCouncil,
		cardchainsimulation.SimulateMsgRegisterForCouncil(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgReportMatch int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgReportMatch, &weightMsgReportMatch, nil,
		func(_ *rand.Rand) {
			weightMsgReportMatch = defaultWeightMsgReportMatch
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgReportMatch,
		cardchainsimulation.SimulateMsgReportMatch(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
