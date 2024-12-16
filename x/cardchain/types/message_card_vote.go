package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCardVote = "card_vote"

var _ sdk.Msg = &MsgCardVote{}

func NewMsgCardVote(creator string, vote *SingleVote) *MsgCardVote {
  return &MsgCardVote{
		Creator: creator,
    Vote: vote,
	}
}

func (msg *MsgCardVote) Route() string {
  return RouterKey
}

func (msg *MsgCardVote) Type() string {
  return TypeMsgCardVote
}

func (msg *MsgCardVote) GetSigners() []sdk.AccAddress {
  creator, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    panic(err)
  }
  return []sdk.AccAddress{creator}
}

func (msg *MsgCardVote) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgCardVote) ValidateBasic() error {
  _, err := sdk.AccAddressFromBech32(msg.Creator)
  	if err != nil {
  		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
  	}
  return nil
}

