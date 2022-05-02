package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetCardRarity = "set_card_rarity"

var _ sdk.Msg = &MsgSetCardRarity{}

func NewMsgSetCardRarity(creator string, cardId uint64, collectionId uint64, rarity string) *MsgSetCardRarity {
	return &MsgSetCardRarity{
		Creator:      creator,
		CardId:       cardId,
		CollectionId: collectionId,
		Rarity:       rarity,
	}
}

func (msg *MsgSetCardRarity) Route() string {
	return RouterKey
}

func (msg *MsgSetCardRarity) Type() string {
	return TypeMsgSetCardRarity
}

func (msg *MsgSetCardRarity) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetCardRarity) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetCardRarity) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
