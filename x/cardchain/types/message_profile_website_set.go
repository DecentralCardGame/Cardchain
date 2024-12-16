package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgProfileWebsiteSet{}

func NewMsgProfileWebsiteSet(creator string, website string) *MsgProfileWebsiteSet {
	return &MsgProfileWebsiteSet{
		Creator: creator,
		Website: website,
	}
}

func (msg *MsgProfileWebsiteSet) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
