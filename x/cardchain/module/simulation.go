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
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
