package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetCollectionArtist = "set_collection_artist"

var _ sdk.Msg = &MsgSetCollectionArtist{}

func NewMsgSetCollectionArtist(creator string, collectionId uint64, artist string) *MsgSetCollectionArtist {
	return &MsgSetCollectionArtist{
		Creator:      creator,
		CollectionId: collectionId,
		Artist:       artist,
	}
}

func (msg *MsgSetCollectionArtist) Route() string {
	return RouterKey
}

func (msg *MsgSetCollectionArtist) Type() string {
	return TypeMsgSetCollectionArtist
}

func (msg *MsgSetCollectionArtist) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetCollectionArtist) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetCollectionArtist) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(errors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
