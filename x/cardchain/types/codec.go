package types

import (
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	// this line is used by starport scaffolding # 1
)

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
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCardTransfer{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCardDonate{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCardArtworkAdd{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCardArtistChange{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCouncilRegister{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCouncilDeregister{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgMatchReport{},
	)
	// this line is used by starport scaffolding # 3

	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateParams{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
