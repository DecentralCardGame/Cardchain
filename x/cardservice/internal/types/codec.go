package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// ModuleCdc is the codec for the module
var ModuleCdc = codec.New()

func init() {
	RegisterCodec(ModuleCdc)
}

// RegisterCodec registers concrete types on the Amino codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgBuyCardScheme{}, "cardservice/BuyCardScheme", nil)
	cdc.RegisterConcrete(MsgSaveCardContent{}, "cardservice/SaveCardContent", nil)
	cdc.RegisterConcrete(MsgVoteCard{}, "cardservice/VoteCard", nil)
	cdc.RegisterConcrete(MsgTransferCard{}, "cardservice/TransferCard", nil)
	cdc.RegisterConcrete(MsgDonateToCard{}, "cardservice/DonateToCard", nil)
	cdc.RegisterConcrete(MsgCreateUser{}, "cardservice/CreateUser", nil)
}
