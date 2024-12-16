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
					RpcMethod:      "CardVote",
					Use:            "card-vote [vote]",
					Short:          "Send a CardVote tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "vote"}},
				},
				{
					RpcMethod:      "CardTransfer",
					Use:            "card-transfer [card-id] [receiver]",
					Short:          "Send a CardTransfer tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "cardId"}, {ProtoField: "receiver"}},
				},
				{
					RpcMethod:      "CardDonate",
					Use:            "card-donate [card-id] [amount]",
					Short:          "Send a CardDonate tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "cardId"}, {ProtoField: "amount"}},
				},
				{
					RpcMethod:      "CardArtworkAdd",
					Use:            "card-artwork-add [card-id] [image] [full-art]",
					Short:          "Send a CardArtworkAdd tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "cardId"}, {ProtoField: "image"}, {ProtoField: "fullArt"}},
				},
				{
					RpcMethod:      "CardArtistChange",
					Use:            "card-artist-change [card-id] [artist]",
					Short:          "Send a CardArtistChange tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "cardId"}, {ProtoField: "artist"}},
				},
				{
					RpcMethod:      "CouncilRegister",
					Use:            "council-register",
					Short:          "Send a CouncilRegister tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{},
				},
				{
					RpcMethod:      "CouncilDeregister",
					Use:            "council-deregister",
					Short:          "Send a CouncilDeregister tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{},
				},
				{
					RpcMethod:      "MatchReport",
					Use:            "match-report [match-id] [played-cards-a] [played-cards-b] [outcome]",
					Short:          "Send a MatchReport tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "matchId"}, {ProtoField: "playedCardsA"}, {ProtoField: "playedCardsB"}, {ProtoField: "outcome"}},
				},
				{
					RpcMethod:      "CouncilCreate",
					Use:            "council-create [card-id]",
					Short:          "Send a CouncilCreate tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "cardId"}},
				},
				{
					RpcMethod:      "MatchReporterAppoint",
					Use:            "match-reporter-appoint [reporter]",
					Short:          "Send a MatchReporterAppoint tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "reporter"}},
				},
				{
					RpcMethod:      "SetCreate",
					Use:            "set-create [name] [artist] [story-writer] [contributors]",
					Short:          "Send a SetCreate tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "name"}, {ProtoField: "artist"}, {ProtoField: "storyWriter"}, {ProtoField: "contributors"}},
				},
				{
					RpcMethod:      "SetCardAdd",
					Use:            "set-card-add [set-id] [card-id]",
					Short:          "Send a SetCardAdd tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "setId"}, {ProtoField: "cardId"}},
				},
				{
					RpcMethod:      "SetCardRemove",
					Use:            "set-card-remove [set-id] [card-id]",
					Short:          "Send a SetCardRemove tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "setId"}, {ProtoField: "cardId"}},
				},
				{
					RpcMethod:      "SetContributorAdd",
					Use:            "set-contributor-add [set-id] [user]",
					Short:          "Send a SetContributorAdd tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "setId"}, {ProtoField: "user"}},
				},
				{
					RpcMethod:      "SetContributorRemove",
					Use:            "set-contributor-remove [set-id] [user]",
					Short:          "Send a SetContributorRemove tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "setId"}, {ProtoField: "user"}},
				},
				{
					RpcMethod:      "SetFinalize",
					Use:            "set-finalize [set-id]",
					Short:          "Send a SetFinalize tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "setId"}},
				},
				{
					RpcMethod:      "SetArtworkAdd",
					Use:            "set-artwork-add [set-id] [image]",
					Short:          "Send a SetArtworkAdd tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "setId"}, {ProtoField: "image"}},
				},
				{
					RpcMethod:      "SetStoryAdd",
					Use:            "set-story-add [set-id] [story]",
					Short:          "Send a SetStoryAdd tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "setId"}, {ProtoField: "story"}},
				},
				{
					RpcMethod:      "BoosterPackBuy",
					Use:            "booster-pack-buy [set-id]",
					Short:          "Send a BoosterPackBuy tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "setId"}},
				},
				{
					RpcMethod:      "SellOfferCreate",
					Use:            "sell-offer-create [card-id] [price]",
					Short:          "Send a SellOfferCreate tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "cardId"}, {ProtoField: "price"}},
				},
				{
					RpcMethod:      "SellOfferBuy",
					Use:            "sell-offer-buy [sell-offer-id]",
					Short:          "Send a SellOfferBuy tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "sellOfferId"}},
				},
				{
					RpcMethod:      "SellOfferRemove",
					Use:            "sell-offer-remove [sell-offer-id]",
					Short:          "Send a SellOfferRemove tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "sellOfferId"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
