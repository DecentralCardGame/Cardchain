package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateuser{}, "cardchain/Createuser", nil)
	cdc.RegisterConcrete(&MsgBuyCardScheme{}, "cardchain/BuyCardScheme", nil)
	cdc.RegisterConcrete(&MsgVoteCard{}, "cardchain/VoteCard", nil)
	cdc.RegisterConcrete(&MsgSaveCardContent{}, "cardchain/SaveCardContent", nil)
	cdc.RegisterConcrete(&MsgTransferCard{}, "cardchain/TransferCard", nil)
	cdc.RegisterConcrete(&MsgDonateToCard{}, "cardchain/DonateToCard", nil)
	cdc.RegisterConcrete(&MsgAddArtwork{}, "cardchain/AddArtwork", nil)
	cdc.RegisterConcrete(&CopyrightProposal{}, "cardchain/CopyrightProposal", nil)
	cdc.RegisterConcrete(&MatchReporterProposal{}, "cardchain/MatchReporterProposal", nil)
	cdc.RegisterConcrete(&SetProposal{}, "cardchain/SetProposal", nil)
	cdc.RegisterConcrete(&MsgChangeArtist{}, "cardchain/ChangeArtist", nil)
	cdc.RegisterConcrete(&MsgRegisterForCouncil{}, "cardchain/RegisterForCouncil", nil)
	cdc.RegisterConcrete(&MsgReportMatch{}, "cardchain/ReportMatch", nil)
	cdc.RegisterConcrete(&MsgApointMatchReporter{}, "cardchain/ApointMatchReporter", nil)
	cdc.RegisterConcrete(&MsgCreateSet{}, "cardchain/CreateSet", nil)
	cdc.RegisterConcrete(&MsgAddCardToSet{}, "cardchain/AddCardToSet", nil)
	cdc.RegisterConcrete(&MsgFinalizeSet{}, "cardchain/FinalizeSet", nil)
	cdc.RegisterConcrete(&MsgBuyBoosterPack{}, "cardchain/BuyBoosterPack", nil)
	cdc.RegisterConcrete(&MsgRemoveCardFromSet{}, "cardchain/RemoveCardFromSet", nil)
	cdc.RegisterConcrete(&MsgRemoveContributorFromSet{}, "cardchain/RemoveContributorFromSet", nil)
	cdc.RegisterConcrete(&MsgAddContributorToSet{}, "cardchain/AddContributorToSet", nil)
	cdc.RegisterConcrete(&MsgCreateSellOffer{}, "cardchain/CreateSellOffer", nil)
	cdc.RegisterConcrete(&MsgBuyCard{}, "cardchain/BuyCard", nil)
	cdc.RegisterConcrete(&MsgRemoveSellOffer{}, "cardchain/RemoveSellOffer", nil)
	cdc.RegisterConcrete(&MsgAddArtworkToSet{}, "cardchain/AddArtworkToSet", nil)
	cdc.RegisterConcrete(&MsgAddStoryToSet{}, "cardchain/AddStoryToSet", nil)
	cdc.RegisterConcrete(&MsgSetCardRarity{}, "cardchain/SetCardRarity", nil)
	cdc.RegisterConcrete(&MsgCreateCouncil{}, "cardchain/CreateCouncil", nil)
	cdc.RegisterConcrete(&MsgCommitCouncilResponse{}, "cardchain/CommitCouncilResponse", nil)
	cdc.RegisterConcrete(&MsgRevealCouncilResponse{}, "cardchain/RevealCouncilResponse", nil)
	cdc.RegisterConcrete(&MsgRestartCouncil{}, "cardchain/RestartCouncil", nil)
	cdc.RegisterConcrete(&MsgRewokeCouncilRegistration{}, "cardchain/RewokeCouncilRegistration", nil)
	cdc.RegisterConcrete(&MsgConfirmMatch{}, "cardchain/ConfirmMatch", nil)
	cdc.RegisterConcrete(&MsgSetProfileCard{}, "cardchain/SetProfileCard", nil)
	cdc.RegisterConcrete(&MsgOpenBoosterPack{}, "cardchain/OpenBoosterPack", nil)
	cdc.RegisterConcrete(&MsgTransferBoosterPack{}, "cardchain/TransferBoosterPack", nil)
	cdc.RegisterConcrete(&MsgSetSetStoryWriter{}, "cardchain/SetSetStoryWriter", nil)
	cdc.RegisterConcrete(&MsgSetSetArtist{}, "cardchain/SetSetArtist", nil)
	cdc.RegisterConcrete(&MsgSetUserWebsite{}, "cardchain/SetUserWebsite", nil)
	cdc.RegisterConcrete(&MsgSetUserBiography{}, "cardchain/SetUserBiography", nil)
	cdc.RegisterConcrete(&MsgMultiVoteCard{}, "cardchain/MultiVoteCard", nil)
	cdc.RegisterConcrete(&MsgOpenMatch{}, "cardchain/MsgOpenMatch", nil)
	cdc.RegisterConcrete(&MsgSetSetName{}, "cardchain/SetSetName", nil)
	cdc.RegisterConcrete(&MsgChangeAlias{}, "cardchain/ChangeAlias", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateuser{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgBuyCardScheme{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgVoteCard{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSaveCardContent{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgTransferCard{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDonateToCard{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddArtwork{},
	)
	registry.RegisterImplementations((*govtypes.Content)(nil),
		&CopyrightProposal{},
	)
	registry.RegisterImplementations((*govtypes.Content)(nil),
		&MatchReporterProposal{},
	)
	registry.RegisterImplementations((*govtypes.Content)(nil),
		&SetProposal{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgChangeArtist{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRegisterForCouncil{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgReportMatch{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgApointMatchReporter{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateSet{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddCardToSet{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgFinalizeSet{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgBuyBoosterPack{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRemoveCardFromSet{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRemoveContributorFromSet{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddContributorToSet{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateSellOffer{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgBuyCard{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRemoveSellOffer{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddArtworkToSet{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddStoryToSet{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetCardRarity{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateCouncil{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCommitCouncilResponse{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRevealCouncilResponse{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRestartCouncil{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRewokeCouncilRegistration{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgConfirmMatch{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetProfileCard{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgOpenBoosterPack{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgTransferBoosterPack{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetSetStoryWriter{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetSetArtist{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetUserWebsite{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetUserBiography{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgMultiVoteCard{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgOpenMatch{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetSetName{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgChangeAlias{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
