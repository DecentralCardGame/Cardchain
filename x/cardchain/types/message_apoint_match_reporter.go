package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgApointMatchReporter = "apoint_match_reporter"

var _ sdk.Msg = &MsgApointMatchReporter{}

func NewMsgApointMatchReporter(creator string, reporter string) *MsgApointMatchReporter {
	return &MsgApointMatchReporter{
		Creator:  creator,
		Reporter: reporter,
	}
}

func (msg *MsgApointMatchReporter) Route() string {
	return RouterKey
}

func (msg *MsgApointMatchReporter) Type() string {
	return TypeMsgApointMatchReporter
}

func (msg *MsgApointMatchReporter) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgApointMatchReporter) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgApointMatchReporter) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
