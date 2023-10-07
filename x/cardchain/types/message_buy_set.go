package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgBuyBoosterPack = "buy_set"

var _ sdk.Msg = &MsgBuyBoosterPack{}

func NewMsgBuyBoosterPack(creator string, setId uint64) *MsgBuyBoosterPack {
	return &MsgBuyBoosterPack{
		Creator: creator,
		SetId:   setId,
	}
}

func (msg *MsgBuyBoosterPack) Route() string {
	return RouterKey
}

func (msg *MsgBuyBoosterPack) Type() string {
	return TypeMsgBuyBoosterPack
}

func (msg *MsgBuyBoosterPack) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgBuyBoosterPack) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgBuyBoosterPack) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(errors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
