package cardchain

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/DecentralCardGame/cardchain/testutil/sample"
	cardchainsimulation "github.com/DecentralCardGame/cardchain/x/cardchain/simulation"
	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
)

// avoid unused import issue
var (
	_ = cardchainsimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgUserCreate = "op_weight_msg_user_create"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUserCreate int = 100

	opWeightMsgCardSchemeBuy = "op_weight_msg_card_scheme_buy"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCardSchemeBuy int = 100

	opWeightMsgCardSaveContent = "op_weight_msg_card_save_content"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCardSaveContent int = 100

	opWeightMsgCardVote = "op_weight_msg_card_vote"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCardVote int = 100

	opWeightMsgCardTransfer = "op_weight_msg_card_transfer"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCardTransfer int = 100

	opWeightMsgCardDonate = "op_weight_msg_card_donate"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCardDonate int = 100

	opWeightMsgCardArtworkAdd = "op_weight_msg_card_artwork_add"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCardArtworkAdd int = 100

	opWeightMsgCardArtistChange = "op_weight_msg_card_artist_change"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCardArtistChange int = 100

	opWeightMsgCouncilRegister = "op_weight_msg_council_register"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCouncilRegister int = 100

	opWeightMsgCouncilDeregister = "op_weight_msg_council_deregister"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCouncilDeregister int = 100

	opWeightMsgMatchReport = "op_weight_msg_match_report"
	// TODO: Determine the simulation weight value
	defaultWeightMsgMatchReport int = 100

	opWeightMsgCouncilCreate = "op_weight_msg_council_create"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCouncilCreate int = 100

	opWeightMsgMatchReporterAppoint = "op_weight_msg_match_reporter_appoint"
	// TODO: Determine the simulation weight value
	defaultWeightMsgMatchReporterAppoint int = 100

	opWeightMsgSetCreate = "op_weight_msg_set_create"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSetCreate int = 100

	opWeightMsgSetCardAdd = "op_weight_msg_set_card_add"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSetCardAdd int = 100

	opWeightMsgSetCardRemove = "op_weight_msg_set_card_remove"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSetCardRemove int = 100

	opWeightMsgSetContributorAdd = "op_weight_msg_set_contributor_add"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSetContributorAdd int = 100

	opWeightMsgSetContributorRemove = "op_weight_msg_set_contributor_remove"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSetContributorRemove int = 100

	opWeightMsgSetFinalize = "op_weight_msg_set_finalize"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSetFinalize int = 100

	opWeightMsgSetArtworkAdd = "op_weight_msg_set_artwork_add"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSetArtworkAdd int = 100

	opWeightMsgSetStoryAdd = "op_weight_msg_set_story_add"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSetStoryAdd int = 100

	opWeightMsgBoosterPackBuy = "op_weight_msg_booster_pack_buy"
	// TODO: Determine the simulation weight value
	defaultWeightMsgBoosterPackBuy int = 100

	opWeightMsgSellOfferCreate = "op_weight_msg_sell_offer_create"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSellOfferCreate int = 100

	opWeightMsgSellOfferBuy = "op_weight_msg_sell_offer_buy"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSellOfferBuy int = 100

	opWeightMsgSellOfferRemove = "op_weight_msg_sell_offer_remove"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSellOfferRemove int = 100

	opWeightMsgCardRaritySet = "op_weight_msg_card_rarity_set"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCardRaritySet int = 100

	opWeightMsgCouncilResponseCommit = "op_weight_msg_council_response_commit"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCouncilResponseCommit int = 100

	opWeightMsgCouncilResponseReveal = "op_weight_msg_council_response_reveal"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCouncilResponseReveal int = 100

	opWeightMsgCouncilRestart = "op_weight_msg_council_restart"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCouncilRestart int = 100

	opWeightMsgMatchConfirm = "op_weight_msg_match_confirm"
	// TODO: Determine the simulation weight value
	defaultWeightMsgMatchConfirm int = 100

	opWeightMsgProfileCardSet = "op_weight_msg_profile_card_set"
	// TODO: Determine the simulation weight value
	defaultWeightMsgProfileCardSet int = 100

	opWeightMsgProfileWebsiteSet = "op_weight_msg_profile_website_set"
	// TODO: Determine the simulation weight value
	defaultWeightMsgProfileWebsiteSet int = 100

	opWeightMsgProfileBioSet = "op_weight_msg_profile_bio_set"
	// TODO: Determine the simulation weight value
	defaultWeightMsgProfileBioSet int = 100

	opWeightMsgBoosterPackOpen = "op_weight_msg_booster_pack_open"
	// TODO: Determine the simulation weight value
	defaultWeightMsgBoosterPackOpen int = 100

	opWeightMsgBoosterPackTransfer = "op_weight_msg_booster_pack_transfer"
	// TODO: Determine the simulation weight value
	defaultWeightMsgBoosterPackTransfer int = 100

	opWeightMsgSetStoryWriterSet = "op_weight_msg_set_story_writer_set"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSetStoryWriterSet int = 100

	opWeightMsgSetArtistSet = "op_weight_msg_set_artist_set"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSetArtistSet int = 100

	opWeightMsgCardVoteMulti = "op_weight_msg_card_vote_multi"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCardVoteMulti int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	cardchainGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&cardchainGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgUserCreate int
	simState.AppParams.GetOrGenerate(opWeightMsgUserCreate, &weightMsgUserCreate, nil,
		func(_ *rand.Rand) {
			weightMsgUserCreate = defaultWeightMsgUserCreate
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUserCreate,
		cardchainsimulation.SimulateMsgUserCreate(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCardSchemeBuy int
	simState.AppParams.GetOrGenerate(opWeightMsgCardSchemeBuy, &weightMsgCardSchemeBuy, nil,
		func(_ *rand.Rand) {
			weightMsgCardSchemeBuy = defaultWeightMsgCardSchemeBuy
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCardSchemeBuy,
		cardchainsimulation.SimulateMsgCardSchemeBuy(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCardSaveContent int
	simState.AppParams.GetOrGenerate(opWeightMsgCardSaveContent, &weightMsgCardSaveContent, nil,
		func(_ *rand.Rand) {
			weightMsgCardSaveContent = defaultWeightMsgCardSaveContent
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCardSaveContent,
		cardchainsimulation.SimulateMsgCardSaveContent(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCardVote int
	simState.AppParams.GetOrGenerate(opWeightMsgCardVote, &weightMsgCardVote, nil,
		func(_ *rand.Rand) {
			weightMsgCardVote = defaultWeightMsgCardVote
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCardVote,
		cardchainsimulation.SimulateMsgCardVote(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCardTransfer int
	simState.AppParams.GetOrGenerate(opWeightMsgCardTransfer, &weightMsgCardTransfer, nil,
		func(_ *rand.Rand) {
			weightMsgCardTransfer = defaultWeightMsgCardTransfer
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCardTransfer,
		cardchainsimulation.SimulateMsgCardTransfer(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCardDonate int
	simState.AppParams.GetOrGenerate(opWeightMsgCardDonate, &weightMsgCardDonate, nil,
		func(_ *rand.Rand) {
			weightMsgCardDonate = defaultWeightMsgCardDonate
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCardDonate,
		cardchainsimulation.SimulateMsgCardDonate(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCardArtworkAdd int
	simState.AppParams.GetOrGenerate(opWeightMsgCardArtworkAdd, &weightMsgCardArtworkAdd, nil,
		func(_ *rand.Rand) {
			weightMsgCardArtworkAdd = defaultWeightMsgCardArtworkAdd
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCardArtworkAdd,
		cardchainsimulation.SimulateMsgCardArtworkAdd(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCardArtistChange int
	simState.AppParams.GetOrGenerate(opWeightMsgCardArtistChange, &weightMsgCardArtistChange, nil,
		func(_ *rand.Rand) {
			weightMsgCardArtistChange = defaultWeightMsgCardArtistChange
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCardArtistChange,
		cardchainsimulation.SimulateMsgCardArtistChange(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCouncilRegister int
	simState.AppParams.GetOrGenerate(opWeightMsgCouncilRegister, &weightMsgCouncilRegister, nil,
		func(_ *rand.Rand) {
			weightMsgCouncilRegister = defaultWeightMsgCouncilRegister
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCouncilRegister,
		cardchainsimulation.SimulateMsgCouncilRegister(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCouncilDeregister int
	simState.AppParams.GetOrGenerate(opWeightMsgCouncilDeregister, &weightMsgCouncilDeregister, nil,
		func(_ *rand.Rand) {
			weightMsgCouncilDeregister = defaultWeightMsgCouncilDeregister
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCouncilDeregister,
		cardchainsimulation.SimulateMsgCouncilDeregister(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgMatchReport int
	simState.AppParams.GetOrGenerate(opWeightMsgMatchReport, &weightMsgMatchReport, nil,
		func(_ *rand.Rand) {
			weightMsgMatchReport = defaultWeightMsgMatchReport
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgMatchReport,
		cardchainsimulation.SimulateMsgMatchReport(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCouncilCreate int
	simState.AppParams.GetOrGenerate(opWeightMsgCouncilCreate, &weightMsgCouncilCreate, nil,
		func(_ *rand.Rand) {
			weightMsgCouncilCreate = defaultWeightMsgCouncilCreate
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCouncilCreate,
		cardchainsimulation.SimulateMsgCouncilCreate(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgMatchReporterAppoint int
	simState.AppParams.GetOrGenerate(opWeightMsgMatchReporterAppoint, &weightMsgMatchReporterAppoint, nil,
		func(_ *rand.Rand) {
			weightMsgMatchReporterAppoint = defaultWeightMsgMatchReporterAppoint
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgMatchReporterAppoint,
		cardchainsimulation.SimulateMsgMatchReporterAppoint(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSetCreate int
	simState.AppParams.GetOrGenerate(opWeightMsgSetCreate, &weightMsgSetCreate, nil,
		func(_ *rand.Rand) {
			weightMsgSetCreate = defaultWeightMsgSetCreate
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSetCreate,
		cardchainsimulation.SimulateMsgSetCreate(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSetCardAdd int
	simState.AppParams.GetOrGenerate(opWeightMsgSetCardAdd, &weightMsgSetCardAdd, nil,
		func(_ *rand.Rand) {
			weightMsgSetCardAdd = defaultWeightMsgSetCardAdd
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSetCardAdd,
		cardchainsimulation.SimulateMsgSetCardAdd(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSetCardRemove int
	simState.AppParams.GetOrGenerate(opWeightMsgSetCardRemove, &weightMsgSetCardRemove, nil,
		func(_ *rand.Rand) {
			weightMsgSetCardRemove = defaultWeightMsgSetCardRemove
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSetCardRemove,
		cardchainsimulation.SimulateMsgSetCardRemove(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSetContributorAdd int
	simState.AppParams.GetOrGenerate(opWeightMsgSetContributorAdd, &weightMsgSetContributorAdd, nil,
		func(_ *rand.Rand) {
			weightMsgSetContributorAdd = defaultWeightMsgSetContributorAdd
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSetContributorAdd,
		cardchainsimulation.SimulateMsgSetContributorAdd(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSetContributorRemove int
	simState.AppParams.GetOrGenerate(opWeightMsgSetContributorRemove, &weightMsgSetContributorRemove, nil,
		func(_ *rand.Rand) {
			weightMsgSetContributorRemove = defaultWeightMsgSetContributorRemove
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSetContributorRemove,
		cardchainsimulation.SimulateMsgSetContributorRemove(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSetFinalize int
	simState.AppParams.GetOrGenerate(opWeightMsgSetFinalize, &weightMsgSetFinalize, nil,
		func(_ *rand.Rand) {
			weightMsgSetFinalize = defaultWeightMsgSetFinalize
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSetFinalize,
		cardchainsimulation.SimulateMsgSetFinalize(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSetArtworkAdd int
	simState.AppParams.GetOrGenerate(opWeightMsgSetArtworkAdd, &weightMsgSetArtworkAdd, nil,
		func(_ *rand.Rand) {
			weightMsgSetArtworkAdd = defaultWeightMsgSetArtworkAdd
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSetArtworkAdd,
		cardchainsimulation.SimulateMsgSetArtworkAdd(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSetStoryAdd int
	simState.AppParams.GetOrGenerate(opWeightMsgSetStoryAdd, &weightMsgSetStoryAdd, nil,
		func(_ *rand.Rand) {
			weightMsgSetStoryAdd = defaultWeightMsgSetStoryAdd
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSetStoryAdd,
		cardchainsimulation.SimulateMsgSetStoryAdd(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgBoosterPackBuy int
	simState.AppParams.GetOrGenerate(opWeightMsgBoosterPackBuy, &weightMsgBoosterPackBuy, nil,
		func(_ *rand.Rand) {
			weightMsgBoosterPackBuy = defaultWeightMsgBoosterPackBuy
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgBoosterPackBuy,
		cardchainsimulation.SimulateMsgBoosterPackBuy(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSellOfferCreate int
	simState.AppParams.GetOrGenerate(opWeightMsgSellOfferCreate, &weightMsgSellOfferCreate, nil,
		func(_ *rand.Rand) {
			weightMsgSellOfferCreate = defaultWeightMsgSellOfferCreate
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSellOfferCreate,
		cardchainsimulation.SimulateMsgSellOfferCreate(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSellOfferBuy int
	simState.AppParams.GetOrGenerate(opWeightMsgSellOfferBuy, &weightMsgSellOfferBuy, nil,
		func(_ *rand.Rand) {
			weightMsgSellOfferBuy = defaultWeightMsgSellOfferBuy
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSellOfferBuy,
		cardchainsimulation.SimulateMsgSellOfferBuy(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSellOfferRemove int
	simState.AppParams.GetOrGenerate(opWeightMsgSellOfferRemove, &weightMsgSellOfferRemove, nil,
		func(_ *rand.Rand) {
			weightMsgSellOfferRemove = defaultWeightMsgSellOfferRemove
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSellOfferRemove,
		cardchainsimulation.SimulateMsgSellOfferRemove(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCardRaritySet int
	simState.AppParams.GetOrGenerate(opWeightMsgCardRaritySet, &weightMsgCardRaritySet, nil,
		func(_ *rand.Rand) {
			weightMsgCardRaritySet = defaultWeightMsgCardRaritySet
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCardRaritySet,
		cardchainsimulation.SimulateMsgCardRaritySet(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCouncilResponseCommit int
	simState.AppParams.GetOrGenerate(opWeightMsgCouncilResponseCommit, &weightMsgCouncilResponseCommit, nil,
		func(_ *rand.Rand) {
			weightMsgCouncilResponseCommit = defaultWeightMsgCouncilResponseCommit
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCouncilResponseCommit,
		cardchainsimulation.SimulateMsgCouncilResponseCommit(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCouncilResponseReveal int
	simState.AppParams.GetOrGenerate(opWeightMsgCouncilResponseReveal, &weightMsgCouncilResponseReveal, nil,
		func(_ *rand.Rand) {
			weightMsgCouncilResponseReveal = defaultWeightMsgCouncilResponseReveal
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCouncilResponseReveal,
		cardchainsimulation.SimulateMsgCouncilResponseReveal(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCouncilRestart int
	simState.AppParams.GetOrGenerate(opWeightMsgCouncilRestart, &weightMsgCouncilRestart, nil,
		func(_ *rand.Rand) {
			weightMsgCouncilRestart = defaultWeightMsgCouncilRestart
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCouncilRestart,
		cardchainsimulation.SimulateMsgCouncilRestart(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgMatchConfirm int
	simState.AppParams.GetOrGenerate(opWeightMsgMatchConfirm, &weightMsgMatchConfirm, nil,
		func(_ *rand.Rand) {
			weightMsgMatchConfirm = defaultWeightMsgMatchConfirm
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgMatchConfirm,
		cardchainsimulation.SimulateMsgMatchConfirm(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgProfileCardSet int
	simState.AppParams.GetOrGenerate(opWeightMsgProfileCardSet, &weightMsgProfileCardSet, nil,
		func(_ *rand.Rand) {
			weightMsgProfileCardSet = defaultWeightMsgProfileCardSet
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgProfileCardSet,
		cardchainsimulation.SimulateMsgProfileCardSet(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgProfileWebsiteSet int
	simState.AppParams.GetOrGenerate(opWeightMsgProfileWebsiteSet, &weightMsgProfileWebsiteSet, nil,
		func(_ *rand.Rand) {
			weightMsgProfileWebsiteSet = defaultWeightMsgProfileWebsiteSet
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgProfileWebsiteSet,
		cardchainsimulation.SimulateMsgProfileWebsiteSet(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgProfileBioSet int
	simState.AppParams.GetOrGenerate(opWeightMsgProfileBioSet, &weightMsgProfileBioSet, nil,
		func(_ *rand.Rand) {
			weightMsgProfileBioSet = defaultWeightMsgProfileBioSet
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgProfileBioSet,
		cardchainsimulation.SimulateMsgProfileBioSet(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgBoosterPackOpen int
	simState.AppParams.GetOrGenerate(opWeightMsgBoosterPackOpen, &weightMsgBoosterPackOpen, nil,
		func(_ *rand.Rand) {
			weightMsgBoosterPackOpen = defaultWeightMsgBoosterPackOpen
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgBoosterPackOpen,
		cardchainsimulation.SimulateMsgBoosterPackOpen(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgBoosterPackTransfer int
	simState.AppParams.GetOrGenerate(opWeightMsgBoosterPackTransfer, &weightMsgBoosterPackTransfer, nil,
		func(_ *rand.Rand) {
			weightMsgBoosterPackTransfer = defaultWeightMsgBoosterPackTransfer
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgBoosterPackTransfer,
		cardchainsimulation.SimulateMsgBoosterPackTransfer(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSetStoryWriterSet int
	simState.AppParams.GetOrGenerate(opWeightMsgSetStoryWriterSet, &weightMsgSetStoryWriterSet, nil,
		func(_ *rand.Rand) {
			weightMsgSetStoryWriterSet = defaultWeightMsgSetStoryWriterSet
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSetStoryWriterSet,
		cardchainsimulation.SimulateMsgSetStoryWriterSet(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSetArtistSet int
	simState.AppParams.GetOrGenerate(opWeightMsgSetArtistSet, &weightMsgSetArtistSet, nil,
		func(_ *rand.Rand) {
			weightMsgSetArtistSet = defaultWeightMsgSetArtistSet
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSetArtistSet,
		cardchainsimulation.SimulateMsgSetArtistSet(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCardVoteMulti int
	simState.AppParams.GetOrGenerate(opWeightMsgCardVoteMulti, &weightMsgCardVoteMulti, nil,
		func(_ *rand.Rand) {
			weightMsgCardVoteMulti = defaultWeightMsgCardVoteMulti
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCardVoteMulti,
		cardchainsimulation.SimulateMsgCardVoteMulti(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgUserCreate,
			defaultWeightMsgUserCreate,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				cardchainsimulation.SimulateMsgUserCreate(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCardSchemeBuy,
			defaultWeightMsgCardSchemeBuy,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				cardchainsimulation.SimulateMsgCardSchemeBuy(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCardSaveContent,
			defaultWeightMsgCardSaveContent,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				cardchainsimulation.SimulateMsgCardSaveContent(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCardVote,
			defaultWeightMsgCardVote,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				cardchainsimulation.SimulateMsgCardVote(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCardTransfer,
			defaultWeightMsgCardTransfer,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				cardchainsimulation.SimulateMsgCardTransfer(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCardDonate,
			defaultWeightMsgCardDonate,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				cardchainsimulation.SimulateMsgCardDonate(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCardArtworkAdd,
			defaultWeightMsgCardArtworkAdd,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				cardchainsimulation.SimulateMsgCardArtworkAdd(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCardArtistChange,
			defaultWeightMsgCardArtistChange,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				cardchainsimulation.SimulateMsgCardArtistChange(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCouncilRegister,
			defaultWeightMsgCouncilRegister,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				cardchainsimulation.SimulateMsgCouncilRegister(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCouncilDeregister,
			defaultWeightMsgCouncilDeregister,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				cardchainsimulation.SimulateMsgCouncilDeregister(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgMatchReport,
			defaultWeightMsgMatchReport,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				cardchainsimulation.SimulateMsgMatchReport(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCouncilCreate,
			defaultWeightMsgCouncilCreate,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				cardchainsimulation.SimulateMsgCouncilCreate(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgMatchReporterAppoint,
			defaultWeightMsgMatchReporterAppoint,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				cardchainsimulation.SimulateMsgMatchReporterAppoint(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgSetCreate,
			defaultWeightMsgSetCreate,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				cardchainsimulation.SimulateMsgSetCreate(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgSetCardAdd,
			defaultWeightMsgSetCardAdd,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				cardchainsimulation.SimulateMsgSetCardAdd(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgSetCardRemove,
			defaultWeightMsgSetCardRemove,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				cardchainsimulation.SimulateMsgSetCardRemove(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgSetContributorAdd,
			defaultWeightMsgSetContributorAdd,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				cardchainsimulation.SimulateMsgSetContributorAdd(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgSetContributorRemove,
			defaultWeightMsgSetContributorRemove,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				cardchainsimulation.SimulateMsgSetContributorRemove(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgSetFinalize,
			defaultWeightMsgSetFinalize,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				cardchainsimulation.SimulateMsgSetFinalize(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgSetArtworkAdd,
			defaultWeightMsgSetArtworkAdd,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				cardchainsimulation.SimulateMsgSetArtworkAdd(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgSetStoryAdd,
			defaultWeightMsgSetStoryAdd,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				cardchainsimulation.SimulateMsgSetStoryAdd(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgBoosterPackBuy,
			defaultWeightMsgBoosterPackBuy,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				cardchainsimulation.SimulateMsgBoosterPackBuy(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgSellOfferCreate,
			defaultWeightMsgSellOfferCreate,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				cardchainsimulation.SimulateMsgSellOfferCreate(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgSellOfferBuy,
			defaultWeightMsgSellOfferBuy,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				cardchainsimulation.SimulateMsgSellOfferBuy(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgSellOfferRemove,
			defaultWeightMsgSellOfferRemove,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				cardchainsimulation.SimulateMsgSellOfferRemove(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCardRaritySet,
			defaultWeightMsgCardRaritySet,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				cardchainsimulation.SimulateMsgCardRaritySet(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCouncilResponseCommit,
			defaultWeightMsgCouncilResponseCommit,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				cardchainsimulation.SimulateMsgCouncilResponseCommit(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCouncilResponseReveal,
			defaultWeightMsgCouncilResponseReveal,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				cardchainsimulation.SimulateMsgCouncilResponseReveal(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCouncilRestart,
			defaultWeightMsgCouncilRestart,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				cardchainsimulation.SimulateMsgCouncilRestart(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgMatchConfirm,
			defaultWeightMsgMatchConfirm,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				cardchainsimulation.SimulateMsgMatchConfirm(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgProfileCardSet,
			defaultWeightMsgProfileCardSet,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				cardchainsimulation.SimulateMsgProfileCardSet(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgProfileWebsiteSet,
			defaultWeightMsgProfileWebsiteSet,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				cardchainsimulation.SimulateMsgProfileWebsiteSet(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgProfileBioSet,
			defaultWeightMsgProfileBioSet,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				cardchainsimulation.SimulateMsgProfileBioSet(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgBoosterPackOpen,
			defaultWeightMsgBoosterPackOpen,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				cardchainsimulation.SimulateMsgBoosterPackOpen(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgBoosterPackTransfer,
			defaultWeightMsgBoosterPackTransfer,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				cardchainsimulation.SimulateMsgBoosterPackTransfer(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgSetStoryWriterSet,
			defaultWeightMsgSetStoryWriterSet,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				cardchainsimulation.SimulateMsgSetStoryWriterSet(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgSetArtistSet,
			defaultWeightMsgSetArtistSet,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				cardchainsimulation.SimulateMsgSetArtistSet(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCardVoteMulti,
			defaultWeightMsgCardVoteMulti,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				cardchainsimulation.SimulateMsgCardVoteMulti(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
