package nameservice

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on the Amino codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgSetName{}, "nameservice/SetName", nil)
	cdc.RegisterConcrete(MsgBuyCardScheme{}, "nameservice/BuyCardScheme", nil)
	cdc.RegisterConcrete(MsgSetType{}, "nameservice/SetType", nil)
}
