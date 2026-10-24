package types

import (
	"github.com/DecentralCardGame/cardchain/util"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (u *User) RemoveEarlyAccess() {
	u.EarlyAccess.Active = false
	u.EarlyAccess.InvitedByUser = ""
}

func (u *User) AddEarlyAccess(inviter string) {
	u.EarlyAccess = &EarlyAccess{
		Active:        true,
		InvitedByUser: inviter,
	}
}

// TransferSchemeToCard Makes a users cardscheme a card
func (u *User) SchemeToCard(cardId uint64) (err error) {
	u.OwnedCardSchemes, err = util.PopItemFromArr(cardId, u.OwnedCardSchemes)
	if err != nil {
		return sdkerrors.ErrUnauthorized
	}

	u.OwnedPrototypes = append(u.OwnedPrototypes, cardId)
	return nil
}
