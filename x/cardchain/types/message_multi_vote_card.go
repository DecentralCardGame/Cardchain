package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgMultiVoteCard = "multi_vote_card"

var _ sdk.Msg = &MsgMultiVoteCard{}

func NewMsgMultiVoteCard(creator string) *MsgMultiVoteCard {
	return &MsgMultiVoteCard{
		Creator: creator,
	}
}

func (msg *MsgMultiVoteCard) Route() string {
	return RouterKey
}

func (msg *MsgMultiVoteCard) Type() string {
	return TypeMsgMultiVoteCard
}

func (msg *MsgMultiVoteCard) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgMultiVoteCard) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgMultiVoteCard) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
