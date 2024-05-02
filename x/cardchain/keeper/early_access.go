package keeper

import (
	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
)

func removeEarlyAccessFromUser(invited *User) {
	invited.EarlyAccess.Active = false
	invited.EarlyAccess.InvitedByUser = ""
}

func AddEarlyAccessToUser(user *User, inviter string) {
	user.EarlyAccess = &types.EarlyAccess{
		Active:        true,
		InvitedByUser: inviter,
	}
}
