package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgProfileCardSet{}

func NewMsgProfileCardSet(creator string, cardId uint64) *MsgProfileCardSet {
	return &MsgProfileCardSet{
		Creator: creator,
		CardId:  cardId,
	}
}

func (msg *MsgProfileCardSet) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
