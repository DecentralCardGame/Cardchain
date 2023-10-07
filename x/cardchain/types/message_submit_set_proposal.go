package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSubmitSetProposal = "submit_set_proposal"

var _ sdk.Msg = &MsgSubmitSetProposal{}

func NewMsgSubmitSetProposal(creator string, setId uint64) *MsgSubmitSetProposal {
	return &MsgSubmitSetProposal{
		Creator: creator,
		SetId:   setId,
	}
}

func (msg *MsgSubmitSetProposal) Route() string {
	return RouterKey
}

func (msg *MsgSubmitSetProposal) Type() string {
	return TypeMsgSubmitSetProposal
}

func (msg *MsgSubmitSetProposal) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSubmitSetProposal) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSubmitSetProposal) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(errors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
