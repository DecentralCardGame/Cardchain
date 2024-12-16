package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetArtistSet{}

func NewMsgSetArtistSet(creator string, setId uint64, artist string) *MsgSetArtistSet {
	return &MsgSetArtistSet{
		Creator: creator,
		SetId:   setId,
		Artist:  artist,
	}
}

func (msg *MsgSetArtistSet) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
