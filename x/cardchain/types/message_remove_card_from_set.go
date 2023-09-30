package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRemoveCardFromSet = "remove_card_from_set"

var _ sdk.Msg = &MsgRemoveCardFromSet{}

func NewMsgRemoveCardFromSet(creator string, setId uint64, cardId uint64) *MsgRemoveCardFromSet {
	return &MsgRemoveCardFromSet{
		Creator:      creator,
		SetId: setId,
		CardId:       cardId,
	}
}

func (msg *MsgRemoveCardFromSet) Route() string {
	return RouterKey
}

func (msg *MsgRemoveCardFromSet) Type() string {
	return TypeMsgRemoveCardFromSet
}

func (msg *MsgRemoveCardFromSet) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRemoveCardFromSet) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRemoveCardFromSet) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(errors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
