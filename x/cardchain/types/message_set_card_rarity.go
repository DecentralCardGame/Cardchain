package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetCardRarity = "set_card_rarity"

var _ sdk.Msg = &MsgSetCardRarity{}

func NewMsgSetCardRarity(creator string, cardId uint64, setId uint64, rarity CardRarity) *MsgSetCardRarity {
	return &MsgSetCardRarity{
		Creator: creator,
		CardId:  cardId,
		SetId:   setId,
		Rarity:  rarity,
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
		return sdkerrors.Wrapf(errors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	_, isValid := CardRarity_name[int32(msg.Rarity)]
	if !isValid {
		return sdkerrors.Wrapf(errors.ErrInvalidRequest, "Invalid cardRarity: (%d)", msg.Rarity)
	}

	return nil
}
