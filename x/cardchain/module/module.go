package cardchain

import (
	"context"
	"encoding/json"
	"fmt"

	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/core/store"
	"cosmossdk.io/depinject"
	"cosmossdk.io/log"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"

	// this line is used by starport scaffolding # 1

	modulev1 "github.com/DecentralCardGame/cardchain/api/cardchain/cardchain/module"
	"github.com/DecentralCardGame/cardchain/util"
	"github.com/DecentralCardGame/cardchain/x/cardchain/keeper"
	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
)

var (
	_ module.AppModuleBasic      = (*AppModule)(nil)
	_ module.AppModuleSimulation = (*AppModule)(nil)
	_ module.HasGenesis          = (*AppModule)(nil)
	_ module.HasInvariants       = (*AppModule)(nil)
	_ module.HasConsensusVersion = (*AppModule)(nil)

	_ appmodule.AppModule       = (*AppModule)(nil)
	_ appmodule.HasBeginBlocker = (*AppModule)(nil)
	_ appmodule.HasEndBlocker   = (*AppModule)(nil)
)

const epochBlockTime = 120000

// ----------------------------------------------------------------------------
// AppModuleBasic
// ----------------------------------------------------------------------------

// AppModuleBasic implements the AppModuleBasic interface that defines the
// independent methods a Cosmos SDK module needs to implement.
type AppModuleBasic struct {
	cdc codec.BinaryCodec
}

func NewAppModuleBasic(cdc codec.BinaryCodec) AppModuleBasic {
	return AppModuleBasic{cdc: cdc}
}

// Name returns the name of the module as a string.
func (AppModuleBasic) Name() string {
	return types.ModuleName
}

// RegisterLegacyAminoCodec registers the amino codec for the module, which is used
// to marshal and unmarshal structs to/from []byte in order to persist them in the module's KVStore.
func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {}

// RegisterInterfaces registers a module's interface types and their concrete implementations as proto.Message.
func (a AppModuleBasic) RegisterInterfaces(reg cdctypes.InterfaceRegistry) {
	types.RegisterInterfaces(reg)
}

// DefaultGenesis returns a default GenesisState for the module, marshalled to json.RawMessage.
// The default GenesisState need to be defined by the module developer and is primarily used for testing.
func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	return cdc.MustMarshalJSON(types.DefaultGenesis())
}

// ValidateGenesis used to validate the GenesisState, given in its json.RawMessage form.
func (AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, config client.TxEncodingConfig, bz json.RawMessage) error {
	var genState types.GenesisState
	if err := cdc.UnmarshalJSON(bz, &genState); err != nil {
		return fmt.Errorf("failed to unmarshal %s genesis state: %w", types.ModuleName, err)
	}
	return genState.Validate()
}

// RegisterGRPCGatewayRoutes registers the gRPC Gateway routes for the module.
func (AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {
	if err := types.RegisterQueryHandlerClient(context.Background(), mux, types.NewQueryClient(clientCtx)); err != nil {
		panic(err)
	}
}

// ----------------------------------------------------------------------------
// AppModule
// ----------------------------------------------------------------------------

// AppModule implements the AppModule interface that defines the inter-dependent methods that modules need to implement
type AppModule struct {
	AppModuleBasic

	keeper        keeper.Keeper
	accountKeeper types.AccountKeeper
	bankKeeper    types.BankKeeper
}

func NewAppModule(
	cdc codec.Codec,
	keeper keeper.Keeper,
	accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
) AppModule {
	return AppModule{
		AppModuleBasic: NewAppModuleBasic(cdc),
		keeper:         keeper,
		accountKeeper:  accountKeeper,
		bankKeeper:     bankKeeper,
	}
}

// RegisterServices registers a gRPC query service to respond to the module-specific gRPC queries
func (am AppModule) RegisterServices(cfg module.Configurator) {
	types.RegisterMsgServer(cfg.MsgServer(), keeper.NewMsgServerImpl(am.keeper))
	types.RegisterQueryServer(cfg.QueryServer(), am.keeper)
}

// RegisterInvariants registers the invariants of the module. If an invariant deviates from its predicted value, the InvariantRegistry triggers appropriate logic (most often the chain will be halted)
func (am AppModule) RegisterInvariants(_ sdk.InvariantRegistry) {}

// InitGenesis performs the module's genesis initialization. It returns no validator updates.
func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, gs json.RawMessage) {
	var genState types.GenesisState
	// Initialize global index to index in genesis state
	cdc.MustUnmarshalJSON(gs, &genState)

	InitGenesis(ctx, am.keeper, genState)
}

// ExportGenesis returns the module's exported genesis state as raw JSON bytes.
func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	genState := ExportGenesis(ctx, am.keeper)
	return cdc.MustMarshalJSON(genState)
}

// ConsensusVersion is a sequence number for state-breaking change of the module.
// It should be incremented on each consensus-breaking change introduced by the module.
// To avoid wrong/empty versions, the initial version should be set to 1.
func (AppModule) ConsensusVersion() uint64 { return 1 }

// BeginBlock contains the logic that is automatically triggered at the beginning of each block.
// The begin block implementation is optional.
func (am AppModule) BeginBlock(_ context.Context) error {
	return nil
}

// EndBlock contains the logic that is automatically triggered at the end of each block.
// The end block implementation is optional.
func (am AppModule) EndBlock(goCtx context.Context) error {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if ctx.BlockHeight()%am.keeper.GetParams(ctx).CardAuctionPriceReductionPeriod == 0 {
		price := am.keeper.CardAuctionPrice.Get(ctx)
		newprice := price.Sub(util.QuoCoin(*price, 100))
		if !newprice.IsLT(sdk.NewInt64Coin("ucredits", 1000000)) { // stop at 1 credit
			am.keeper.CardAuctionPrice.Set(ctx, &newprice)
		}
		am.keeper.Logger().Info(fmt.Sprintf(":: CardAuctionPrice: %s", am.keeper.CardAuctionPrice.Get(ctx)))
	}

	// automated nerf/buff happens here
	if ctx.BlockHeight()%epochBlockTime == 0 {
		am.keeper.UpdateNerfLevels(ctx)
		matchesEnabled, _ := am.keeper.FeatureFlagModuleInstance.Get(ctx, string(types.FeatureFlagName_Matches))
		if matchesEnabled { // Only give voterigths to all users, when matches are not anabled
			am.keeper.AddVoteRightsToAllUsers(ctx)
		}
	}

	if ctx.BlockHeight()%500 == 0 { //HourlyFaucet
		am.keeper.DistributeHourlyFaucet(ctx)

		incentives := util.QuoCoin(*am.keeper.Pools.Get(ctx, keeper.PublicPoolKey), 10)
		am.keeper.SubPoolCredits(ctx, keeper.PublicPoolKey, incentives)
		winnersIncentives := util.MulCoinFloat(incentives, float64(am.keeper.GetWinnerIncentives(ctx)))
		balancersIncentives := util.MulCoinFloat(incentives, float64(am.keeper.GetBalancerIncentives(ctx)))
		am.keeper.Logger().Info(fmt.Sprintf(":: Incentives (w, b): %s, %s", winnersIncentives, balancersIncentives))
		am.keeper.AddPoolCredits(ctx, keeper.WinnersPoolKey, winnersIncentives)
		am.keeper.AddPoolCredits(ctx, keeper.BalancersPoolKey, balancersIncentives)
		am.keeper.Logger().Info(fmt.Sprintf(":: PublicPool: %s", am.keeper.Pools.Get(ctx, keeper.PublicPoolKey)))
	}

	// Setting running averages
	if ctx.BlockHeight()%500 == 0 {
		iter := am.keeper.RunningAverages.GetItemIterator(ctx)
		for ; iter.Valid(); iter.Next() {
			idx, average := iter.Value()
			average.Arr = append(average.Arr, 0)
			if len(average.Arr) > 24 {
				average.Arr = average.Arr[1:]
			}
			am.keeper.RunningAverages.Set(ctx, am.keeper.RunningAverages.KeyWords[idx], average)
		}
	}

	err := am.keeper.CheckTrial(ctx)
	if err != nil {
		am.keeper.Logger().Error(fmt.Sprintf("%s", err))
	}

	// Checks for empty pools
	for idx, coin := range am.keeper.Pools.GetAll(ctx) {
		if coin.IsLT(sdk.NewInt64Coin(coin.Denom, 0)) {
			am.keeper.Logger().Error(fmt.Sprintf(":: %s: %v", am.keeper.Pools.KeyWords[idx], coin))
		}
	}

	am.keeper.MatchWorker(ctx)

	return nil
}

// IsOnePerModuleType implements the depinject.OnePerModuleType interface.
func (am AppModule) IsOnePerModuleType() {}

// IsAppModule implements the appmodule.AppModule interface.
func (am AppModule) IsAppModule() {}

// ----------------------------------------------------------------------------
// App Wiring Setup
// ----------------------------------------------------------------------------

func init() {
	appmodule.Register(
		&modulev1.Module{},
		appmodule.Provide(ProvideModule),
	)
}

type ModuleInputs struct {
	depinject.In

	StoreService store.KVStoreService
	Cdc          codec.Codec
	Config       *modulev1.Module
	Logger       log.Logger

	AccountKeeper     types.AccountKeeper
	BankKeeper        types.BankKeeper
	FeatureFlagKeeper types.FeatureFlagKeeper
}

type ModuleOutputs struct {
	depinject.Out

	CardchainKeeper keeper.Keeper
	Module          appmodule.AppModule
}

func ProvideModule(in ModuleInputs) ModuleOutputs {
	// default to governance authority if not provided
	authority := authtypes.NewModuleAddress(govtypes.ModuleName)
	if in.Config.Authority != "" {
		authority = authtypes.NewModuleAddressOrBech32Address(in.Config.Authority)
	}
	k := keeper.NewKeeper(
		in.Cdc,
		in.StoreService,
		in.Logger,
		authority.String(),
		in.BankKeeper,
		in.FeatureFlagKeeper,
	)
	m := NewAppModule(
		in.Cdc,
		k,
		in.AccountKeeper,
		in.BankKeeper,
	)

	return ModuleOutputs{CardchainKeeper: k, Module: m}
}
