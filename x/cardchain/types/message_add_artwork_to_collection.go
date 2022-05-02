package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAddArtworkToCollection = "add_artwork_to_collection"

var _ sdk.Msg = &MsgAddArtworkToCollection{}

func NewMsgAddArtworkToCollection(creator string, collectionId uint64, image []byte) *MsgAddArtworkToCollection {
	return &MsgAddArtworkToCollection{
		Creator:      creator,
		CollectionId: collectionId,
		Image:        image,
	}
}

func (msg *MsgAddArtworkToCollection) Route() string {
	return RouterKey
}

func (msg *MsgAddArtworkToCollection) Type() string {
	return TypeMsgAddArtworkToCollection
}

func (msg *MsgAddArtworkToCollection) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddArtworkToCollection) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddArtworkToCollection) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
