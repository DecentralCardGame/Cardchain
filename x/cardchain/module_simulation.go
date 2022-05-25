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

	opWeightMsgSubmitMatchReporterProposal = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSubmitMatchReporterProposal int = 100

	opWeightMsgApointMatchReporter = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgApointMatchReporter int = 100

	opWeightMsgCreateCollection = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateCollection int = 100

	opWeightMsgAddCardToCollection = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAddCardToCollection int = 100

	opWeightMsgFinalizeCollection = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgFinalizeCollection int = 100

	opWeightMsgBuyCollection = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgBuyCollection int = 100

	opWeightMsgRemoveCardFromCollection = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRemoveCardFromCollection int = 100

	opWeightMsgRemoveContributorFromCollection = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRemoveContributorFromCollection int = 100

	opWeightMsgAddContributorToCollection = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAddContributorToCollection int = 100

	opWeightMsgSubmitCollectionProposal = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSubmitCollectionProposal int = 100

	opWeightMsgCreateSellOffer = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateSellOffer int = 100

	opWeightMsgBuyCard = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgBuyCard int = 100

	opWeightMsgRemoveSellOffer = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRemoveSellOffer int = 100

	opWeightMsgAddArtworkToCollection = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAddArtworkToCollection int = 100

	opWeightMsgAddStoryToCollection = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAddStoryToCollection int = 100

	opWeightMsgSetCardRarity = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSetCardRarity int = 100

	opWeightMsgCreateCouncil = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateCouncil int = 100

	opWeightMsgCommitCouncilResponse = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCommitCouncilResponse int = 100

	opWeightMsgRevealCouncilResponse = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRevealCouncilResponse int = 100

	opWeightMsgRestartCouncil = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRestartCouncil int = 100

	opWeightMsgRewokeCouncilRegistration = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRewokeCouncilRegistration int = 100

	opWeightMsgConfirmMatch = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgConfirmMatch int = 100

	opWeightMsgSetProfileCard = "op_weight_msg_set_profile_card"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSetProfileCard int = 100

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

	var weightMsgSubmitMatchReporterProposal int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSubmitMatchReporterProposal, &weightMsgSubmitMatchReporterProposal, nil,
		func(_ *rand.Rand) {
			weightMsgSubmitMatchReporterProposal = defaultWeightMsgSubmitMatchReporterProposal
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSubmitMatchReporterProposal,
		cardchainsimulation.SimulateMsgSubmitMatchReporterProposal(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgApointMatchReporter int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgApointMatchReporter, &weightMsgApointMatchReporter, nil,
		func(_ *rand.Rand) {
			weightMsgApointMatchReporter = defaultWeightMsgApointMatchReporter
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgApointMatchReporter,
		cardchainsimulation.SimulateMsgApointMatchReporter(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateCollection int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateCollection, &weightMsgCreateCollection, nil,
		func(_ *rand.Rand) {
			weightMsgCreateCollection = defaultWeightMsgCreateCollection
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateCollection,
		cardchainsimulation.SimulateMsgCreateCollection(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgAddCardToCollection int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgAddCardToCollection, &weightMsgAddCardToCollection, nil,
		func(_ *rand.Rand) {
			weightMsgAddCardToCollection = defaultWeightMsgAddCardToCollection
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAddCardToCollection,
		cardchainsimulation.SimulateMsgAddCardToCollection(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgFinalizeCollection int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgFinalizeCollection, &weightMsgFinalizeCollection, nil,
		func(_ *rand.Rand) {
			weightMsgFinalizeCollection = defaultWeightMsgFinalizeCollection
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgFinalizeCollection,
		cardchainsimulation.SimulateMsgFinalizeCollection(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgBuyCollection int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgBuyCollection, &weightMsgBuyCollection, nil,
		func(_ *rand.Rand) {
			weightMsgBuyCollection = defaultWeightMsgBuyCollection
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgBuyCollection,
		cardchainsimulation.SimulateMsgBuyCollection(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgRemoveCardFromCollection int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRemoveCardFromCollection, &weightMsgRemoveCardFromCollection, nil,
		func(_ *rand.Rand) {
			weightMsgRemoveCardFromCollection = defaultWeightMsgRemoveCardFromCollection
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRemoveCardFromCollection,
		cardchainsimulation.SimulateMsgRemoveCardFromCollection(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgRemoveContributorFromCollection int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRemoveContributorFromCollection, &weightMsgRemoveContributorFromCollection, nil,
		func(_ *rand.Rand) {
			weightMsgRemoveContributorFromCollection = defaultWeightMsgRemoveContributorFromCollection
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRemoveContributorFromCollection,
		cardchainsimulation.SimulateMsgRemoveContributorFromCollection(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgAddContributorToCollection int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgAddContributorToCollection, &weightMsgAddContributorToCollection, nil,
		func(_ *rand.Rand) {
			weightMsgAddContributorToCollection = defaultWeightMsgAddContributorToCollection
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAddContributorToCollection,
		cardchainsimulation.SimulateMsgAddContributorToCollection(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSubmitCollectionProposal int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSubmitCollectionProposal, &weightMsgSubmitCollectionProposal, nil,
		func(_ *rand.Rand) {
			weightMsgSubmitCollectionProposal = defaultWeightMsgSubmitCollectionProposal
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSubmitCollectionProposal,
		cardchainsimulation.SimulateMsgSubmitCollectionProposal(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateSellOffer int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateSellOffer, &weightMsgCreateSellOffer, nil,
		func(_ *rand.Rand) {
			weightMsgCreateSellOffer = defaultWeightMsgCreateSellOffer
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateSellOffer,
		cardchainsimulation.SimulateMsgCreateSellOffer(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgBuyCard int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgBuyCard, &weightMsgBuyCard, nil,
		func(_ *rand.Rand) {
			weightMsgBuyCard = defaultWeightMsgBuyCard
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgBuyCard,
		cardchainsimulation.SimulateMsgBuyCard(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgRemoveSellOffer int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRemoveSellOffer, &weightMsgRemoveSellOffer, nil,
		func(_ *rand.Rand) {
			weightMsgRemoveSellOffer = defaultWeightMsgRemoveSellOffer
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRemoveSellOffer,
		cardchainsimulation.SimulateMsgRemoveSellOffer(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgAddArtworkToCollection int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgAddArtworkToCollection, &weightMsgAddArtworkToCollection, nil,
		func(_ *rand.Rand) {
			weightMsgAddArtworkToCollection = defaultWeightMsgAddArtworkToCollection
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAddArtworkToCollection,
		cardchainsimulation.SimulateMsgAddArtworkToCollection(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgAddStoryToCollection int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgAddStoryToCollection, &weightMsgAddStoryToCollection, nil,
		func(_ *rand.Rand) {
			weightMsgAddStoryToCollection = defaultWeightMsgAddStoryToCollection
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAddStoryToCollection,
		cardchainsimulation.SimulateMsgAddStoryToCollection(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSetCardRarity int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSetCardRarity, &weightMsgSetCardRarity, nil,
		func(_ *rand.Rand) {
			weightMsgSetCardRarity = defaultWeightMsgSetCardRarity
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSetCardRarity,
		cardchainsimulation.SimulateMsgSetCardRarity(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateCouncil int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateCouncil, &weightMsgCreateCouncil, nil,
		func(_ *rand.Rand) {
			weightMsgCreateCouncil = defaultWeightMsgCreateCouncil
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateCouncil,
		cardchainsimulation.SimulateMsgCreateCouncil(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCommitCouncilResponse int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCommitCouncilResponse, &weightMsgCommitCouncilResponse, nil,
		func(_ *rand.Rand) {
			weightMsgCommitCouncilResponse = defaultWeightMsgCommitCouncilResponse
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCommitCouncilResponse,
		cardchainsimulation.SimulateMsgCommitCouncilResponse(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgRevealCouncilResponse int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRevealCouncilResponse, &weightMsgRevealCouncilResponse, nil,
		func(_ *rand.Rand) {
			weightMsgRevealCouncilResponse = defaultWeightMsgRevealCouncilResponse
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRevealCouncilResponse,
		cardchainsimulation.SimulateMsgRevealCouncilResponse(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgRestartCouncil int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRestartCouncil, &weightMsgRestartCouncil, nil,
		func(_ *rand.Rand) {
			weightMsgRestartCouncil = defaultWeightMsgRestartCouncil
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRestartCouncil,
		cardchainsimulation.SimulateMsgRestartCouncil(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgRewokeCouncilRegistration int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRewokeCouncilRegistration, &weightMsgRewokeCouncilRegistration, nil,
		func(_ *rand.Rand) {
			weightMsgRewokeCouncilRegistration = defaultWeightMsgRewokeCouncilRegistration
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRewokeCouncilRegistration,
		cardchainsimulation.SimulateMsgRewokeCouncilRegistration(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgConfirmMatch int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgConfirmMatch, &weightMsgConfirmMatch, nil,
		func(_ *rand.Rand) {
			weightMsgConfirmMatch = defaultWeightMsgConfirmMatch
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgConfirmMatch,
		cardchainsimulation.SimulateMsgConfirmMatch(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSetProfileCard int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSetProfileCard, &weightMsgSetProfileCard, nil,
		func(_ *rand.Rand) {
			weightMsgSetProfileCard = defaultWeightMsgSetProfileCard
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSetProfileCard,
		cardchainsimulation.SimulateMsgSetProfileCard(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
