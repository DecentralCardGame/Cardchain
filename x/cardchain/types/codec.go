package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgUserCreate{}, "cardchain/UserCreate", nil)
	cdc.RegisterConcrete(&MsgCardSchemeBuy{}, "cardchain/CardSchemeBuy", nil)
	cdc.RegisterConcrete(&MsgCardSaveContent{}, "cardchain/CardSaveContent", nil)
	cdc.RegisterConcrete(&MsgCardVote{}, "cardchain/CardVote", nil)
// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUserCreate{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCardSchemeBuy{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCardSaveContent{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
	&MsgCardVote{},
)
// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
