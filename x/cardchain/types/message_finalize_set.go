package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgFinalizeSet = "finalize_set"

var _ sdk.Msg = &MsgFinalizeSet{}

func NewMsgFinalizeSet(creator string, setId uint64) *MsgFinalizeSet {
	return &MsgFinalizeSet{
		Creator:      creator,
		SetId: setId,
	}
}

func (msg *MsgFinalizeSet) Route() string {
	return RouterKey
}

func (msg *MsgFinalizeSet) Type() string {
	return TypeMsgFinalizeSet
}

func (msg *MsgFinalizeSet) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgFinalizeSet) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgFinalizeSet) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(errors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
