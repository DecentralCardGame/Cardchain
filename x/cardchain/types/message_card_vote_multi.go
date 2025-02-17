package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCardVoteMulti{}

func NewMsgCardVoteMulti(creator string, votes []*SingleVote) *MsgCardVoteMulti {
	return &MsgCardVoteMulti{
		Creator: creator,
		Votes:   votes,
	}
}

func (msg *MsgCardVoteMulti) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
