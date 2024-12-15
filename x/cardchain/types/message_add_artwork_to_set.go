package types

import (
	fmt "fmt"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAddArtworkToSet = "add_artwork_to_set"

var _ sdk.Msg = &MsgAddArtworkToSet{}

func NewMsgAddArtworkToSet(creator string, setId uint64, image []byte) *MsgAddArtworkToSet {
	return &MsgAddArtworkToSet{
		Creator: creator,
		SetId:   setId,
		Image:   image,
	}
}

func (msg *MsgAddArtworkToSet) Route() string {
	return RouterKey
}

func (msg *MsgAddArtworkToSet) Type() string {
	return TypeMsgAddArtworkToSet
}

func (msg *MsgAddArtworkToSet) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddArtworkToSet) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddArtworkToSet) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(errors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if len(msg.Image) > ArtworkMaxSize {
		return sdkerrors.Wrap(ErrImageSizeExceeded, fmt.Sprint(len(msg.Image)))
	}
	return nil
}
