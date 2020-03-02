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
	cdc.RegisterConcrete(MsgSetName{}, "cardservice/SetName", nil)
	cdc.RegisterConcrete(MsgBuyName{}, "cardservice/BuyName", nil)
	cdc.RegisterConcrete(MsgDeleteName{}, "cardservice/DeleteName", nil)
}
