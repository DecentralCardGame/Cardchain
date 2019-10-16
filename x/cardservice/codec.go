package cardservice

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on the Amino codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgBuyCardScheme{}, "cardservice/BuyCardScheme", nil)
	cdc.RegisterConcrete(MsgSaveCardContent{}, "cardservice/SaveCardContent", nil)
	cdc.RegisterConcrete(MsgVoteCard{}, "cardservice/VoteCard", nil)
	cdc.RegisterConcrete(MsgTransferCard{}, "cardservice/TransferCard", nil)
	cdc.RegisterConcrete(MsgDonateToCard{}, "cardservice/DonateToCard", nil)
	cdc.RegisterConcrete(MsgCreateUser{}, "cardservice/CreateUser", nil)
	cdc.RegisterConcrete(Card{}, "cardservice/Card", nil)
	cdc.RegisterConcrete(User{}, "cardservice/User", nil)
	cdc.RegisterConcrete(VoteRight{}, "cardservice/VoteRight", nil)
}
