package types

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
