package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateuser{}, "cardchain/Createuser", nil)
	cdc.RegisterConcrete(&MsgBuyCardScheme{}, "cardchain/BuyCardScheme", nil)
	cdc.RegisterConcrete(&MsgVoteCard{}, "cardchain/VoteCard", nil)
	cdc.RegisterConcrete(&MsgSaveCardContent{}, "cardchain/SaveCardContent", nil)
	cdc.RegisterConcrete(&MsgTransferCard{}, "cardchain/TransferCard", nil)
	cdc.RegisterConcrete(&MsgDonateToCard{}, "cardchain/DonateToCard", nil)
	cdc.RegisterConcrete(&MsgAddArtwork{}, "cardchain/AddArtwork", nil)
	cdc.RegisterConcrete(&MsgSubmitCopyrightProposal{}, "cardchain/SubmitCopyrightProposal", nil)
	cdc.RegisterConcrete(&CopyrightProposal{}, "cardchain/CopyrightProposal", nil)
	cdc.RegisterConcrete(&MatchReporterProposal{}, "cardchain/MatchReporterProposal", nil)
	cdc.RegisterConcrete(&MsgChangeArtist{}, "cardchain/ChangeArtist", nil)
	cdc.RegisterConcrete(&MsgRegisterForCouncil{}, "cardchain/RegisterForCouncil", nil)
	cdc.RegisterConcrete(&MsgReportMatch{}, "cardchain/ReportMatch", nil)
	cdc.RegisterConcrete(&MsgSubmitMatchReporterProposal{}, "cardchain/SubmitMatchReporterProposal", nil)
	cdc.RegisterConcrete(&MsgApointMatchReporter{}, "cardchain/ApointMatchReporter", nil)
	cdc.RegisterConcrete(&MsgCreateCollection{}, "cardchain/CreateCollection", nil)
	cdc.RegisterConcrete(&MsgAddCardToCollection{}, "cardchain/AddCardToCollection", nil)
	cdc.RegisterConcrete(&MsgFinalizeCollection{}, "cardchain/FinalizeCollection", nil)
	cdc.RegisterConcrete(&MsgBuyCollection{}, "cardchain/BuyCollection", nil)
	cdc.RegisterConcrete(&MsgRemoveCardFromCollection{}, "cardchain/RemoveCardFromCollection", nil)
	cdc.RegisterConcrete(&MsgRemoveContributorFromCollection{}, "cardchain/RemoveContributorFromCollection", nil)
	cdc.RegisterConcrete(&MsgAddContributorToCollection{}, "cardchain/AddContributorToCollection", nil)
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
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSubmitCopyrightProposal{},
	)
	registry.RegisterImplementations((*govtypes.Content)(nil),
		&CopyrightProposal{},
	)
	registry.RegisterImplementations((*govtypes.Content)(nil),
		&MatchReporterProposal{},
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
		&MsgSubmitMatchReporterProposal{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgApointMatchReporter{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateCollection{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddCardToCollection{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgFinalizeCollection{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgBuyCollection{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRemoveCardFromCollection{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRemoveContributorFromCollection{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddContributorToCollection{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
