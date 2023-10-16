package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAddCardToSet = "add_card_to_set"

var _ sdk.Msg = &MsgAddCardToSet{}

func NewMsgAddCardToSet(creator string, setId uint64, cardId uint64) *MsgAddCardToSet {
	return &MsgAddCardToSet{
		Creator: creator,
		SetId:   setId,
		CardId:  cardId,
	}
}

func (msg *MsgAddCardToSet) Route() string {
	return RouterKey
}

func (msg *MsgAddCardToSet) Type() string {
	return TypeMsgAddCardToSet
}

func (msg *MsgAddCardToSet) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddCardToSet) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddCardToSet) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(errors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
