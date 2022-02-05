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
	cdc.RegisterConcrete(&MsgChangeArtist{}, "cardchain/ChangeArtist", nil)
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
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgChangeArtist{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
