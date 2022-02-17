package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgVoteCard = "vote_card"

var _ sdk.Msg = &MsgVoteCard{}

func NewMsgVoteCard(creator string, cardId uint64, voteType string) *MsgVoteCard {
	return &MsgVoteCard{
		Creator:  creator,
		CardId:   cardId,
		VoteType: voteType,
	}
}

func (msg *MsgVoteCard) Route() string {
	return RouterKey
}

func (msg *MsgVoteCard) Type() string {
	return TypeMsgVoteCard
}

func (msg *MsgVoteCard) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgVoteCard) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgVoteCard) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
