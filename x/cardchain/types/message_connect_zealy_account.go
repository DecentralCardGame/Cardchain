package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgConnectZealyAccount = "connect_zealy_account"

var _ sdk.Msg = &MsgConnectZealyAccount{}

func NewMsgConnectZealyAccount(creator string, zealyId string) *MsgConnectZealyAccount {
	return &MsgConnectZealyAccount{
		Creator: creator,
		ZealyId: zealyId,
	}
}

func (msg *MsgConnectZealyAccount) Route() string {
	return RouterKey
}

func (msg *MsgConnectZealyAccount) Type() string {
	return TypeMsgConnectZealyAccount
}

func (msg *MsgConnectZealyAccount) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgConnectZealyAccount) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgConnectZealyAccount) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
