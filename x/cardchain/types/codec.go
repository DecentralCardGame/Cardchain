package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateuser{}, "cardchain/Createuser", nil)
	cdc.RegisterConcrete(&MsgBuyCardScheme{}, "cardchain/BuyCardScheme", nil)
	cdc.RegisterConcrete(&MsgVoteCard{}, "cardchain/VoteCard", nil)
	cdc.RegisterConcrete(&MsgSaveCardContent{}, "cardchain/SaveCardContent", nil)
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
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)