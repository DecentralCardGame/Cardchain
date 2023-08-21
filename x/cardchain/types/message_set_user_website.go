package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetUserWebsite = "set_user_website"

var _ sdk.Msg = &MsgSetUserWebsite{}

func NewMsgSetUserWebsite(creator string, website string) *MsgSetUserWebsite {
	return &MsgSetUserWebsite{
		Creator: creator,
		Website: website,
	}
}

func (msg *MsgSetUserWebsite) Route() string {
	return RouterKey
}

func (msg *MsgSetUserWebsite) Type() string {
	return TypeMsgSetUserWebsite
}

func (msg *MsgSetUserWebsite) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetUserWebsite) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetUserWebsite) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(errors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
