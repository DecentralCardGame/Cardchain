package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetSetArtist = "set_set_artist"

var _ sdk.Msg = &MsgSetSetArtist{}

func NewMsgSetSetArtist(creator string, setId uint64, artist string) *MsgSetSetArtist {
	return &MsgSetSetArtist{
		Creator: creator,
		SetId:   setId,
		Artist:  artist,
	}
}

func (msg *MsgSetSetArtist) Route() string {
	return RouterKey
}

func (msg *MsgSetSetArtist) Type() string {
	return TypeMsgSetSetArtist
}

func (msg *MsgSetSetArtist) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetSetArtist) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetSetArtist) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(errors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
