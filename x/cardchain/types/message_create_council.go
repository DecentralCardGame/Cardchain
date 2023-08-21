package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateCouncil = "create_council"

var _ sdk.Msg = &MsgCreateCouncil{}

func NewMsgCreateCouncil(creator string, cardId uint64) *MsgCreateCouncil {
	return &MsgCreateCouncil{
		Creator: creator,
		CardId:  cardId,
	}
}

func (msg *MsgCreateCouncil) Route() string {
	return RouterKey
}

func (msg *MsgCreateCouncil) Type() string {
	return TypeMsgCreateCouncil
}

func (msg *MsgCreateCouncil) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateCouncil) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateCouncil) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(errors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
