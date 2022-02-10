package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgBuyCollection = "buy_collection"

var _ sdk.Msg = &MsgBuyCollection{}

func NewMsgBuyCollection(creator string) *MsgBuyCollection {
	return &MsgBuyCollection{
		Creator: creator,
	}
}

func (msg *MsgBuyCollection) Route() string {
	return RouterKey
}

func (msg *MsgBuyCollection) Type() string {
	return TypeMsgBuyCollection
}

func (msg *MsgBuyCollection) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgBuyCollection) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgBuyCollection) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
