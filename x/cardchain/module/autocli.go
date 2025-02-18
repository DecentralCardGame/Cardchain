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

				{
					RpcMethod:      "Card",
					Use:            "card [card-id]",
					Short:          "Query card",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "cardId"}},
				},

				{
					RpcMethod:      "User",
					Use:            "user [address]",
					Short:          "Query user",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "address"}},
				},

				{
					RpcMethod:      "Cards",
					Use:            "cards [owner]",
					Short:          "Query cards",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "owner"}},
				},

				{
					RpcMethod:      "Match",
					Use:            "match [match-id]",
					Short:          "Query match",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "matchId"}},
				},

				{
					RpcMethod:      "Set",
					Use:            "set [set-id]",
					Short:          "Query set",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "setId"}},
				},

				{
					RpcMethod:      "SellOffer",
					Use:            "sell-offer [sell-offer-id]",
					Short:          "Query sellOffer",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "sellOfferId"}},
				},

				{
					RpcMethod:      "Council",
					Use:            "council [council-id]",
					Short:          "Query council",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "councilId"}},
				},

				{
					RpcMethod:      "Server",
					Use:            "server [server-id]",
					Short:          "Query server",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "serverId"}},
				},

				{
					RpcMethod:      "Encounter",
					Use:            "encounter [encounter-id]",
					Short:          "Query encounter",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "encounterId"}},
				},

				{
					RpcMethod:      "Encounters",
					Use:            "encounters",
					Short:          "Query encounters",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{},
				},

				{
					RpcMethod:      "EncounterWithImage",
					Use:            "encounter-with-image",
					Short:          "Query encounterWithImage",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{},
				},

				{
					RpcMethod:      "EncountersWithImage",
					Use:            "encounters-with-image",
					Short:          "Query encountersWithImage",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{},
				},

				{
					RpcMethod:      "CardchainInfo",
					Use:            "cardchain-info",
					Short:          "Query cardchainInfo",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{},
				},

				{
					RpcMethod:      "SetRarityDistribution",
					Use:            "set-rarity-distribution [set-id]",
					Short:          "Query setRarityDistribution",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "setId"}},
				},

				{
					RpcMethod:      "AccountFromZealy",
					Use:            "account-from-zealy [zealy-id]",
					Short:          "Query accountFromZealy",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "zealyId"}},
				},

				{
					RpcMethod:      "VotingResults",
					Use:            "voting-results",
					Short:          "Query votingResults",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{},
				},

				{
					RpcMethod:      "Matches",
					Use:            "matches [timestamp-down] [timestamp-up] [contains-users] [reporter] [outcome] [cards-played]",
					Short:          "Query matches",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "timestampDown"}, {ProtoField: "timestampUp"}, {ProtoField: "containsUsers"}, {ProtoField: "reporter"}, {ProtoField: "outcome"}, {ProtoField: "cardsPlayed"}},
				},

				{
					RpcMethod:      "Sets",
					Use:            "sets [status] [contributors] [contains-cards] [owner]",
					Short:          "Query sets",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "status"}, {ProtoField: "contributors"}, {ProtoField: "containsCards"}, {ProtoField: "owner"}},
				},

				{
					RpcMethod:      "CardContent",
					Use:            "card-content [card-id]",
					Short:          "Query cardContent",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "cardId"}},
				},

				{
					RpcMethod:      "CardContents",
					Use:            "card-contents [card-ids]",
					Short:          "Query cardContents",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "cardIds"}},
				},

				{
					RpcMethod:      "SellOffers",
					Use:            "sell-offers [price-down] [price-up] [seller] [buyer] [card] [status]",
					Short:          "Query sellOffers",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "priceDown"}, {ProtoField: "priceUp"}, {ProtoField: "seller"}, {ProtoField: "buyer"}, {ProtoField: "card"}, {ProtoField: "status"}},
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
				{
					RpcMethod:      "CardRaritySet",
					Use:            "card-rarity-set [card-id] [set-id] [rarity]",
					Short:          "Send a CardRaritySet tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "cardId"}, {ProtoField: "setId"}, {ProtoField: "rarity"}},
				},
				{
					RpcMethod:      "CouncilResponseCommit",
					Use:            "council-response-commit [council-id] [response] [suggestion]",
					Short:          "Send a CouncilResponseCommit tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "councilId"}, {ProtoField: "response"}, {ProtoField: "suggestion"}},
				},
				{
					RpcMethod:      "CouncilResponseReveal",
					Use:            "council-response-reveal [council-id] [response] [secret]",
					Short:          "Send a CouncilResponseReveal tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "councilId"}, {ProtoField: "response"}, {ProtoField: "secret"}},
				},
				{
					RpcMethod:      "CouncilRestart",
					Use:            "council-restart [council-id]",
					Short:          "Send a CouncilRestart tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "councilId"}},
				},
				{
					RpcMethod:      "MatchConfirm",
					Use:            "match-confirm [match-id] [outcome] [voted-cards]",
					Short:          "Send a MatchConfirm tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "matchId"}, {ProtoField: "outcome"}, {ProtoField: "votedCards"}},
				},
				{
					RpcMethod:      "ProfileCardSet",
					Use:            "profile-card-set [card-id]",
					Short:          "Send a ProfileCardSet tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "cardId"}},
				},
				{
					RpcMethod:      "ProfileWebsiteSet",
					Use:            "profile-website-set [website]",
					Short:          "Send a ProfileWebsiteSet tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "website"}},
				},
				{
					RpcMethod:      "ProfileBioSet",
					Use:            "profile-bio-set [bio]",
					Short:          "Send a ProfileBioSet tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "bio"}},
				},
				{
					RpcMethod:      "BoosterPackOpen",
					Use:            "booster-pack-open [booster-pack-id]",
					Short:          "Send a BoosterPackOpen tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "boosterPackId"}},
				},
				{
					RpcMethod:      "BoosterPackTransfer",
					Use:            "booster-pack-transfer [booster-pack-id] [receiver]",
					Short:          "Send a BoosterPackTransfer tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "boosterPackId"}, {ProtoField: "receiver"}},
				},
				{
					RpcMethod:      "SetStoryWriterSet",
					Use:            "set-story-writer-set [set-id] [story-writer]",
					Short:          "Send a SetStoryWriterSet tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "setId"}, {ProtoField: "storyWriter"}},
				},
				{
					RpcMethod:      "SetArtistSet",
					Use:            "set-artist-set [set-id] [artist]",
					Short:          "Send a SetArtistSet tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "setId"}, {ProtoField: "artist"}},
				},
				{
					RpcMethod:      "CardVoteMulti",
					Use:            "card-vote-multi [votes]",
					Short:          "Send a CardVoteMulti tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "votes"}},
				},
				{
					RpcMethod:      "MatchOpen",
					Use:            "match-open [player-a] [player-b] [player-a-deck] [player-b-deck]",
					Short:          "Send a MatchOpen tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "playerA"}, {ProtoField: "playerB"}, {ProtoField: "playerADeck"}, {ProtoField: "playerBDeck"}},
				},
				{
					RpcMethod:      "SetNameSet",
					Use:            "set-name-set [set-id] [name]",
					Short:          "Send a SetNameSet tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "setId"}, {ProtoField: "name"}},
				},
				{
					RpcMethod:      "ProfileAliasSet",
					Use:            "profile-alias-set [alias]",
					Short:          "Send a ProfileAliasSet tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "alias"}},
				},
				{
					RpcMethod:      "EarlyAccessInvite",
					Use:            "early-access-invite [user]",
					Short:          "Send a EarlyAccessInvite tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "user"}},
				},
				{
					RpcMethod:      "ZealyConnect",
					Use:            "zealy-connect [zealy-id]",
					Short:          "Send a ZealyConnect tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "zealyId"}},
				},
				{
					RpcMethod:      "EncounterCreate",
					Use:            "encounter-create [name] [drawlist] [parameters] [image]",
					Short:          "Send a EncounterCreate tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "name"}, {ProtoField: "drawlist"}, {ProtoField: "parameters"}, {ProtoField: "image"}},
				},
				{
					RpcMethod:      "EncounterDo",
					Use:            "encounter-do [encounter-id] [user]",
					Short:          "Send a EncounterDo tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "encounterId"}, {ProtoField: "user"}},
				},
				{
					RpcMethod:      "EncounterClose",
					Use:            "encounter-close [encounter-id] [user] [won]",
					Short:          "Send a EncounterClose tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "encounterId"}, {ProtoField: "user"}, {ProtoField: "won"}},
				},

				{
					RpcMethod:      "EarlyAccessDisinvite",
					Use:            "early-access-disinvite [user]",
					Short:          "Send a EarlyAccessDisinvite tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "user"}},
				},
				{
					RpcMethod:      "CardBan",
					Use:            "ban-card [card-id]",
					Short:          "Send a CardBan tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "cardId"}},
				},
				{
					RpcMethod:      "EarlyAccessGrant",
					Use:            "early-access-grant [users]",
					Short:          "Send a EarlyAccessGrant tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "users"}},
				},
				{
					RpcMethod:      "SetActivate",
					Use:            "set-activate [set-id]",
					Short:          "Send a SetActivate tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "setId"}},
				},
				{
					RpcMethod:      "CardCopyrightClaim",
					Use:            "card-copyright-claim [card-id]",
					Short:          "Send a CardCopyrightClaim tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "cardId"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
