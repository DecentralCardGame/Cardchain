package cardchain

import (
	"math/rand"

	"github.com/DecentralCardGame/Cardchain/testutil/sample"
	cardchainsimulation "github.com/DecentralCardGame/Cardchain/x/cardchain/simulation"
	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = cardchainsimulation.FindAccount
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

	opWeightMsgCreateSet = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateSet int = 100

	opWeightMsgAddCardToSet = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAddCardToSet int = 100

	opWeightMsgFinalizeSet = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgFinalizeSet int = 100

	opWeightMsgBuyBoosterPack = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgBuyBoosterPack int = 100

	opWeightMsgRemoveCardFromSet = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRemoveCardFromSet int = 100

	opWeightMsgRemoveContributorFromSet = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRemoveContributorFromSet int = 100

	opWeightMsgAddContributorToSet = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAddContributorToSet int = 100

	opWeightMsgSubmitSetProposal = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSubmitSetProposal int = 100

	opWeightMsgCreateSellOffer = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateSellOffer int = 100

	opWeightMsgBuyCard = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgBuyCard int = 100

	opWeightMsgRemoveSellOffer = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRemoveSellOffer int = 100

	opWeightMsgAddArtworkToSet = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAddArtworkToSet int = 100

	opWeightMsgAddStoryToSet = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAddStoryToSet int = 100

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

	opWeightMsgOpenBoosterPack = "op_weight_msg_open_booster_pack"
	// TODO: Determine the simulation weight value
	defaultWeightMsgOpenBoosterPack int = 100

	opWeightMsgTransferBoosterPack = "op_weight_msg_transfer_booster_pack"
	// TODO: Determine the simulation weight value
	defaultWeightMsgTransferBoosterPack int = 100

	opWeightMsgSetSetStoryWriter = "op_weight_msg_set_set_story_writer"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSetSetStoryWriter int = 100

	opWeightMsgSetSetArtist = "op_weight_msg_set_set_artist"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSetSetArtist int = 100

	opWeightMsgSetUserWebsite = "op_weight_msg_set_user_website"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSetUserWebsite int = 100

	opWeightMsgSetUserBiography = "op_weight_msg_set_user_biography"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSetUserBiography int = 100

	opWeightMsgMultiVoteCard = "op_weight_msg_multi_vote_card"
	// TODO: Determine the simulation weight value
	defaultWeightMsgMultiVoteCard int = 100

	opWeightMsgOpenMatch = "op_weight_msg_msg_open_match"
	// TODO: Determine the simulation weight value
	defaultWeightMsgOpenMatch int = 100

	opWeightMsgSetSetName = "op_weight_msg_set_set_name"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSetSetName int = 100

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

	var weightMsgCreateSet int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateSet, &weightMsgCreateSet, nil,
		func(_ *rand.Rand) {
			weightMsgCreateSet = defaultWeightMsgCreateSet
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateSet,
		cardchainsimulation.SimulateMsgCreateSet(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgAddCardToSet int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgAddCardToSet, &weightMsgAddCardToSet, nil,
		func(_ *rand.Rand) {
			weightMsgAddCardToSet = defaultWeightMsgAddCardToSet
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAddCardToSet,
		cardchainsimulation.SimulateMsgAddCardToSet(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgFinalizeSet int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgFinalizeSet, &weightMsgFinalizeSet, nil,
		func(_ *rand.Rand) {
			weightMsgFinalizeSet = defaultWeightMsgFinalizeSet
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgFinalizeSet,
		cardchainsimulation.SimulateMsgFinalizeSet(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgBuyBoosterPack int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgBuyBoosterPack, &weightMsgBuyBoosterPack, nil,
		func(_ *rand.Rand) {
			weightMsgBuyBoosterPack = defaultWeightMsgBuyBoosterPack
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgBuyBoosterPack,
		cardchainsimulation.SimulateMsgBuyBoosterPack(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgRemoveCardFromSet int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRemoveCardFromSet, &weightMsgRemoveCardFromSet, nil,
		func(_ *rand.Rand) {
			weightMsgRemoveCardFromSet = defaultWeightMsgRemoveCardFromSet
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRemoveCardFromSet,
		cardchainsimulation.SimulateMsgRemoveCardFromSet(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgRemoveContributorFromSet int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRemoveContributorFromSet, &weightMsgRemoveContributorFromSet, nil,
		func(_ *rand.Rand) {
			weightMsgRemoveContributorFromSet = defaultWeightMsgRemoveContributorFromSet
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRemoveContributorFromSet,
		cardchainsimulation.SimulateMsgRemoveContributorFromSet(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgAddContributorToSet int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgAddContributorToSet, &weightMsgAddContributorToSet, nil,
		func(_ *rand.Rand) {
			weightMsgAddContributorToSet = defaultWeightMsgAddContributorToSet
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAddContributorToSet,
		cardchainsimulation.SimulateMsgAddContributorToSet(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSubmitSetProposal int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSubmitSetProposal, &weightMsgSubmitSetProposal, nil,
		func(_ *rand.Rand) {
			weightMsgSubmitSetProposal = defaultWeightMsgSubmitSetProposal
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSubmitSetProposal,
		cardchainsimulation.SimulateMsgSubmitSetProposal(am.accountKeeper, am.bankKeeper, am.keeper),
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

	var weightMsgAddArtworkToSet int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgAddArtworkToSet, &weightMsgAddArtworkToSet, nil,
		func(_ *rand.Rand) {
			weightMsgAddArtworkToSet = defaultWeightMsgAddArtworkToSet
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAddArtworkToSet,
		cardchainsimulation.SimulateMsgAddArtworkToSet(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgAddStoryToSet int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgAddStoryToSet, &weightMsgAddStoryToSet, nil,
		func(_ *rand.Rand) {
			weightMsgAddStoryToSet = defaultWeightMsgAddStoryToSet
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAddStoryToSet,
		cardchainsimulation.SimulateMsgAddStoryToSet(am.accountKeeper, am.bankKeeper, am.keeper),
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

	var weightMsgOpenBoosterPack int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgOpenBoosterPack, &weightMsgOpenBoosterPack, nil,
		func(_ *rand.Rand) {
			weightMsgOpenBoosterPack = defaultWeightMsgOpenBoosterPack
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgOpenBoosterPack,
		cardchainsimulation.SimulateMsgOpenBoosterPack(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgTransferBoosterPack int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgTransferBoosterPack, &weightMsgTransferBoosterPack, nil,
		func(_ *rand.Rand) {
			weightMsgTransferBoosterPack = defaultWeightMsgTransferBoosterPack
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgTransferBoosterPack,
		cardchainsimulation.SimulateMsgTransferBoosterPack(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSetSetStoryWriter int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSetSetStoryWriter, &weightMsgSetSetStoryWriter, nil,
		func(_ *rand.Rand) {
			weightMsgSetSetStoryWriter = defaultWeightMsgSetSetStoryWriter
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSetSetStoryWriter,
		cardchainsimulation.SimulateMsgSetSetStoryWriter(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSetSetArtist int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSetSetArtist, &weightMsgSetSetArtist, nil,
		func(_ *rand.Rand) {
			weightMsgSetSetArtist = defaultWeightMsgSetSetArtist
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSetSetArtist,
		cardchainsimulation.SimulateMsgSetSetArtist(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSetUserWebsite int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSetUserWebsite, &weightMsgSetUserWebsite, nil,
		func(_ *rand.Rand) {
			weightMsgSetUserWebsite = defaultWeightMsgSetUserWebsite
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSetUserWebsite,
		cardchainsimulation.SimulateMsgSetUserWebsite(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSetUserBiography int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSetUserBiography, &weightMsgSetUserBiography, nil,
		func(_ *rand.Rand) {
			weightMsgSetUserBiography = defaultWeightMsgSetUserBiography
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSetUserBiography,
		cardchainsimulation.SimulateMsgSetUserBiography(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgMultiVoteCard int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgMultiVoteCard, &weightMsgMultiVoteCard, nil,
		func(_ *rand.Rand) {
			weightMsgMultiVoteCard = defaultWeightMsgMultiVoteCard
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgMultiVoteCard,
		cardchainsimulation.SimulateMsgMultiVoteCard(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgOpenMatch int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgOpenMatch, &weightMsgOpenMatch, nil,
		func(_ *rand.Rand) {
			weightMsgOpenMatch = defaultWeightMsgOpenMatch
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgOpenMatch,
		cardchainsimulation.SimulateMsgOpenMatch(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSetSetName int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSetSetName, &weightMsgSetSetName, nil,
		func(_ *rand.Rand) {
			weightMsgSetSetName = defaultWeightMsgSetSetName
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSetSetName,
		cardchainsimulation.SimulateMsgSetSetName(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
