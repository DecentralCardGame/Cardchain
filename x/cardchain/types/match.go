package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// getMatchAddresses Get's and verifies the players of a match
func (m Match) GetMatchAddresses() (addresses []sdk.AccAddress, err error) {
	for _, player := range []string{m.PlayerA.Addr, m.PlayerB.Addr} {
		var address sdk.AccAddress
		address, err = sdk.AccAddressFromBech32(player)
		if err != nil {
			err = errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "Invalid player")
			return
		}
		addresses = append(addresses, address)
	}

	return
}
