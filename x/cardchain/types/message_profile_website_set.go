package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const WEBSITE_MAX_LENGTH = 50

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

	if len(msg.Website) > WEBSITE_MAX_LENGTH {
		return errorsmod.Wrapf(ErrInvalidData, "Website length exceded %d chars", WEBSITE_MAX_LENGTH)
	}

	return nil
}
