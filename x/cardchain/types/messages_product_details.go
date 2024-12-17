package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateProductDetails{}

func NewMsgCreateProductDetails(creator string, name string, desc string) *MsgCreateProductDetails {
	return &MsgCreateProductDetails{
		Creator: creator,
		Name:    name,
		Desc:    desc,
	}
}

func (msg *MsgCreateProductDetails) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateProductDetails{}

func NewMsgUpdateProductDetails(creator string, id uint64, name string, desc string) *MsgUpdateProductDetails {
	return &MsgUpdateProductDetails{
		Id:      id,
		Creator: creator,
		Name:    name,
		Desc:    desc,
	}
}

func (msg *MsgUpdateProductDetails) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteProductDetails{}

func NewMsgDeleteProductDetails(creator string, id uint64) *MsgDeleteProductDetails {
	return &MsgDeleteProductDetails{
		Id:      id,
		Creator: creator,
	}
}

func (msg *MsgDeleteProductDetails) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
