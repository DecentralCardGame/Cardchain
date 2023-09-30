package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRemoveContributorFromSet = "remove_contributor_from_set"

var _ sdk.Msg = &MsgRemoveContributorFromSet{}

func NewMsgRemoveContributorFromSet(creator string, setId uint64, user string) *MsgRemoveContributorFromSet {
	return &MsgRemoveContributorFromSet{
		Creator:      creator,
		SetId: setId,
		User:         user,
	}
}

func (msg *MsgRemoveContributorFromSet) Route() string {
	return RouterKey
}

func (msg *MsgRemoveContributorFromSet) Type() string {
	return TypeMsgRemoveContributorFromSet
}

func (msg *MsgRemoveContributorFromSet) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRemoveContributorFromSet) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRemoveContributorFromSet) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(errors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
