package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCardRaritySet{}

func NewMsgCardRaritySet(creator string, cardId uint64, setId uint64, rarity CardRarity) *MsgCardRaritySet {
	return &MsgCardRaritySet{
		Creator: creator,
		CardId:  cardId,
		SetId:   setId,
		Rarity:  rarity,
	}
}

func (msg *MsgCardRaritySet) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
