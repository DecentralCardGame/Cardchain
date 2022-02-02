package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSubmitCopyrightProposal = "submit_copyright_proposal"

var _ sdk.Msg = &MsgSubmitCopyrightProposal{}

func NewMsgSubmitCopyrightProposal(creator string, cardId uint64, description string, link string) *MsgSubmitCopyrightProposal {
	return &MsgSubmitCopyrightProposal{
		Creator:     creator,
		CardId:      cardId,
		Description: description,
		Link:        link,
	}
}

func (msg *MsgSubmitCopyrightProposal) Route() string {
	return RouterKey
}

func (msg *MsgSubmitCopyrightProposal) Type() string {
	return TypeMsgSubmitCopyrightProposal
}

func (msg *MsgSubmitCopyrightProposal) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSubmitCopyrightProposal) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSubmitCopyrightProposal) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
