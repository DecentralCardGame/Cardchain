package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgBuySet = "buy_set"

var _ sdk.Msg = &MsgBuySet{}

func NewMsgBuySet(creator string, setId uint64) *MsgBuySet {
	return &MsgBuySet{
		Creator:      creator,
		SetId: setId,
	}
}

func (msg *MsgBuySet) Route() string {
	return RouterKey
}

func (msg *MsgBuySet) Type() string {
	return TypeMsgBuySet
}

func (msg *MsgBuySet) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgBuySet) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgBuySet) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(errors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
