package cardchain

import (
	"math/rand"

	"github.com/DecentralCardGame/cardchain/testutil/sample"
	cardchainsimulation "github.com/DecentralCardGame/cardchain/x/cardchain/simulation"
	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
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

// GenerateGenesisState creates a randomized GenState of the module
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

	var weightMsgUserCreate int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUserCreate, &weightMsgUserCreate, nil,
		func(_ *rand.Rand) {
			weightMsgUserCreate = defaultWeightMsgUserCreate
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUserCreate,
		cardchainsimulation.SimulateMsgUserCreate(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCardSchemeBuy int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCardSchemeBuy, &weightMsgCardSchemeBuy, nil,
		func(_ *rand.Rand) {
			weightMsgCardSchemeBuy = defaultWeightMsgCardSchemeBuy
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCardSchemeBuy,
		cardchainsimulation.SimulateMsgCardSchemeBuy(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCardSaveContent int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCardSaveContent, &weightMsgCardSaveContent, nil,
		func(_ *rand.Rand) {
			weightMsgCardSaveContent = defaultWeightMsgCardSaveContent
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCardSaveContent,
		cardchainsimulation.SimulateMsgCardSaveContent(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCardVote int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCardVote, &weightMsgCardVote, nil,
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
