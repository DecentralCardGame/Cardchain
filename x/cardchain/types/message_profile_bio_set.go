package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const BIO_MAX_LENGTH = 400

var _ sdk.Msg = &MsgProfileBioSet{}

func NewMsgProfileBioSet(creator string, bio string) *MsgProfileBioSet {
	return &MsgProfileBioSet{
		Creator: creator,
		Bio:     bio,
	}
}

func (msg *MsgProfileBioSet) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if len(msg.Bio) > BIO_MAX_LENGTH {
		return errorsmod.Wrapf(ErrInvalidData, "Website length exceded %d chars", BIO_MAX_LENGTH)
	}

	return nil
}
