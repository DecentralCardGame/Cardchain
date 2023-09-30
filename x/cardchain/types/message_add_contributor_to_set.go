package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAddContributorToSet = "add_contributor_to_set"

var _ sdk.Msg = &MsgAddContributorToSet{}

func NewMsgAddContributorToSet(creator string, setId uint64, user string) *MsgAddContributorToSet {
	return &MsgAddContributorToSet{
		Creator:      creator,
		SetId: setId,
		User:         user,
	}
}

func (msg *MsgAddContributorToSet) Route() string {
	return RouterKey
}

func (msg *MsgAddContributorToSet) Type() string {
	return TypeMsgAddContributorToSet
}

func (msg *MsgAddContributorToSet) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddContributorToSet) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddContributorToSet) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(errors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
