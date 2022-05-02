package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSubmitMatchReporterProposal = "submit_match_reporter_proposal"

var _ sdk.Msg = &MsgSubmitMatchReporterProposal{}

func NewMsgSubmitMatchReporterProposal(creator string, reporter string, deposit string, description string) *MsgSubmitMatchReporterProposal {
	return &MsgSubmitMatchReporterProposal{
		Creator:     creator,
		Reporter:    reporter,
		Deposit:     deposit,
		Description: description,
	}
}

func (msg *MsgSubmitMatchReporterProposal) Route() string {
	return RouterKey
}

func (msg *MsgSubmitMatchReporterProposal) Type() string {
	return TypeMsgSubmitMatchReporterProposal
}

func (msg *MsgSubmitMatchReporterProposal) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSubmitMatchReporterProposal) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSubmitMatchReporterProposal) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
