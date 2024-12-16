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
// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
