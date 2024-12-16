package cardchain

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "github.com/DecentralCardGame/cardchain/api/cardchain/cardchain"
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
					RpcMethod:      "UserCreate",
					Use:            "user-create [new-user] [alias]",
					Short:          "Send a UserCreate tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "newUser"}, {ProtoField: "alias"}},
				},
				{
					RpcMethod:      "CardSchemeBuy",
					Use:            "card-scheme-buy [bid]",
					Short:          "Send a CardSchemeBuy tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "bid"}},
				},
				{
					RpcMethod:      "CardSaveContent",
					Use:            "card-save-content [card-id] [content] [notes] [artist] [balance-anchor]",
					Short:          "Send a CardSaveContent tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "cardId"}, {ProtoField: "content"}, {ProtoField: "notes"}, {ProtoField: "artist"}, {ProtoField: "balanceAnchor"}},
				},
				{
			RpcMethod: "CardVote",
			Use: "card-vote [vote]",
			Short: "Send a CardVote tx",
			PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "vote"},},
		},
		// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
