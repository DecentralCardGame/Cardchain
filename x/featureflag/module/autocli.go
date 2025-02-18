package featureflag

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "github.com/DecentralCardGame/cardchain/api/cardchain/featureflag"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod:      "Flag",
					Use:            "flag [module] [name]",
					Short:          "Query flag",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "module"}, {ProtoField: "name"}},
				},

				{
					RpcMethod:      "Flags",
					Use:            "flags",
					Short:          "Query flags",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{},
				},

				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "Set",
					Use:            "set [module] [name] [value]",
					Short:          "Send a Set tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "module"}, {ProtoField: "name"}, {ProtoField: "value"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
